package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	result := CountValidMessages(lib.MustReadFile("input.txt"))
	fmt.Printf("result 1: %d\n", result)

	result = CountValidMessages(lib.MustReadFile("input2.txt"))
	fmt.Printf("result 2: %d\n", result)
}

var cache = make(map[string]map[string]bool)

func CountValidMessages(input string) int {
	c := strings.Split(input, "\n\n")

	rules := map[int]string{}
	for _, rule := range strings.Split(c[0], "\n") {
		id, pattern := parseRule(rule)
		rules[id] = pattern
	}

	var validMessages int
	for _, msg := range strings.Split(c[1], "\n") {
		if isValid(msg, rules, rules[0]) {
			validMessages++
		}
	}

	cache = make(map[string]map[string]bool)

	return validMessages
}

func parseRule(input string) (int, string) {
	parts := strings.Split(input, ": ")
	id, _ := strconv.Atoi(parts[0])
	return id, parts[1]
}

func isValid(msg string, rules map[int]string, pattern string) bool {
	if pc, ok := cache[pattern]; ok {
		if mc, ok := pc[msg]; ok {
			return mc
		}
	} else {
		cache[pattern] = make(map[string]bool)
	}

	// "a"
	if strings.HasPrefix(pattern, `"`) {
		result := strings.Contains(pattern, msg)
		cache[pattern][msg] = result
		return result
	}

	// 1 | 2 | 3
	if strings.Contains(pattern, `|`) {
		for _, branch := range strings.Split(pattern, " | ") {
			if isValid(msg, rules, branch) {
				cache[pattern][msg] = true
				return true
			}
		}
		cache[pattern][msg] = false
		return false
	}

	// 1 2 3
	parts := strings.Split(pattern, " ")
	ruleID, _ := strconv.Atoi(parts[0])
	suffixPattern := strings.Join(parts[1:], " ")

	for i := 1; i <= len(msg); i++ {
		prefix, suffix := msg[:i], msg[i:]
		switch {
		case suffix == "" && suffixPattern == "":
			if isValid(prefix, rules, rules[ruleID]) {
				cache[pattern][msg] = true
				return true
			}
		case suffix != "" && suffixPattern != "":
			if isValid(prefix, rules, rules[ruleID]) && isValid(suffix, rules, suffixPattern) {
				cache[pattern][msg] = true
				return true
			}
		case suffix == "" && suffixPattern != "":
			cache[pattern][msg] = false
			return false
		}
	}

	cache[pattern][msg] = false
	return false
}
