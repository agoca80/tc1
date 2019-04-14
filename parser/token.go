package parser

import (
	"strconv"
	"strings"
)

// Token ...
type Token []byte

// Number ...
func (t Token) Number() (number int, ok bool) {
	number, err := strconv.Atoi(strings.TrimSpace(string(t)))
	if err != nil {
		return
	}

	return number, true
}

// String ...
func (t Token) String() string {
	return string(t)
}
