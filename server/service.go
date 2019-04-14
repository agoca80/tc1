package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

// Testing ...
var Testing = false

type service struct {
	terminate  chan bool
	Duplicates int
	Uniques    int
	Total      int
	input      io.Writer

	net.Listener
	io.Writer
	Memory
	sync.WaitGroup
}

func newService() *service {
	writer, err := os.OpenFile("numbers.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	input, err := os.OpenFile("input", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	return &service{
		Listener:  listener,
		terminate: make(chan bool),
		Writer:    writer,
		input:     input,
	}
}

// Start ...
func Start() {
	s := newService()

	var (
		clients = make(chan io.ReadCloser)
		numbers = make(chan int, 5*1024)
		uniques = make(chan int)
	)

	s.remind()

	go s.reporter()
	go s.dispatcher(clients)
	for i := 0; i < 5; i++ {
		s.Add(1)
		go s.worker(clients, numbers)
	}

	go s.filter(numbers, uniques)

	go func() {
		s.Wait()
		close(numbers)
	}()

	for unique := range (<-chan int)(uniques) {
		fmt.Fprintln(s, unique)
	}

	s.store()
}

func (s *service) Running() bool {
	select {
	case <-s.terminate:
		return false
	default:
		return true
	}
}

func (s *service) String() string {
	return fmt.Sprintf("Received %v unique numbers, %v duplicates. Unique total: %v", s.Uniques, s.Duplicates, s.Total)
}
