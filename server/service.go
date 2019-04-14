package server

import (
	"io"
	"net"
)

// Service ...
type Service struct {
	terminate chan bool
	Stats

	net.Listener
	input  io.Writer
	output io.Writer

	Memory
}

// New ...
func New(listener net.Listener, input, output io.Writer) *Service {

	return &Service{
		Listener:  listener,
		terminate: make(chan bool),
		input:     input,
		output:    output,
	}
}

// Start ...
func (s *Service) Start() {
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

func (s *Service) Running() bool {
	select {
	case <-s.terminate:
		return false
	default:
		return true
	}
}
