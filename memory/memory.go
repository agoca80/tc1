package memory

// Inspiration taken from
// https://github.com/ShawnMilo/bitmap
// https://github.com/adonovan/gopl.io/blob/master/ch6/intset/intset.go

// Size ...
const size = 1000000000

// Memory ...
type memory []uint64

func newMemory() memory {
	return make(memory, size)
}

func (m memory) Remembers(n int) (ok bool) {
	w, b := n/64, uint(n%64)
	ok = m[w]&(1<<b) != 0
	m[w] |= 1 << b
	return
}
