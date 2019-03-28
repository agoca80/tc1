package server

import (
	"fmt"
	"io"
)

// filter processes the numbers from the workers
func (s *service) filter(numbers <-chan int, uniques chan<- int) {
	for number := range numbers {
		// Only for testing purposes
		if Testing {
			io.WriteString(s.input, fmt.Sprintf("%09d\n", number))
		}
		if s.Remembers(number) {
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
