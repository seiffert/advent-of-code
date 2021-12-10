package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("total syntax error score: %d\n", TotalSyntaxErrorScore(lib.MustReadFile("input.txt")))
	fmt.Printf("middle completion score: %d\n", MiddleCompletionScore(lib.MustReadFile("input.txt")))
}

type syntaxError struct {
	IllegalCharacter rune
}

func (se syntaxError) Error() string {
	return fmt.Sprintf("syntax error: illegal character \"%c\"", se.IllegalCharacter)
}

func TotalSyntaxErrorScore(in string) (out int) {
	scoreByChar := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	for _, l := range strings.Split(in, "\n") {
		if _, err := complete(l); err != nil {
			if serr, ok := err.(syntaxError); err != nil && ok {
				out += scoreByChar[serr.IllegalCharacter]
			}
		}
	}
	return
}

func MiddleCompletionScore(in string) (out int) {
	var scores []int
	scoreByChar := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	for _, l := range strings.Split(in, "\n") {
		compl, err := complete(l)
		if err == nil {
			var score int
			for _, c := range compl {
				score = score*5 + scoreByChar[c]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func complete(l string) (compl string, err error) {
	var open []rune
	closings := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	for _, c := range l {
		if _, ok := closings[c]; ok {
			open = append([]rune{c}, open...)
		} else if closings[open[0]] == c {
			open = open[1:]
		} else {
			return "", syntaxError{c}
		}
	}
	for _, c := range open {
		compl += string(closings[c])
	}
	return
}
