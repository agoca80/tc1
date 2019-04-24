package main

import (
	"os"

	"github.com/agoca80/tc1/service"
)

// Server ...
func Server() {
	output, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	srv, err := service.New(clients, reports, *size, in, output)
	if err != nil {
		panic(err)
	}

	srv.Start()
	if err != nil {
		panic(err)
	}
}
