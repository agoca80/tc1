package service

import (
	"io"

	"github.com/agoca80/tc1/parser"
)

func (p *pool) process(clients <-chan io.ReadCloser, numbers chan<- int) {
	for client := range clients {
		p.serialize(client, numbers)
		client.Close()
	}
}

func (p *pool) serialize(client io.Reader, numbers chan<- int) {
	stream := parser.New(client)

	for p.Running() {
		switch stream.Next() {
		case parser.Invalid, parser.Finish:
			return

		case parser.Terminate:
			p.Stop()
			return

		case parser.Number:
			numbers <- stream.Number
		}
	}
}
