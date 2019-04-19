package service

import (
	"fmt"
	"time"
)

func (s *Service) reporter(period time.Duration) {
	clock := time.NewTicker(period)
	defer clock.Stop()
	for {
		select {

		case <-s.terminate:
			return

		case <-clock.C:
			fmt.Println(s)
			s.Uniques, s.Duplicates = 0, 0

		}
	}
}
