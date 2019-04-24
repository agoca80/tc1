package service

import (
	"fmt"

	"github.com/agoca80/tc1/filter"
)

func (s *Service) reporter() func() {
	var stats filter.Statistics
	return func() {
		current := s.Filter.Stats()
		fmt.Printf(
			"Received %v unique numbers, %v duplicates. Unique total: %v\n",
			current.Uniques-stats.Uniques,
			current.Duplicates-stats.Duplicates,
			current.Uniques,
		)
		stats = current
	}
}
