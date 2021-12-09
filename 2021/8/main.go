package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("number of unique #digit numbers: %d\n", CountUniqueDigitNumbers(lib.MustReadFile("input.txt")))
	fmt.Printf("sum of output values: %d\n", SumOutputValues(lib.MustReadFile("input.txt")))
}

func CountUniqueDigitNumbers(in string) (out int) {
	for _, l := range strings.Split(in, "\n") {
		t := strings.Split(l, " | ")
		for _, n := range strings.Split(t[1], " ") {
			l := len(n)
			if l >= 2 && l <= 4 || l == 7 {
				out++
			}
		}
	}
	return
}

func SumOutputValues(in string) (out int) {
	for _, l := range strings.Split(in, "\n") {
		t := strings.Split(l, " | ")

		outputNums := canonicalCodes(strings.Split(t[1], " "))
		unsolvedCodes := canonicalCodes(strings.Split(t[0], " "))

		digitByCode, codeByDigit := make(map[string]int), make(map[int]string)
		for len(codeByDigit) != 10 {
			for i, c := range unsolvedCodes {
				digit := -1
				switch len(c) {
				case 2:
					digit = 1
				case 4:
					digit = 4
				case 3:
					digit = 7
				case 7:
					digit = 8
				case 6:
					switch {
					case codeByDigit[4] != "" && sublen(codeByDigit[4], c) == 0:
						digit = 9
					case codeByDigit[9] != "" && codeByDigit[1] != "" && sublen(codeByDigit[1], c) == 0:
						digit = 0
					case codeByDigit[1] != "" && sublen(codeByDigit[1], c) != 0:
						digit = 6
					default:
						continue
					}
				case 5:
					fmt.Printf("c-4: %d\n", sublen(c, codeByDigit[4]))
					switch {
					case codeByDigit[1] != "" && sublen(codeByDigit[1], c) == 0:
						digit = 3
					case codeByDigit[3] != "" && codeByDigit[4] != "" && sublen(c, codeByDigit[4]) == 3:
						digit = 2
					case codeByDigit[2] != "" && codeByDigit[3] != "":
						digit = 5
					default:
						continue
					}
				default:
					lib.Abort("invalid number of signals")
				}

				digitByCode[c] = digit
				codeByDigit[digit] = c
				unsolvedCodes = append(unsolvedCodes[:i], unsolvedCodes[i+1:]...)
				break
			}
		}

		var outNum string
		for _, n := range outputNums {
			outNum += strconv.Itoa(digitByCode[n])
		}
		val, _ := strconv.Atoi(outNum)
		out += val
	}
	return
}

func canonicalCodes(in []string) (out []string) {
	for _, c := range in {
		s := strings.Split(c, "")
		sort.Strings(s)
		out = append(out, strings.Join(s, ""))
	}
	return
}

func sublen(a, b string) int {
	for _, c := range b {
		a = strings.Replace(a, string(c), "", -1)
	}

	return len(a)
}
