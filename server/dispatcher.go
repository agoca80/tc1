package server

import "io"

func (s *Service) dispatcher(clients chan<- io.ReadCloser) {
	defer close(clients)
	for s.Running() {
		client, err := s.Accept()
		switch {

		case err != nil && err.Error() == "use of closed network connection":
			return

		case err == nil:
			clients <- client

		}
	}
}
