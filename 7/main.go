package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	GateRegexp   = regexp.MustCompile("^([a-zA-Z0-9 ]+) -> ([a-z]+)$")
	LshiftRegexp = regexp.MustCompile("^([a-z0-9]+) LSHIFT ([0-9a-z]+)$")
	RshiftRegexp = regexp.MustCompile("^([a-z0-9]+) RSHIFT ([0-9a-z]+)$")
	AndRegexp    = regexp.MustCompile("^([a-z0-9]+) AND ([a-z0-9]+)$")
	OrRegexp     = regexp.MustCompile("^([a-z0-9]+) OR ([a-z0-9]+)$")
	NotRegexp    = regexp.MustCompile("^NOT ([a-z0-9]+)$")
	ScalarRegexp = regexp.MustCompile("^([0-9]+)$")
	VarRegexp    = regexp.MustCompile("^([a-z]+)$")

	cache = map[string]uint16{}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gates := map[string]Gate{}
	for scanner.Scan() {
		line := scanner.Text()

		wire, gate := parseGate(line)
		gates[wire] = gate
	}

	wire := os.Args[1]
	wire2 := os.Args[2]

	result := gates[wire](gates)
	fmt.Printf("wire %s: %d\n", wire, result)
	fmt.Printf("now setting wire %s to %d\n", wire2, result)
	gates[wire2] = ScalarGate(result)
	cache = map[string]uint16{}

	fmt.Printf("wire %s: %d\n", wire, gates[wire](gates))
}

type Gate func(map[string]Gate) uint16

func CachingGate(s string, g Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		if v, ok := cache[s]; ok {
			return v
		}

		cache[s] = g(gates)
		return cache[s]
	}
}

func LshiftGate(in Gate, by Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		return in(gates) << by(gates)
	}
}

func RshiftGate(in Gate, by Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		return in(gates) >> by(gates)
	}
}

func ScalarGate(v uint16) Gate {
	return func(gates map[string]Gate) uint16 {
		return v
	}
}

func VarGate(in string) Gate {
	return func(gates map[string]Gate) uint16 {
		return gates[in](gates)
	}
}

func AndGate(a, b Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		return a(gates) & b(gates)
	}
}

func OrGate(a, b Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		return a(gates) | b(gates)
	}
}

func NotGate(in Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		return 65535 - in(gates)
	}
}

func parseGate(s string) (string, Gate) {
	matches := GateRegexp.FindAllStringSubmatch(s, -1)
	if len(matches) < 1 {
		log.Fatalf("No match found in line %s", s)
	}

	in := matches[0][1]
	out := matches[0][2]

	var gate Gate
	switch {
	case LshiftRegexp.MatchString(in):
		operands := LshiftRegexp.FindAllStringSubmatch(in, -1)
		gate = LshiftGate(parseSimple(operands[0][1]), parseSimple(operands[0][2]))
	case RshiftRegexp.MatchString(in):
		operands := RshiftRegexp.FindAllStringSubmatch(in, -1)
		gate = RshiftGate(parseSimple(operands[0][1]), parseSimple(operands[0][2]))
	case AndRegexp.MatchString(in):
		operands := AndRegexp.FindAllStringSubmatch(in, -1)
		gate = AndGate(parseSimple(operands[0][1]), parseSimple(operands[0][2]))
	case OrRegexp.MatchString(in):
		operands := OrRegexp.FindAllStringSubmatch(in, -1)
		gate = OrGate(parseSimple(operands[0][1]), parseSimple(operands[0][2]))
	case NotRegexp.MatchString(in):
		operand := NotRegexp.FindAllStringSubmatch(in, -1)
		gate = NotGate(parseSimple(operand[0][1]))
	case ScalarRegexp.MatchString(in):
		fallthrough
	case VarRegexp.MatchString(in):
		gate = parseSimple(in)
	default:
		log.Fatalf("Could not parse %s", s)
	}

	return out, CachingGate(s, gate)
}

func parseSimple(s string) Gate {
	if _, err := strconv.ParseUint(s, 10, 16); err != nil {
		return VarGate(s)
	}

	return ScalarGate(i(s))
}

func i(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)

	return uint16(i)
}
