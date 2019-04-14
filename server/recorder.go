package server

import "fmt"

func (s *service) record(uniques <-chan int) {
	for unique := range uniques {
		fmt.Fprintln(s, unique)
	}
}
