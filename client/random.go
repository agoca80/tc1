package client

import (
	"fmt"
	"math/rand"
)

type random struct{}

// Random ...
func Random() Generator {
	return &random{}
}

func (r *random) Number() []byte {
	return []byte(fmt.Sprintf("%09d\n", rand.Intn(1000000000)))
}
