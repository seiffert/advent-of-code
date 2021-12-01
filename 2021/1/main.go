package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	in := strings.Split(lib.MustReadFile("input.txt"), "\n")

	fmt.Printf("%d larger measures found\n", CountIncreases(in))
	fmt.Printf("%d larger windows of 3 found\n", CountWindowIncreases(in, 3))
}

func CountIncreases(in []string) (out int) {
	return CountWindowIncreases(in, 1)
}

func CountWindowIncreases(in []string, ws int) (out int) {
	for i := 0; i < len(in)-ws; i++ {
		if sum(in[i:i+ws]) < sum(in[i+1:i+ws+1]) {
			out++
		}
	}
	return
}

func sum(in []string) (out int) {
	for _, i := range in {
		out += lib.MustInt(i)
	}
	return
}
