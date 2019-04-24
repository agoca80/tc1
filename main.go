package main

import (
	"flag"

	"github.com/agoca80/tc1/client"
)

var (
	during = flag.Int("t", 1, "Send termination signal after t seconds")
	wait   = flag.Int("w", 1, "microseconds between client numbers")
	size   = flag.Int("r", 1000000, "range of numbers")
)

const (
	in      = "/tmp/input"
	out     = "/tmp/output"
	dump    = "/tmp/dump"
	clients = 5
	reports = 100
)

func init() {
	flag.Parse()
}

func main() {
	go client.Start(*during, *wait, *size)

	Server()
}
