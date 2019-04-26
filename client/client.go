package client

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/agoca80/tc1/runner"
)

// PanicOnErr if there is an error during send
var (
	Address    = "127.0.0.1:4000"
	PanicOnErr = false
)

func send(conn io.Writer, msg []byte) (err error) {
	_, err = conn.Write(msg)
	return
}

// Send ...
func Send(conn io.Writer, msg []byte) (err error) {
	_, err = conn.Write(msg)
	return
}

// Terminate sends the termination sequence to the service
func Terminate() (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		return
	}
	return send(conn, []byte("terminate\n"))
}

// Taste ...
func taste() {
	for {
		if conn, err := net.Dial("tcp", Address); err == nil {
			fmt.Fprintln(os.Stderr, "tasted")
			conn.Close()
			break
		}
	}
}

// minion ...
func minion(during, wait int, leader runner.Runner, generator Generator) {
	conn, err := net.Dial("tcp", Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for leader.Running() {
		Send(conn, generator.Number())
		if wait > 0 {
			time.Sleep(time.Duration(wait) * time.Microsecond)
		}
	}
}

// Start ...
func Start(during, wait int, size int) {
	platoon := runner.New()
	generator := Random(size)

	taste()
	for i := 0; i < 5; i++ {
		go minion(during, wait, platoon, generator)
	}

	time.Sleep(time.Duration(during) * time.Second)
	platoon.Stop()
	Terminate()
}
