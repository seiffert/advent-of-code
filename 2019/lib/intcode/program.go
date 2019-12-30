package intcode

import (
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func Parse(in string) []int {
	vs := strings.Split(in, ",")
	p := make([]int, 0, len(vs))

	for _, v := range vs {
		i, err := strconv.Atoi(v)
		if err != nil {
			lib.Abort("invalid intcode instruction: %w", err)
		}
		p = append(p, i)
	}
	return p
}
