package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestCountPaths(t *testing.T) {
	r := require.New(t)

	r.NoError(filepath.Walk("samples/", func(path string, info os.FileInfo, err error) error {
		t.Run(path, func(t *testing.T) {
			r = require.New(t)

			r.NoError(err)
			if info.IsDir() {
				return
			}

			s := strings.Split(filepath.Base(path), ".")
			r.Equal(lib.MustInt(s[0]), CountPaths(lib.MustReadFile(path), false))
			r.Equal(lib.MustInt(s[1]), CountPaths(lib.MustReadFile(path), true))
		})

		return nil
	}))

}
