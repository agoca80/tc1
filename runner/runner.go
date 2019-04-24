package runner

// Runner ...
type Runner chan struct{}

// New ...
func New() Runner {
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
