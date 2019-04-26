package process

import "io"

// Pool ...
type Pool interface {
	Process(<-chan io.ReadCloser, chan<- int)
}
