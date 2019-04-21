package service

import (
	"io"
	"sync"
)

type pool struct {
	size int

	runner
	sync.WaitGroup
}

func newPool(size int, leader runner) *pool {
	return &pool{
		runner: leader,
		size:   size,
	}
}

func (p *pool) run(clients <-chan io.ReadCloser, numbers chan<- int) {
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
