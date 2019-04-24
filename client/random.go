package client

import (
	"fmt"
	"math/rand"
)

type random int

// Random ...
func Random(size int) Generator {
	return random(size)
}

func (r random) Number() []byte {
	return []byte(fmt.Sprintf("%09d\n", rand.Intn(int(r))))
}
