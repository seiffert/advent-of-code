package lib

import (
	"io/ioutil"
)

func MustReadFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		Abort("failed to read input file: %v", err)
	}

	return data
}
