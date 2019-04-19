package service

import "fmt"

// Stats ...
type Stats struct {
	Uniques    int
	Duplicates int
	Total      int
}

func (s *Stats) String() string {
	return fmt.Sprintf("Received %v unique numbers, %v duplicates. Unique total: %v", s.Uniques, s.Duplicates, s.Total)
}
