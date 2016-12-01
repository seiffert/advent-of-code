package main

import (
	"fmt"
	"os"
)

func main() {
	count := 29000000
	for {
		for i := 1; true; i++ {
			for j := 1; j <= i; j++ {
				if i%j == 0 {
					count -= j * 10
				}
			}
			fmt.Printf("house %d: %d left\n", i, count)
			if count <= 0 {
				os.Exit(0)
			}
		}
	}
}
