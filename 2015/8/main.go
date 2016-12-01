package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int
	var doubleSum int
	for scanner.Scan() {
		s := scanner.Text()
		u, err := strconv.Unquote(s)
		if err != nil {
			log.Fatalf("Error unquoting %s: %s", s, err)
		}

		d := strconv.Quote(s)

		fmt.Printf("len(%s)=%d - len(%s)=%d = %d\n", s, len(s), u, len(u), len(s)-len(u))
		fmt.Printf("len(%s)=%d - len(%s)=%d = %d\n", d, len(d), s, len(s), len(d)-len(s))
		sum += len(s) - len(u)
		doubleSum += len(d) - len(s)
	}

	fmt.Printf("Difference: %d\n", sum)
	fmt.Printf("Double Enquoted Difference: %d\n", doubleSum)
}
