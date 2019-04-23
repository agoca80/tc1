package service

// Runner ...
type Runner chan struct{}

// NewRunner ...
func NewRunner() Runner {
	return make(Runner)
}

// Running ...
func (r Runner) Running() bool {
	select {
	case <-r:
		return false
	default:
		return true
	}
}

// Stop ...
func (r Runner) Stop() {
	if r.Running() {
		close(r)
	}
}
