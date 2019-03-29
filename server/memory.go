package server

// Inspiration taken from
// https://github.com/ShawnMilo/bitmap
// https://github.com/adonovan/gopl.io/blob/master/ch6/intset/intset.go

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/bits"
	"os"
)

// Memory ...
type Memory []uint64

const size = 1000000000

// no slice boundaries performed.
// let it panic if they are exceeded
func (m Memory) remembers(n int) (ok bool) {
	w, b := n/64, uint(n%64)
	ok = m[w]&(1<<b) != 0
	m[w] |= 1 << b
	return
}

const dump = "dump.gz"

func (s *service) store() {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, &s.Memory)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(dump, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	gz := gzip.NewWriter(file)
	_, err = gz.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	gz.Close()
	file.Close()
}

func (s *service) remind() {
	s.Memory = make(Memory, size/64+1)
	_, err := os.Stat(dump)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}

	load, err := ioutil.ReadFile(dump)
	if err != nil {
		panic(err)
	}

	gz, err := gzip.NewReader(bytes.NewBuffer(load))
	if err != nil {
		panic(err)
	}

	err = binary.Read(gz, binary.BigEndian, &s.Memory)
	if err != nil {
		panic(err)
	}

	for _, word := range s.Memory {
		s.Total += bits.OnesCount64(word)
	}
	fmt.Println(s)
}
