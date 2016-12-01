package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := os.Args[1]
	num, _ := strconv.Atoi(os.Args[2])

	var out string
	for i := 0; i < num; i++ {
		out = ""
		var curChar rune
		var count int
		for _, c := range in {
			if curChar == 0 || c == curChar {
				curChar = c
				count++
			} else {
				out += fmt.Sprintf("%d%s", count, string(curChar))
				count = 1
				curChar = c
			}
		}
		out += fmt.Sprintf("%d%s", count, string(curChar))
		in = out
	}

	fmt.Printf("length: %d", len(out))
}
