package client

import (
	"fmt"
	"math/rand"
)

type randomp struct {
	size   int
	number [][]byte
}

// Randomp ...
func Randomp(pool int) Generator {
	r := &randomp{
		size:   pool,
		number: make([][]byte, pool),
	}
	for i := 0; i < pool; i++ {
		r.number[i] = []byte(fmt.Sprintf("%09d\n", rand.Intn(1000000000)))
	}
	return r
}

func (r *randomp) Number() []byte {
	return r.number[rand.Intn(r.size)]
}
