package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var triangles int
	for {
		coords := [][]int{}
		for i := 0; i < 3; i++ {
			if !s.Scan() {
				fmt.Printf("%d triangles", triangles)
				os.Exit(0)
			}
			line := s.Text()
			a, b, c := line[0:5], line[5:10], line[10:15]

			ai, _ := strconv.Atoi(strings.Trim(a, " "))
			bi, _ := strconv.Atoi(strings.Trim(b, " "))
			ci, _ := strconv.Atoi(strings.Trim(c, " "))

			coords = append(coords, []int{ai, bi, ci})
		}

		for i := 0; i < 3; i++ {
			if IsTriangle(coords[0][i],coords[1][i],coords[2][i]) {
				triangles++
			}
		}
	}
}
