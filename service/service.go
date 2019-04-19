package service

import (
	"io"
	"net"

	"github.com/agoca80/tc1/memory"
)

// Service ...
type Service struct {
	terminate chan bool
	Stats

	net.Listener
	input  io.Writer
	output io.Writer

	memory.Persistent
}

// New ...
func New(listener net.Listener, input, output io.Writer, memory memory.Persistent) *Service {
	return &Service{
		Listener:   listener,
		Persistent: memory,
		terminate:  make(chan bool),
		input:      input,
		output:     output,
	}
}

// Start ...
func (s *Service) Start() {
	var (
		clients = make(chan io.ReadCloser)
		numbers = make(chan int, 5*1024)
		uniques = make(chan int)
	)

	go s.reporter(Report)
	go s.dispatcher(clients)
	go NewPool(
		Clients,
		func() { s.process(clients, numbers) },
		func() { close(numbers) },
	)
	go s.filter(numbers, uniques)

	s.record(uniques)
}

// Running ...
func (s *Service) Running() bool {
	select {
	case <-s.terminate:
		return false
	default:
		return true
	}
}
