package recorder

import (
	"fmt"
	"io"
)

type recorder struct {
	io.WriteCloser
}

// Record ...
func (r *recorder) Record(numbers <-chan int) {
	for number := range numbers {
		fmt.Fprintln(r, number)
	}
	r.Close()
}
