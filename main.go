package main

import (
	"flag"

	"github.com/agoca80/tc1/client"
)

var (
	server = flag.Bool("server", false, "Run the server instead of the client")
)

const (
	in      = "/tmp/input"
	out     = "/tmp/output"
	dump    = "/tmp/dump"
	clients = 5
	reports = 100
	size    = 1000000000
)

func init() {
	flag.Parse()
}

func main() {
	switch *server {
	case true:
		Server()
	case false:
		client.Start()
	}
}
