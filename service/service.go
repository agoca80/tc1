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
	reports time.Duration
	workers int

	net.Listener
	input  io.Writer
	output io.Writer

	Memory memory.Interface

	runner
}

// New ...
func New(workers, reports, size int, listener net.Listener, input, output io.Writer, memory memory.Interface) *Service {
	return &Service{
		reports:  time.Duration(reports) * time.Millisecond,
		Listener: listener,
		Memory:   memory,
		input:    input,
		output:   output,
		workers:  workers,
		runner:   newRunner(),
	}
}

// Start ...
func (s *Service) Start() {
	var (
		clients = make(chan io.ReadCloser)
		numbers = make(chan int, s.workers)
		uniques = make(chan int)
		report  = s.reporter()
	)

	go s.dispatcher(clients)
	go s.newPool(clients, numbers)
	go s.filter(numbers, uniques)
	go s.record(uniques)

	clock := time.NewTicker(s.reports)
	defer clock.Stop()
	for {
		select {

		case <-s.runner:
			report()
			return

		case <-clock.C:
			report()
		}
	}
}

func (s *Service) reporter() func() {
	var stats memory.Statistics
	return func() {
		current := s.Memory.Stats()
		fmt.Printf(
			"Received %v unique numbers, %v duplicates. Unique total: %v\n",
			current.Uniques-stats.Uniques,
			current.Duplicates-stats.Duplicates,
			current.Uniques,
		)
		stats = current
	}
}
