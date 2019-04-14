package main

import (
	"net"
	"os"

	"github.com/agoca80/tc1/client"
	"github.com/agoca80/tc1/server"
)

// Server ...
func Server() {
	output, err := os.OpenFile("numbers.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	input, err := os.OpenFile("input", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	service := server.New(listener, input, output)
	service.Start()
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
