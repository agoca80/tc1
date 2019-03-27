package server

import (
	"bufio"
	"io"
)

// worker accept incoming connections
// as long as the termination signal is not broadcasted
func (s *service) worker(numbers chan<- int) {
	for !s.terminating() {
		conn, err := s.Accept()
		if err == nil {
			s.serialize(conn, numbers)
		}
	}

	// notify the waiter goroutine that the worker as finished
	s.Done()
}

// serialize will keep running for an incoming connection
// as long as the termination signal is not broadcasted
// or there is any error reading the connection
//
// Using conn as io.ReadCloser makes testing it easier,
// as it does not require a listening server to provider net connection
func (s *service) serialize(conn io.ReadCloser, numbers chan<- int) {
	reader := bufio.NewReader(conn)

	for !s.terminating() {
		number, err := newNumber(reader)
		switch {
		case err != nil:
			// error while reading from the connection
			// close the connection and go back to process another one
			conn.Close()
			return

		case number == terminate:
			// Broadcast the termination signal and go back
			close(s.terminate)
			s.Close()

		case number != invalid:
			// The number is valid, send it to the filter
			numbers <- number
		}
	}
}
