package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

// Testing enable input recollection for inspection on the filter
var Testing = false

type service struct {
	// This will broadcast the termination signal to all goroutines
	terminate chan bool

	// These are updated from two places:
	//
	// - The service String method, called from the reporter,
	//   setting Duplicates and Uniques to 0 before returning
	//
	// - The service filter method, increasing the Duplicates and Uniques count
	//
	// These method are executed from different goroutines
	// It will be eventually consistent
	// when the filter is stopped and prints the actual value
	Duplicates int
	Uniques    int
	// This is only updated from the filter, and read from
	// the service String method.
	Total int

	// Pipeline
	//
	// dispacher <-listener conns<-
	// workers   <-conns    numbers<-
	// filter    <-numbers  uniques<- (numbers save to input file if Testing is true)
	// service   <-uniques  writer<-
	//
	net.Listener
	// The listener will be closed when the broadcast signal arrives
	// conns will be dispatched from the dispatcher to the workers
	// as long as the termination signal has not been broadcasted.
	conns chan net.Conn //
	// workers will process connections while they are valid and
	// the termination signal has not been broadcasted.
	// conns will be serialized to numbers and sent to the filter
	numbers chan int
	input   io.Writer
	// numbers are provided by the workers to the filter.
	// If the number was seen for the first time,
	// it will be send back to the service through uniques,
	// and optionally to input if Testing is true
	uniques chan int
	// The service will read numbers from here and write them to io.Writer
	io.Writer

	// Numbers the service has already seen.
	// All access are serialized through the pipeline
	*Memory

	// pool of workers
	// This keeps track of the workers which have finished.
	// Once the group in completely finished, the waiter will close
	// the numbers channel, to propagate the channel closures through the pipeline
	sync.WaitGroup
}

func newService() *service {
	writer, err := os.OpenFile("numbers.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	input, err := os.OpenFile("input", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	return &service{
		conns:     make(chan net.Conn),
		Listener:  listener,
		Memory:    newMemory(),
		numbers:   make(chan int, 5*1024),
		terminate: make(chan bool),
		uniques:   make(chan int),

		Writer: writer,
		input:  input,
	}
}

// Start ...
func Start() {
	s := newService()

	go s.reporter()

	// Create the workers pool
	for i := 0; i < 5; i++ {
		s.Add(1)
		go s.worker(s.numbers)
	}

	go s.filter(s.numbers, s.uniques)

	// waiter goroutine.
	// It will wait from the pool of workers to finish
	// to close numbers and propagate the channels closures
	// to the next stages in the pipeline
	go func() {
		s.Wait()
		close(s.numbers)
	}()

	// Last step in the pipeline
	for unique := range (<-chan int)(s.uniques) {
		fmt.Fprintln(s, unique)
	}

}

// terminating poll the terminate broadcast channel
// if the channel is not ready, the service is still running
// and the terminate signal has not arrived yet.
func (s *service) terminating() bool {
	select {
	case <-s.terminate:
		return true
	default:
		return false
	}
}

// String for representing the service during reports
func (s *service) String() string {
	return fmt.Sprintf("Received %v unique numbers, %v duplicates. Unique total: %v", s.Uniques, s.Duplicates, s.Total)
}
