package memory

// Statistics ...
type Statistics struct {
	Uniques    int
	Duplicates int
}

// Stats ...
func (m *memory) Stats() Statistics {
	return m.Statistics
}
