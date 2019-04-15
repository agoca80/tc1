package parser

import (
	"bufio"
	"io"
)

const (
	Invalid   = "invalid"
	Number    = "number"
	Terminate = "terminate"
	Finish    = "finish"
)

// Parser ...
type Parser struct {
	Number int
	*bufio.Reader
}

// New ...
func New(reader io.Reader) (p *Parser) {
	return &Parser{
		Reader: bufio.NewReader(reader),
	}
}

// Next ...
func (p *Parser) Next() string {
	token, err := p.ReadString('\n')
	if err != nil || len(token) != 10 {
		return Invalid
	}

	if string(token) == "terminate\n" {
		return Terminate
	}

	number, ok := Token(token[0:9]).Number()
	if !ok {
		return Invalid
	}

	p.Number = number
	return Number
}
