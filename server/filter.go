package server

import "fmt"

func (s *Service) filter(numbers <-chan int, uniques chan<- int) {
	for number := range numbers {
		fmt.Fprintln(s.input, number)
		if s.remembers(number) {
			s.Duplicates++
		} else {
			s.Uniques++
			s.Total++
			uniques <- number
		}
	}
	close(uniques)

	fmt.Println(s)
}
