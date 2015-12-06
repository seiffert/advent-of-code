package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lights := [1000][1000]bool{}
	r := regexp.MustCompile("^(.*) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$")

	var count int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()

		m := r.FindAllStringSubmatch(instruction, -1)
		ax, ay, bx, by := toInt(m[0][2]), toInt(m[0][3]), toInt(m[0][4]), toInt(m[0][5])

		for x := ax; x <= bx; x++ {
			for y := ay; y <= by; y++ {
				switch m[0][1] {
				case "turn on":
					if !lights[x][y] {
						count++
					}
					lights[x][y] = true
				case "turn off":
					if lights[x][y] {
						count--
					}
					lights[x][y] = false
				case "toggle":
					if lights[x][y] {
						count--
					} else {
						count++
					}
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	fmt.Printf("%d\n", count)
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
