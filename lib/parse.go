package lib

import "strconv"

func MustInt(in string) int {
	n, err := strconv.Atoi(in)
	if err != nil {
		Abort("failed to parse int %q: %v", in, err)
	}
	return n
}
