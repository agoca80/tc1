package memory

// Memory ...
type Memory interface {
	Remembers(int) bool
}

// Stats ...
type Stats interface {
}

// Persistent ...
type Persistent interface {
	Memory
	Stats
	Load(dump string) error
	Store(dump string) error
}

// NewMemory ...
func NewMemory(dump string) (Persistent, error) {
	memory := &persistent{
		memory: newMemory(),
	}

	err := memory.Load(dump)
	return memory, err
}
