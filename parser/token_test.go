package parser

import (
	"fmt"
	"testing"
)

func token(str string) Token {
	return Token([]byte(str))
}

func TestToken_Number(t *testing.T) {
	type check struct {
		Token
		number int
		ok     bool
	}

	fail := func(c check) {
		fmt.Printf("FAIL %+v\n", c)
		t.Fail()
	}

	checks := []check{
		check{token(""), 0, false},
		check{token("aaaaaaaaa"), 0, false},
		check{token("terminate"), 0, false},
		check{token("000000000"), 0, true},
	}

	for _, c := range checks {
		number, ok := c.Number()
		if ok != c.ok {
			fail(c)
		} else if ok && number != c.number {
			fail(c)
		}
	}
}
