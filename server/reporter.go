package server

import (
	"fmt"
	"time"
)

// reporter spends most of the time waiting
// It prints the service uniques, duplicates and total uniques count
//
// As service implements Stringer getting it from the Report type,
// we don't need any additional logic.
//
// NOTE: The reporting statistics alre never locked:
// - Only the filter updates the total uniques counter
// - Uniques and Duplicates are updated from the filter
//   and from the report String() method, so, there is no strong consistency
//   Eventually consistency will be provided by stopping the filter
func (s *service) reporter() {
	clock := time.NewTicker(10 * time.Second)
	for {
		select {

		case <-s.terminate:
			clock.Stop()
			return

		case <-clock.C:
			fmt.Println(s)
			s.Uniques, s.Duplicates = 0, 0

		}
	}
}
