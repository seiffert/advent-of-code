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

	GateConstructors = map[*regexp.Regexp]GateConstructor{
		regexp.MustCompile("^([a-z0-9]+) LSHIFT ([0-9a-z]+)$"): LshiftGate,
		regexp.MustCompile("^([a-z0-9]+) RSHIFT ([0-9a-z]+)$"): RshiftGate,
		regexp.MustCompile("^([a-z0-9]+) AND ([a-z0-9]+)$"):    AndGate,
		regexp.MustCompile("^([a-z0-9]+) OR ([a-z0-9]+)$"):     OrGate,
		regexp.MustCompile("^NOT ([a-z0-9]+)$"):                NotGate,
		regexp.MustCompile("^([0-9]+)$"):                       SimpleGate,
		regexp.MustCompile("^([a-z]+)$"):                       SimpleGate,
	}

	cache = map[string]uint16{}
)

type Gate func(map[string]Gate) uint16

type GateConstructor func(...string) Gate

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

func parseGate(s string) (string, Gate) {
	matches := GateRegexp.FindAllStringSubmatch(s, -1)
	if len(matches) < 1 {
		log.Fatalf("No match found in line %s", s)
	}

	in := matches[0][1]
	out := matches[0][2]

	for regexp, constructor := range GateConstructors {
		if regexp.MatchString(in) {
			args := regexp.FindAllStringSubmatch(in, -1)
			return out, CachingGate(s, constructor(args[0][1:]...))
		}
	}

	return "", nil
}

func CachingGate(s string, g Gate) Gate {
	return func(gates map[string]Gate) uint16 {
		if v, ok := cache[s]; ok {
			return v
		}

		cache[s] = g(gates)
		return cache[s]
	}
}

func LshiftGate(args ...string) Gate {
	in := SimpleGate(args[0])
	by := SimpleGate(args[1])
	return func(gates map[string]Gate) uint16 {
		return in(gates) << by(gates)
	}
}

func RshiftGate(args ...string) Gate {
	in := SimpleGate(args[0])
	by := SimpleGate(args[1])
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

func AndGate(args ...string) Gate {
	a := SimpleGate(args[0])
	b := SimpleGate(args[1])
	return func(gates map[string]Gate) uint16 {
		return a(gates) & b(gates)
	}
}

func OrGate(args ...string) Gate {
	a := SimpleGate(args[0])
	b := SimpleGate(args[1])
	return func(gates map[string]Gate) uint16 {
		return a(gates) | b(gates)
	}
}

func NotGate(args ...string) Gate {
	in := SimpleGate(args[0])
	return func(gates map[string]Gate) uint16 {
		return 65535 - in(gates)
	}
}

func SimpleGate(args ...string) Gate {
	if i, err := strconv.ParseUint(args[0], 10, 16); err != nil {
		return VarGate(args[0])
	} else {
		return ScalarGate(uint16(i))
	}
}
