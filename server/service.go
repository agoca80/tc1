package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Testing ...
var Testing = false

type service struct {
	terminate  chan bool
	Duplicates int
	Uniques    int
	Total      int

	net.Listener
	input  io.Writer
	output io.Writer

	Memory
}

func newService() *service {
	output, err := os.OpenFile("numbers.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
		input:     input,
		output:    output,
	}
}

// Start ...
func Start() {
	s := newService()
	s.remind()

	var (
		clients = make(chan io.ReadCloser)
		numbers = make(chan int, 5*1024)
		uniques = make(chan int)
	)

	NewWorker(func() { s.reporter(Report) })
	NewWorker(func() { s.dispatcher(clients) })
	NewWorker(NewPool(
		Clients,
		func() { s.process(clients, numbers) },
		func() { close(numbers) },
	))
	NewWorker(func() { s.filter(numbers, uniques) })

	s.record(uniques)
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
