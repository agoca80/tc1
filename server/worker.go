package server

import (
	"bufio"
	"io"
)

func (s *service) worker(numbers chan<- int) {
	for !s.terminating() {
		conn, err := s.Accept()
		if err == nil {
			s.serialize(conn, numbers)
		}
	}

	s.Done()
}

func (s *service) serialize(conn io.ReadCloser, numbers chan<- int) {
	reader := bufio.NewReader(conn)

	for !s.terminating() {
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
