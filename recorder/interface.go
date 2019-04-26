package recorder

import "os"

// Recorder ...
type Recorder interface {
	Record(<-chan int)
}

// New ...
func New(file string) (r Recorder, err error) {
	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	r = &recorder{
		WriteCloser: output,
	}

	return
}
