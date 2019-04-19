package service

import "sync"

// NewPool ...
func NewPool(size int, work func(), finish func()) func() {
	var pool sync.WaitGroup
	for ; 0 < size; size-- {
		pool.Add(1)
		go func() {
			work()
			pool.Done()
		}()
	}

	return func() {
		pool.Wait()
		finish()
	}
}
