package service

import (
	"fmt"
	"io"
	"time"

	"github.com/agoca80/tc1/memory"
)

// Service ...
type Service struct {
	reports time.Duration
	workers int

	input  io.Writer
	output io.Writer

	Memory memory.Interface

	Runner
	*dispatcher
	*pool
}

// New ...
func New(workers, reports, size int, input, output io.Writer, memory memory.Interface) *Service {
	service := NewRunner()

	return &Service{
		reports:    time.Duration(reports) * time.Millisecond,
		Memory:     memory,
		input:      input,
		output:     output,
		workers:    workers,
		Runner:     service,
		dispatcher: newDispatcher(service),
		pool:       newPool(workers, service),
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

	go s.dispatcher.run(clients)
	go s.pool.run(clients, numbers)
	go s.filter(numbers, uniques)
	go s.record(uniques)

	clock := time.NewTicker(s.reports)
	defer clock.Stop()
	for {
		select {

		case <-s.Runner:
			s.dispatcher.Close()
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
