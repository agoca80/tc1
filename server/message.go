package server

import (
	"bufio"
	"strconv"
)

const (
	invalid   = -1
	terminate = 1000000000
)

// newNumber process the input from a bufio.Reader
// scaning one line at a time
//
// if there is an issue reading from the buffer
// this will cause the calling function (process) to deem the connection invalid
// and it will close it
//
// We will use 2 numbers out of the range to signal invalid and termination messages
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
