package server

import (
	"io"

	"github.com/agoca80/tc1/parser"
)

func (s *service) process(clients <-chan io.ReadCloser, numbers chan<- int) {
	for client := range clients {
		s.serialize(client, numbers)
		client.Close()
	}
}

func (s *service) serialize(client io.Reader, numbers chan<- int) {
	p := parser.New(client)

	for s.Running() {
		switch p.Next() {
		case parser.Invalid, parser.Finish:
			return

		case parser.Terminate:
			close(s.terminate)
			s.Close()
			return

		case parser.Number:
			numbers <- p.Number
		}
	}
}
