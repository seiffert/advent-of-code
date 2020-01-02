package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

type (
	Dimension struct {
		Width, Height int
	}
	Image struct {
		Layers []Layer
	}
	Layer [][]int
)

func main() {
	if len(os.Args) < 2 {
		lib.Abort("expects dimension as first argument")
	}
	dimensionArg := os.Args[1]

	var inputArg string
	if len(os.Args) == 3 {
		inputArg = os.Args[2]
	} else {
		fileInput, err := ioutil.ReadFile("input.txt")
		if err != nil {
			lib.Abort("error reading input file: %w", err)
		}
		inputArg = strings.TrimSpace(string(fileInput))
	}

	dimension := parseDimension(dimensionArg)
	img := parseImage(inputArg, dimension)

	var fewestZerosLayer Layer
	var fewestZeros int
	for _, layer := range img.Layers {
		zeros := layer.Count(0)
		if fewestZerosLayer == nil || zeros < fewestZeros {
			fewestZeros = zeros
			fewestZerosLayer = layer
		}
	}

	fmt.Println("Result step 1:", fewestZerosLayer.Count(1)*fewestZerosLayer.Count(2))
	fmt.Println("")

	fmt.Println("Flattened image:")
	fmt.Println("")
	fmt.Println(img.Flatten())
}

func parseDimension(in string) Dimension {
	dim := strings.Split(in, "x")
	if len(dim) != 2 {
		lib.Abort("invalid dimension %q", in)
	}

	w, err := strconv.Atoi(dim[0])
	if err != nil {
		lib.Abort("invalid width %q", dim[0])
	}

	h, err := strconv.Atoi(dim[1])
	if err != nil {
		lib.Abort("invalid height %q", dim[1])
	}

	return Dimension{w, h}
}

func parseImage(in string, dim Dimension) *Image {
	layerLength := dim.Width * dim.Height
	inLength := len(in)

	if inLength%layerLength != 0 {
		lib.Abort("invalid image input length: %d, expected multiple of %d", inLength, layerLength)
	}

	img := &Image{}
	for i := 0; i < inLength/layerLength; i++ {
		layer := Layer{}
		for y := 0; y < dim.Height; y++ {
			row := make([]int, dim.Width)
			for x := 0; x < dim.Width; x++ {
				inPixel := string(in[i*layerLength+y*dim.Width+x])
				pixel, err := strconv.Atoi(inPixel)
				if err != nil {
					lib.Abort("invalid pixel value %q", inPixel)
				}

				row[x] = pixel
			}
			layer = append(layer, row)
		}
		img.Layers = append(img.Layers, layer)
	}
	return img
}

func (i Image) Flatten() Layer {
	result := Layer{}
	for y := 0; y < len(i.Layers[0]); y++ {
		row := make([]int, len(i.Layers[0][0]))
		for x := 0; x < len(i.Layers[0][0]); x++ {
			for layer := 0; layer < len(i.Layers); layer++ {
				if i.Layers[layer][y][x] != 2 {
					row[x] = i.Layers[layer][y][x]
					break
				}
			}
		}
		result = append(result, row)
	}
	return result
}

func (l Layer) String() string {
	var result strings.Builder
	for _, row := range l {
		for _, p := range row {
			result.WriteString(strconv.Itoa(p))
		}
		result.WriteRune('\n')
	}
	return result.String()
}

func (l Layer) Count(q int) int {
	var count int
	for _, r := range l {
		for _, c := range r {
			if c == q {
				count++
			}
		}
	}
	return count
}
