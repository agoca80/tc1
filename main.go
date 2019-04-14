package main

import (
	"os"

	"github.com/agoca80/tc1/client"
	"github.com/agoca80/tc1/server"
)

// Server ...
func Server() {
	if os.Getenv("testing") == "true" {
		server.Testing = true
	}
	server.Start()
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
