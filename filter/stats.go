package filter

// Statistics ...
type Statistics struct {
	Uniques    int
	Duplicates int
}

// Stats ...
func (f *filter) Stats() Statistics {
	return f.Statistics
}
