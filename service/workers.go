package service

import (
	"io"
	"sync"
)

func (s *Service) newPool(clients <-chan io.ReadCloser, numbers chan<- int) {
	var pool sync.WaitGroup
	for i := 0; i < s.workers; i++ {
		pool.Add(1)
		go func() {
			s.process(clients, numbers)
			pool.Done()
		}()
	}

	go func() {
		pool.Wait()
		close(numbers)
	}()
}
