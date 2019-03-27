package main

import (
	"testing"

	"github.com/agoca80/tc1/server"
)

func TestServer(t *testing.T) {
	server.Testing = true
	Server()
}
