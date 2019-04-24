package filter

import (
	"os"

	"github.com/agoca80/tc1/memory"
)

// Filter ...
type Filter interface {
	Run(<-chan int, chan<- int)
	Stats() Statistics
}

// New ...
func New(input string, size int) Filter {
	writeCloser, err := os.OpenFile(input, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	return &filter{
		WriteCloser: writeCloser,
		Memory:      memory.NewMemory(size),
	}
}
