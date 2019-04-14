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
	net.Listener
	numbers chan int
	input   io.Writer
	uniques chan int
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
		numbers:   make(chan int, 5*1024),
		terminate: make(chan bool),
		uniques:   make(chan int),

		Writer: writer,
		input:  input,
	}
}

// Start ...
func Start() {
	s := newService()

	var (
		clients = make(chan io.ReadCloser)
	)

	s.remind()

	go s.reporter()
	go s.dispatcher(clients)
	for i := 0; i < 5; i++ {
		s.Add(1)
		go s.worker(clients, s.numbers)
	}

	go s.filter(s.numbers, s.uniques)

	go func() {
		s.Wait()
		close(s.numbers)
	}()

	for unique := range (<-chan int)(s.uniques) {
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
