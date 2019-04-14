package server

import (
	"fmt"
	"time"
)

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
