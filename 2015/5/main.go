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

	var countVocals, countDoubles int
	for reader.Len() > 0 {
		r, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("error reading rune: %s", err)
		}

		if strings.Contains(vocals, string(r)) {
			countVocals++
		}
		if reader.Len() == 0 {
			continue
		}

		next, err := reader.ReadByte()
		if err != nil {
			log.Fatalf("error reading rune: %s", err)
		}
		if err := reader.UnreadByte(); err != nil {
			log.Fatalf("error reading rune: %s", err)
		}

		if string(next) == string(r) {
			countDoubles++
		}
		x := string(r) + string(next)
		if x == "ab" || x == "cd" || x == "pq" || x == "xy" {
			return false
		}
	}

	return countVocals > 2 && countDoubles > 0
}
