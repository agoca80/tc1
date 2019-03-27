package client

import "fmt"

type constant struct {
	msg []byte
}

// Constant generates always the same number
func Constant(n int) Generator {
	return &constant{
		msg: []byte(fmt.Sprintf("%09d\n", n)),
	}
}

func (c *constant) Number() []byte {
	return c.msg
}
