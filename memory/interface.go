package memory

// Interface ...
type Interface interface {
	Remembers(int) bool
	Load(string) error
	Dump(string) error
	Stats() Statistics
}

// NewMemory ...
func NewMemory(size int) Interface {
	return &memory{
		Bitmap: newBitmap(size),
	}
}
