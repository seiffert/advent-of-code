package lib

import "strconv"

func MustInt(in string) int {
	n, err := strconv.Atoi(in)
	if err != nil {
		Abort("failed to parse int %q: %v", in, err)
	}
	return n
}

func MustAllInts(in []string) (out []int) {
	for _, i := range in {
		out = append(out, MustInt(i))
	}
	return
}
