package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	r := require.New(t)

	pwds, err := ParsePasswordList([]string{
		`1-3 a: abcde`,
		`1-3 b: cdefg`,
		`2-9 c: ccccccccc`,
	})
	r.NoError(err)
	r.Len(pwds, 3)

	r.Equal(2, CountValidPasswords(pwds))
}
