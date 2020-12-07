package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNIntsWithSum(t *testing.T) {
	var (
		r     = require.New(t)
		input = `1721
979
366
299
675
1456
`
	)

	list, err := parseIntList(input)
	r.NoError(err)

	res2, err := FindNIntsWithSum(2, 2020, list)
	r.NoError(err)
	r.ElementsMatch([]int{1721, 299}, res2)

	res3, err := FindNIntsWithSum(3, 2020, list)
	r.NoError(err)
	r.ElementsMatch([]int{979, 366, 675}, res3)
}
