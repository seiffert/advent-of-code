package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	var (
		input      = lib.MustReadFile("input.txt")
		lines      = strings.Split(input, "\n")
		arrival, _ = strconv.Atoi(lines[0])

		schedule = NewSchedule(lines[1])
		busID, t = schedule.A(arrival)
	)

	fmt.Printf("bus ID %d, arrives at %d, result: %d\n", busID, t, busID*(t-arrival))
	fmt.Printf("b: %d\n", schedule.B())
}

type Schedule []int

func NewSchedule(input string) Schedule {
	var s Schedule
	for _, busID := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(busID)
		s = append(s, i)
	}
	return s
}

func (s Schedule) A(t int) (int, int) {
	closest := 0
	for _, b := range s {
		if b == 0 {
			continue
		}
		if closest == 0 || t%b > t%closest {
			closest = b
		}
	}
	return closest, (t/closest + 1) * closest
}

func (s Schedule) B() int64 {
	start, step, stepI := int64(s[0]), s[0], 0

outer:
	for n := start; ; n += int64(step) {
		for i := 1; i < len(s); i++ {
			if s[i] == 0 {
				continue
			}
			if (n+int64(i))%int64(s[i]) != 0 {
				continue outer
			}
			if stepI < i {
				step *= s[i]
				stepI = i
			}
		}
		return n
	}
}
