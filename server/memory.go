package server

// Inspiration taken from
// https://github.com/ShawnMilo/bitmap
// https://github.com/adonovan/gopl.io/blob/master/ch6/intset/intset.go

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
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

func (m Memory) store() {
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(j)
	if err != nil {
		panic(err)
	}

	err = gz.Close()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dump, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
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

	buf := bytes.NewBuffer(load)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		panic(err)
	}

	inner, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(inner, &s.Memory)
	if err != nil {
		panic(err)
	}

	for _, word := range s.Memory {
		s.Total += bits.OnesCount64(word)
	}
	fmt.Println(s)
}
