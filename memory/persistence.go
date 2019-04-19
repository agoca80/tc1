package memory

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"os"
)

// Dump ...
func (m *memory) Dump(dump string) (err error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err = encoder.Encode(m)
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

func (m *memory) Load(dump string) (err error) {
	_, err = os.Stat(dump)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return
	}

	file, err := os.OpenFile(dump, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return
	}

	decoder := gob.NewDecoder(gz)
	err = decoder.Decode(m)
	if err != nil {
		return
	}

	return
}
