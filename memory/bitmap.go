package memory

// Bitmap ...
type Bitmap []uint64

func newBitmap(size int) Bitmap {
	return make(Bitmap, size/64+1)
}

// Remembers ...
func (b Bitmap) Remembers(n int) (ok bool) {
	w, m := n/64, uint(n%64)
	ok = b[w]&(1<<m) != 0
	b[w] |= 1 << m
	return
}
