package filter

import "testing"

func TestFilter_Run(t *testing.T) {
	var (
		check = struct {
			input  []int
			output []int
		}{
			[]int{0, 0, 0, 0, 0, 1, 1, 0, 2, 1, 0, 3, 3, 3, 3, 3, 1, 4, 5},
			[]int{0, 1, 2, 3, 4, 5},
		}

		filter = New("/tmp/in.test", 8)
		in     = make(chan int)
		out    = make(chan int)
	)

	go filter.Run(in, out)
	go func() {
		for _, number := range check.input {
			in <- number
		}
		close(in)
	}()

	output := []int{}
	for number := range out {
		output = append(output, number)
	}

	if len(output) != len(check.output) {
		t.Fail()
	} else {
		for i := range output {
			if output[i] != check.output[i] {
				t.Fail()
			}
		}
	}
}
