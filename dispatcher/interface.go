package dispatcher

import (
	"io"

	"github.com/agoca80/tc1/runner"
)

// Dispatcher ...
type Dispatcher interface {
	Dispatch(chan<- io.ReadCloser)
	io.Closer
}

// New ...
func New(service runner.Runner) Dispatcher {
	return newDispatcher(service)
}
