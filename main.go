package main

import (
	"fmt"
	"os"

	"github.com/agoca80/tc1/client"
	"github.com/agoca80/tc1/server"
)

// Server wrapper
func Server() {
	if os.Getenv("testing") == "true" {
		server.Testing = true
	}
	server.Start()
}

// Client wrapper
func Client() { client.Start() }

// The main function only has 2 inputs:
// - The first argument
func main() {
	switch os.Args[1] {
	case "server":
		Server()
	case "client":
		Client()
	default:
		fmt.Println("Read the README file for instructions")
		fmt.Println()
	}
}
