package service

import "fmt"

func (s *Service) filter(numbers <-chan int, uniques chan<- int) {
	for number := range numbers {
		fmt.Fprintln(s.input, number)
		if !s.Memory.Remembers(number) {
			uniques <- number
		}
	}
	close(uniques)
}
