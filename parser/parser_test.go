package parser

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	type check struct {
		stream string
		kind   []string
		number []int
	}

	parser := func(stream string) *Parser {
		return New(bytes.NewReader([]byte(stream)))
	}

	checks := []check{
		check{
			"\nignored",
			[]string{Invalid},
			[]int{},
		},
		check{
			"terminate\nignored",
			[]string{Terminate},
			[]int{},
		},
		check{
			"000000000\nterminate\n0000000001\n",
			[]string{Number, Terminate},
			[]int{0},
		},
		check{
			"000000000\n999999999\nterminate\n0000000001\n",
			[]string{Number, Number, Terminate},
			[]int{0, 999999999},
		},
	}

	fail := func(c check) {
		fmt.Printf("FAIL %+v\n", c)
		t.Fail()
	}

	for _, c := range checks {
		p := parser(c.stream)
		for i, kind := 0, p.Next(); true; i, kind = i+1, p.Next() {
			if kind != c.kind[i] {
				fail(c)
				break
			}

			if kind == Terminate || kind == Invalid {
				break
			}

			if kind != Number {
				continue
			}

			if p.Number != c.number[i] {
				fail(c)
			}
		}
	}
}
