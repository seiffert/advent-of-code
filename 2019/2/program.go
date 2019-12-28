package main

import (
	"strconv"
	"strings"
)

func Parse(in string) []int {
	vs := strings.Split(in, ",")
	p := make([]int, 0, len(vs))

	for _, v := range vs {
		i, err := strconv.Atoi(v)
		if err != nil {
			i = -1
		}
		p = append(p, i)
	}
	return p
}
