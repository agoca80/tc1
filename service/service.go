package service

import (
	"io"
	"time"

	"github.com/agoca80/tc1/filter"
)

// Service ...
type Service struct {
	reports time.Duration
	workers int

	output io.Writer

	Runner
	*dispatcher
	*pool
	filter.Filter
}

// New ...
func New(workers, reports, size int, input string, output io.Writer) (s *Service, err error) {
	service := NewRunner()
	filter := filter.New(input, size)

	s = &Service{
		Filter:     filter,
		reports:    time.Duration(reports) * time.Millisecond,
		output:     output,
		workers:    workers,
		Runner:     service,
		dispatcher: newDispatcher(service),
		pool:       newPool(workers, service),
	}

	return
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
	go s.Filter.Run(numbers, uniques)
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
