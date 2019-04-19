package service

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/agoca80/tc1/memory"
)

// Service ...
type Service struct {
	clients   int
	reports   time.Duration
	size      int
	terminate chan bool
	Stats

	net.Listener
	input  io.Writer
	output io.Writer

	Memory memory.Interface
}

// New ...
func New(clients, reports, size int, listener net.Listener, input, output io.Writer, memory memory.Interface) *Service {
	return &Service{
		size:      size,
		clients:   clients,
		reports:   time.Duration(reports) * time.Millisecond,
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
		numbers = make(chan int, s.clients)
		uniques = make(chan int)
	)

	go s.dispatcher(clients)
	go s.newPool(clients, numbers)
	go s.filter(numbers, uniques)
	go s.record(uniques)

	clock := time.NewTicker(s.reports)
	defer clock.Stop()
	for {
		select {

		case <-s.terminate:
			return

		case <-clock.C:
			fmt.Println(s)
			s.Uniques, s.Duplicates = 0, 0

		}
	}
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
