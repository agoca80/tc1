package service

import (
	"io"
	"net"

	"github.com/agoca80/tc1/memory"
)

// Service ...
type Service struct {
	size      int
	terminate chan bool
	Stats

	net.Listener
	input  io.Writer
	output io.Writer

	Memory memory.Interface
}

// New ...
func New(size int, listener net.Listener, input, output io.Writer, memory memory.Interface) *Service {
	return &Service{
		size:      size,
		Listener:  listener,
		Memory:    memory,
		terminate: make(chan bool),
		input:     input,
		output:    output,
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
	go s.newPool(clients, numbers)
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
