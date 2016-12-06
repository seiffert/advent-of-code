package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	fmt.Println(ErrorCorrect(lines))
}

func ErrorCorrect(lines []string) string {
	c := NewErrorCorrector()
	for _, line := range lines {
		for i, char := range line {
			c.Add(i, byte(char))
		}
	}
	return c.Correct()
}

func NewErrorCorrector() *ErrorCorrector {
	return &ErrorCorrector{}
}

type ErrorCorrector struct {
	chars []map[byte]int
}

func (ec *ErrorCorrector) Add(position int, char byte) {
	if len(ec.chars) <= position {
		ec.chars = append(ec.chars, map[byte]int{})
	}
	ec.chars[position][char] += 1
}

func (ec *ErrorCorrector) Correct() string {
	var result []byte
	for i, c := range ec.chars {
		result = append(result, byte(0))
		numMin := 1000
		for c, num := range c {
			if num < numMin {
				result[i] = c
				numMin = num
			}
		}
	}
	return string(result[0:])
}
