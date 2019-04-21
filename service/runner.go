package service

type runner chan struct{}

func newRunner() runner {
	return make(runner)
}

func (r runner) Running() bool {
	select {
	case <-r:
		return false
	default:
		return true
	}
}

func (r runner) Stop() {
	if r.Running() {
		close(r)
	}
}
