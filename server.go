package main

import (
	"net"
	"os"

	"github.com/agoca80/tc1/memory"
	"github.com/agoca80/tc1/service"
)

// Server ...
func Server() {
	memory := memory.NewMemory(size)

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

	srv := service.New(clients, reports, size, listener, input, output, memory)
	srv.Start()

	err = srv.Memory.Dump(dump)
	if err != nil {
		panic(err)
	}
}
