package main

import (
	"os"

	"github.com/agoca80/tc1/client"
)

const (
	in   = "/tmp/input"
	out  = "/tmp/output"
	dump = "/tmp/dump"
)

// Client ...
func Client() { client.Start() }

func main() {
	if len(os.Args) < 2 || os.Args[1] == "server" {
		Server()
	} else {
		Client()
	}
}
