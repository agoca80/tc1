package server

import "fmt"

// filter processes the numbers from the workers
func (s *service) filter(numbers <-chan int, uniques chan<- int) {
	for number := range numbers {
		// Only for testing purposes
		if Testing {
			fmt.Fprintln(s.input, number)
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
