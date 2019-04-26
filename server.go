package main

import (
	"github.com/agoca80/tc1/service"
)

// Server ...
func Server() {
	srv, err := service.New(clients, reports, *size, in, out)
	if err != nil {
		panic(err)
	}

	srv.Start()
	if err != nil {
		panic(err)
	}
}
