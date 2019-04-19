package memory

import "testing"

func TestBitmap_Remembers(t *testing.T) {
	bitmap := newBitmap(16)

	checks := []struct {
		num int
		ok  bool
	}{
		{0, false},
		{0, true},
		{1, false},
		{0, true},
		{2, false},
		{1, true},
		{2, true},
	}

	for _, check := range checks {
		if bitmap.Remembers(check.num) != check.ok {
			t.Fail()
		}
	}
}
