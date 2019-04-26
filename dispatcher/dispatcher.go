package dispatcher

import (
	"io"
	"net"

	"github.com/agoca80/tc1/runner"
)

type dispatcher struct {
	net.Listener
	runner.Runner
}

func newDispatcher(service runner.Runner) *dispatcher {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	return &dispatcher{
		Listener: listener,
		Runner:   service,
	}
}

func (d *dispatcher) Dispatch(clients chan<- io.ReadCloser) {
	for d.Running() {
		client, err := d.Accept()
		switch {

		case err != nil && err.Error() == "use of closed network connection":
			return

		case err == nil:
			clients <- client

		}
	}
	close(clients)
}
