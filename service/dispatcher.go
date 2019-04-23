package service

import (
	"io"
	"net"
)

type dispatcher struct {
	net.Listener
	Runner
}

func newDispatcher(leader Runner) *dispatcher {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	return &dispatcher{
		Listener: listener,
		Runner:   leader,
	}
}

func (d *dispatcher) run(clients chan<- io.ReadCloser) {
	defer close(clients)
	for d.Running() {
		client, err := d.Accept()
		switch {

		case err != nil && err.Error() == "use of closed network connection":
			return

		case err == nil:
			clients <- client

		}
	}
}
