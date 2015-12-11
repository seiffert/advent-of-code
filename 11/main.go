package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ASCIIa = 97
	ASCIIi = 105
	ASCIIl = 108
	ASCIIo = 111
	ASCIIz = 122
)

func main() {
	for _, w := range os.Args[1:] {
		pwd := generateNextPassword(w)
		fmt.Printf("Next password after %s should be: %s\n", w, pwd)

		pwd = generateNextPassword(pwd)
		fmt.Printf("After that, the next should be: %s\n", pwd)
	}
}

func generateNextPassword(w string) string {
	for {
		w = increment(w)
		if isValid(w) {
			return w
		}
	}
}

func increment(w string) string {
	for i := 0; i < len(w); i++ {
		c := w[i]
		if c == ASCIIi || c == ASCIIl || c == ASCIIo {
			return w[:i] + string(c+1) + strings.Repeat("a", len(w[i+1:]))
		}
	}
	for i := len(w) - 1; i >= 0; i-- {
		l := w[i]
		if l != ASCIIz {
			w = w[:i] + string(l+1) + w[i+1:]
			break
		} else {
			w = w[:i] + string(ASCIIa) + w[i+1:]
		}
	}

	return w
}

func isValid(w string) bool {
	var (
		hasStraight   bool
		prev          rune
		inPair        bool
		pairChar      rune
		pairs         int
		straightCount = 1
	)

	for _, c := range w {
		inPair = (prev == c && !inPair && pairChar != c)
		if inPair {
			pairs++
			pairChar = c
		}

		if c == prev+1 {
			straightCount++
			if straightCount == 3 {
				hasStraight = true
			}
		} else {
			straightCount = 1
		}
		prev = c
	}

	return hasStraight && pairs > 1
}
