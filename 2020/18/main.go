package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	var sumWithoutPrecedence, sumWithPrecedence int
	for _, line := range strings.Split(lib.MustReadFile("input.txt"), "\n") {
		sumWithoutPrecedence += CalculateWithoutPrecedence(line)
		sumWithPrecedence += CalculateWithPrecedence(line)
	}

	fmt.Printf("result without precedence: %d\n", sumWithoutPrecedence)
	fmt.Printf("result with precedence: %d\n", sumWithPrecedence)
}

func CalculateWithoutPrecedence(input string) int {
	result, _ := calculate(sortWithoutPrecedence(tokenize(input)))
	return result
}

func CalculateWithPrecedence(input string) int {
	result, _ := calculate(sortWithPrecedence(tokenize(input)))
	return result
}

type token string

func tokenize(input string) []token {
	var ts []token
	for i, c := range input {
		switch c {
		case ' ':
			continue
		case '(', ')', '+', '-', '*', '/':
			ts = append(ts, token(c))
		default:
			num := string(c)
			for la := 1; la+i < len(input); la++ {
				if input[i+la] >= '0' && input[i+la] <= 9 {
					num += string(input[i+la])
				}
			}
			ts = append(ts, token(num))
		}
	}
	return ts
}

func calculate(ts []token) (int, []token) {
	operator, ts := ts[len(ts)-1], ts[:len(ts)-1]
	var numA, numB int

	if !isOperator(ts[len(ts)-1]) {
		numA, _ = strconv.Atoi(string(ts[len(ts)-1]))
		ts = ts[:len(ts)-1]
	} else {
		numA, ts = calculate(ts)
	}

	if !isOperator(ts[len(ts)-1]) {
		numB, _ = strconv.Atoi(string(ts[len(ts)-1]))
		ts = ts[:len(ts)-1]
	} else {
		numB, ts = calculate(ts)
	}

	switch operator {
	case "+":
		return numA + numB, ts
	case "-":
		return numA - numB, ts
	case "*":
		return numA * numB, ts
	case "/":
		return numA / numB, ts
	}

	return -1, ts
}

func isOperator(in token) bool {
	return in == "+" || in == "-" || in == "*" || in == "/"
}

func sortWithoutPrecedence(ts []token) []token {
	var ops []token
	var q []token

	for i := len(ts) - 1; i >= 0; i-- {
		switch ts[i] {
		case "+", "*", "-", "/":
			ops = append(ops, ts[i])
		case ")":
			ops = append(ops, ts[i])
		case "(":
			for j := len(ops) - 1; j >= 0; j-- {
				if ops[j] == ")" {
					ops = ops[:j]
					break
				}
				q = append(q, ops[j])
			}
		default:
			q = append(q, ts[i])
		}
	}
	for i := len(ops) - 1; i >= 0; i-- {
		q = append(q, ops[i])
	}

	return q
}

func sortWithPrecedence(ts []token) []token {
	var ops []token
	var q []token

	for i := 0; i < len(ts); i++ {
		switch ts[i] {
		case "+", "-":
			ops = append(ops, ts[i])
		case "*", "/":
			for j := len(ops) - 1; j >= 0; j-- {
				op := ops[j]

				if op == "+" || op == "-" {
					q = append(q, op)
					ops = ops[:j]
				} else {
					break
				}
			}
			ops = append(ops, ts[i])
		case "(":
			ops = append(ops, ts[i])
		case ")":
			for j := len(ops) - 1; j >= 0; j-- {
				op := ops[j]
				ops = ops[:j]
				if op == "(" {
					break
				}
				q = append(q, op)
			}
		default:
			q = append(q, ts[i])
		}
	}
	for i := len(ops) - 1; i >= 0; i-- {
		q = append(q, ops[i])
	}

	return q
}
