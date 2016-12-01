package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	zeroes, _ := strconv.Atoi(os.Args[2])

	var i int
	for {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), strings.Repeat("0", zeroes)) {
			fmt.Printf("Result: %d\n", i)
			return
		}
		i++
	}
}
