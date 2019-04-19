package service

import "fmt"

func (s *Service) record(uniques <-chan int) {
	for unique := range uniques {
		fmt.Fprintln(s.output, unique)
	}
}
