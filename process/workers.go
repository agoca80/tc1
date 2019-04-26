package process

import (
	"io"
	"sync"

	"github.com/agoca80/tc1/runner"
)

type pool struct {
	size int

	runner.Runner
	sync.WaitGroup
}

// New ...
func New(size int, service runner.Runner) Pool {
	return &pool{
		Runner: service,
		size:   size,
	}
}

// Process ...
func (p *pool) Process(clients <-chan io.ReadCloser, numbers chan<- int) {
	for i := 0; i < p.size; i++ {
		p.Add(1)
		go func() {
			p.process(clients, numbers)
			p.Done()
		}()
	}

	go func() {
		p.Wait()
		close(numbers)
	}()
}
