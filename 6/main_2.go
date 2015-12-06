package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/jpeg"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lights := [1000][1000]int{}
	r := regexp.MustCompile("^(.*) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$")

	i := image.NewPaletted(image.Rect(0, 0, 1000, 1000), palette.WebSafe)

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
					count++
					lights[x][y]++
				case "turn off":
					if lights[x][y] > 0 {
						count--
						lights[x][y]--
					}
				case "toggle":
					count += 2
					lights[x][y] += 2
				}
			}
		}
	}
	for x, _ := range lights {
		for y, _ := range lights[x] {
			i.Set(x, y, color.RGBA{uint8(lights[x][y] * 10), 0, uint8(lights[x][y] * 2), 255})
		}
	}

	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt jpeg.Options
	opt.Quality = 80
	err = jpeg.Encode(out, i, &opt) // put quality to 80%
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", count)
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
