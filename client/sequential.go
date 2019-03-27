package client

import "fmt"

// Sequential generates a monotonic incremental number generator
type sequential struct {
	current int
	first   int
	max     int
	step    int
}

// Sequential generates a monotonic incremental number generator
func Sequential(first, max, step int) Generator {
	return &sequential{current: first, first: first, max: max, step: step}
}

// Number ...
func (s *sequential) Number() (n []byte) {
	n = []byte(fmt.Sprintf("%09d\n", s.current))
	s.current += s.step
	if s.current > s.max {
		s.current = s.first
	}
	return
}
