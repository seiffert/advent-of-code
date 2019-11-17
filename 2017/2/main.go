package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input file: %v", err)
		os.Exit(1)
	}

	sum, err := Checksum(string(in), minMaxChecksum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating checksum: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Checksum based on min/max: %d\n", sum)

	sum, err = Checksum(string(in), evenlyDevisibleChecksum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating checksum: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Checksum based on evenly divisible numbers: %d\n", sum)
}

func Checksum(in string, rowChecksum func([]int) int) (int, error) {
	sheet, err := parse(in)
	if err != nil {
		return 0, fmt.Errorf("could not parse input: %v", err)
	}

	var sum int
	for _, row := range sheet {
		sum += rowChecksum(row)
	}
	return sum, nil
}

func parse(in string) ([][]int, error) {
	var result [][]int
	rows := strings.Split(in, "\n")

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		var resultRow []int
		cells := regexp.MustCompile("[ \t]").Split(row, -1)
		for _, cell := range cells {
			num, err := strconv.Atoi(cell)
			if err != nil {
				return nil, fmt.Errorf("not a number: %q", cell)
			}
			resultRow = append(resultRow, num)
		}
		result = append(result, resultRow)
	}
	return result, nil
}

func minMaxChecksum(row []int) int {
	var (
		max = 0
		min = math.MaxInt32
	)
	for _, num := range row {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return max - min
}

func evenlyDevisibleChecksum(row []int) int {
	for x, numX := range row {
		for y, numY := range row {
			if x == y {
				continue
			}
			if float32(numX)/float32(numY) == float32(numX/numY) {
				return numX / numY
			}
			if float32(numY)/float32(numX) == float32(numY/numX) {
				return numY / numX
			}
		}
	}
	return 0
}
