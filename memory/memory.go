package memory

type memory struct {
	Statistics
	Bitmap
}

func (m *memory) Remembers(n int) (ok bool) {
	if ok = m.Bitmap.Remembers(n); ok {
		m.Duplicates++
	} else {
		m.Uniques++
	}
	return
}
