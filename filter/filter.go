package filter

import (
	"fmt"
	"io"

	"github.com/agoca80/tc1/memory"
)

type filter struct {
	dump string
	Statistics
	io.WriteCloser
	memory.Memory
}

// Run ...
func (f *filter) Run(numbers <-chan int, uniques chan<- int) {
	for number := range numbers {
		fmt.Fprintln(f, number)
		if f.Remembers(number) {
			f.Duplicates++
		} else {
			f.Uniques++
			uniques <- number
		}
	}
	close(uniques)
}

// Stop ...
func (f *filter) Stop() error {
	return f.Dump(f.dump)
}
