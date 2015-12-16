package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	LineRegexp = regexp.MustCompile("^Sue ([0-9]+): (.*)$")
	props      = map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}
)

func main() {
	var mostCommonProps int
	var answer string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var commonProps int
		l := s.Text()

		m := LineRegexp.FindAllStringSubmatch(l, -1)
		num := m[0][1]
		propStrings := strings.Split(m[0][2], ", ")
		for _, p := range propStrings {
			pp := strings.Split(p, ": ")
			switch pp[0] {
			case "cats":
				fallthrough
			case "trees":
				if i(props[pp[0]]) < i(pp[1]) {
					commonProps++
				}
			case "pomeranians":
				fallthrough
			case "goldfish":
				if i(props[pp[0]]) > i(pp[1]) {
					commonProps++
				}
			default:
				if props[pp[0]] == pp[1] {
					commonProps++
				}
			}
		}

		if commonProps > mostCommonProps {
			mostCommonProps = commonProps
			answer = num

			log.Printf("New mcpa: %s with %d common props", num, commonProps)
		}
	}

	log.Printf("Answer: %s", answer)
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
