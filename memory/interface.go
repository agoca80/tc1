package memory

// Memory ...
type Memory interface {
	Remembers(int) bool
}

// NewMemory ...
func NewMemory(size int) Memory {
	return newBitmap(size)
}
