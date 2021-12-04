package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	in := lib.MustReadFile("input.txt")

	fmt.Printf("power consumption: %d\n", PowerConsumption(in))
	fmt.Printf("life support rating: %d\n", LifeSupportRating(in))
}

func PowerConsumption(in string) int {
	lines := strings.Split(in, "\n")
	if len(lines) == 0 {
		lib.Abort("empty input")
	}

	var (
		b              = bits(lines)
		gamma, epsilon []byte
	)
	for _, b := range b {
		if b >= len(lines)/2 {
			gamma, epsilon = append(gamma, '1'), append(epsilon, '0')
		} else {
			gamma, epsilon = append(gamma, '0'), append(epsilon, '1')
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 32)
	e, _ := strconv.ParseInt(string(epsilon), 2, 32)

	return int(g * e)
}

func LifeSupportRating(in string) int {
	lines := strings.Split(in, "\n")
	if len(lines) == 0 {
		lib.Abort("empty input")
	}

	oxNums := make([]string, len(lines))
	copy(oxNums, lines)
	co2Nums := make([]string, len(lines))
	copy(co2Nums, lines)

	for i := 0; len(oxNums) > 1; i++ {
		var (
			b   = bits(oxNums)
			tmp []string
		)
		for _, n := range oxNums {
			mostCommonlySet := float64(b[i]) >= float64(len(oxNums))/2
			if mostCommonlySet && n[i] == '1' || !mostCommonlySet && n[i] == '0' {
				tmp = append(tmp, n)
			}
		}
		oxNums = tmp
	}

	for i := 0; len(co2Nums) > 1; i++ {
		var (
			b   = bits(co2Nums)
			tmp []string
		)
		for _, n := range co2Nums {
			mostCommonlySet := float64(b[i]) >= float64(len(co2Nums))/2
			if mostCommonlySet && n[i] == '0' || !mostCommonlySet && n[i] == '1' {
				tmp = append(tmp, n)
			}
		}
		co2Nums = tmp
	}

	ox, _ := strconv.ParseInt(string(oxNums[0]), 2, 32)
	co2, _ := strconv.ParseInt(string(co2Nums[0]), 2, 32)

	return int(ox * co2)
}

func bits(in []string) (bits []int) {
	bits = make([]int, len(in[0]))
	for _, l := range in {
		for i := 0; i < len(l); i++ {
			if l[i] == '1' {
				bits[i]++
			}
		}
	}
	return
}
