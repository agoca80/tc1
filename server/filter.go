package server

import (
	"fmt"
	"io"
	"os"
)

// filter processes the numbers from the workers
func (s *service) filter(numbers <-chan int, uniques chan<- int) {
	var err error
	var input io.WriteCloser
	if Testing == true {
		input, err = os.OpenFile("input", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer input.Close()
	}

	for number := range numbers {
		// Only for testing purposes
		if Testing {
			fmt.Fprintln(input, number)
		}
		if s.remembers(number) {
			s.Duplicates++
		} else {
			s.Uniques++
			s.Total++
			uniques <- number
		}
	}
	close(uniques)

	// Print last report
	fmt.Println(s)
}
