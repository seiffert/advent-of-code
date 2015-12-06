package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	vocals = "aeiou"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nice, naughty int
	for scanner.Scan() {
		word := scanner.Text()
		if isNice(word) {
			nice++
		} else {
			naughty++
		}
	}

	fmt.Printf("Nice: %d, Naughty: %d\n", nice, naughty)
}

func isNice(word string) bool {
	reader := strings.NewReader(word)

	var aba, doubles int
	for reader.Len() > 2 {
		r, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("error reading rune: %s", err)
		}

		next, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("error reading rune: %s", err)
		}
		next2, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("error reading rune: %s", err)
		}

		if err := reader.UnreadByte(); err != nil {
			log.Fatalf("error reading rune: %s", err)
		}
		if err := reader.UnreadByte(); err != nil {
			log.Fatalf("error reading rune: %s", err)
		}

		if strings.Count(word, string(r)+string(next)) >= 2 {
			doubles++
		}

		if string(r) == string(next2) {
			aba++
		}
	}

	return doubles > 0 && aba > 0
}
