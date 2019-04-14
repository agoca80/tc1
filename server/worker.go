package server

import (
	"bufio"
	"io"
)

func (s *service) worker(clients <-chan io.ReadCloser, numbers chan<- int) {
	for client := range clients {
		s.serialize(client, numbers)
	}

	s.Done()
}

func (s *service) serialize(conn io.ReadCloser, numbers chan<- int) {
	reader := bufio.NewReader(conn)

	for s.Running() {
		number, err := newNumber(reader)
		switch {
		case err != nil || number == invalid:
			conn.Close()
			return

		case number == terminate:
			close(s.terminate)
			s.Close()

		default:
			numbers <- number
		}
	}
}
