package memory

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
	"os"
)

type persistent struct {
	memory
}

// Store ...
func (p persistent) Store(dump string) (err error) {
	buffer := new(bytes.Buffer)
	err = binary.Write(buffer, binary.BigEndian, p.memory)
	if err != nil {
		return
	}

	file, err := os.OpenFile(dump, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	gz := gzip.NewWriter(file)
	_, err = gz.Write(buffer.Bytes())
	if err != nil {
		return
	}
	defer gz.Close()

	return
}

func (p persistent) Load(dump string) (err error) {
	_, err = os.Stat(dump)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return
	}

	load, err := ioutil.ReadFile(dump)
	if err != nil {
		return
	}

	gz, err := gzip.NewReader(bytes.NewBuffer(load))
	if err != nil {
		return
	}

	err = binary.Read(gz, binary.BigEndian, &p.memory)
	if err != nil {
		return
	}

	return
}
