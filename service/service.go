package service

import (
	"io"
	"time"

	"github.com/agoca80/tc1/dispatcher"
	"github.com/agoca80/tc1/process"

	"github.com/agoca80/tc1/recorder"

	"github.com/agoca80/tc1/filter"
	"github.com/agoca80/tc1/runner"
)

// Service ...
type Service struct {
	runner.Runner
	reports time.Duration
	workers int

	dispatcher.Dispatcher
	process.Pool
	filter.Filter
	recorder.Recorder
}

// New ...
func New(workers, reports, size int, input, output string) (s *Service, err error) {
	recorder, err := recorder.New(output)
	if err != nil {
		return
	}

	service := runner.New()

	s = &Service{
		Runner:     service,
		reports:    time.Duration(reports) * time.Millisecond,
		Dispatcher: dispatcher.New(service),
		Pool:       process.New(workers, service),
		Filter:     filter.New(input, size),
		Recorder:   recorder,
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
		clock   = time.NewTicker(s.reports)
	)
	defer clock.Stop()

	go s.Dispatch(clients)
	go s.Process(clients, numbers)
	go s.Filter.Run(numbers, uniques)
	go s.Record(uniques)

	for {
		select {

		case <-s.Runner:
			s.Close()
			report()
			return

		case <-clock.C:
			report()
		}
	}
}
