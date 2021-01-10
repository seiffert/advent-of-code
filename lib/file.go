package lib

import (
	"bytes"
	"io/ioutil"
)

func MustReadFile(name string) string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		Abort("failed to read input file: %v", err)
	}

	return string(bytes.TrimSpace(data))
}
