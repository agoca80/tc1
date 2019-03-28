package server

// Memory ...
type Memory struct {
	memory
}

const size = 1000000000

type memory []uint64

func newMemory() *Memory {
	return &Memory{
		memory: make([]uint64, size/64+1),
	}
}

// no slice boundaries performed.
// let it panic if they are exceeded
func (m memory) Remembers(n int) (ok bool) {
	w, b := n/64, uint(n%64)
	ok = m[w]&(1<<b) != 0
	m[w] |= 1 << b
	return
}
