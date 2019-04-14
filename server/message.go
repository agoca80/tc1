package server

import (
	"bufio"
	"strconv"
)

const (
	invalid   = -1
	terminate = 1000000000
)

func newNumber(reader *bufio.Reader) (int, error) {
	message, err := reader.ReadString('\n')
	if err != nil || len(message) != 10 {
		return invalid, err
	}

	if message == "terminate\n" {
		return terminate, nil
	}

	number, err := strconv.Atoi(message[0:9])
	if err != nil {
		return invalid, nil
	}

	return number, nil
}
