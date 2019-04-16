package main

import (
	"net"
	"os"

	"github.com/agoca80/tc1/memory"

	"github.com/agoca80/tc1/client"
	"github.com/agoca80/tc1/server"
)

const (
	in   = "/tmp/input"
	out  = "/tmp/output"
	dump = "/tmp/dump"
)

// Server ...
func Server() {
	memory, err := memory.NewMemory(dump)
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	input, err := os.OpenFile(in, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	service := server.New(listener, input, output, memory)
	service.Start()

	err = service.Store(dump)
	if err != nil {
		panic(err)
	}
}

// Client ...
func Client() { client.Start() }

func main() {
	if len(os.Args) < 2 || os.Args[1] == "server" {
		Server()
	} else {
		Client()
	}
}
