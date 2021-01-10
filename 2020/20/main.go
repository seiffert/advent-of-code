package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	p := NewSatelliteImage(lib.MustReadFile("input.txt"))

	fmt.Printf("corners multiplied: %d\n", p.MultiplyCorners())
	fmt.Printf("water roughness: %d\n", p.WaterRoughness())
}

type (
	SatelliteImage struct {
		image  map[int]map[int]*tile
		pixels [][]rune
	}
	tile struct {
		id     int
		pixels [][]rune
	}
)

func NewSatelliteImage(input string) *SatelliteImage {
	var tiles []tile
	for _, rawTile := range strings.Split(input, "\n\n") {
		lines := strings.Split(rawTile, "\n")
		id, _ := strconv.Atoi(strings.Trim(lines[0], "Tile :"))
		t := tile{id: id}
		for _, line := range lines[1:] {
			var r []rune
			for _, p := range line {
				r = append(r, p)
			}
			t.pixels = append(t.pixels, r)
		}
		tiles = append(tiles, t)
	}

	return reassembleSatelliteImage(tiles)
}

func reassembleSatelliteImage(tiles []tile) *SatelliteImage {
	// assumption: the image is a square
	l := math.Sqrt(float64(len(tiles)))

	images := []*SatelliteImage{{
		image: make(map[int]map[int]*tile),
	}}
	for y := 0; float64(y) < l; y++ {
		for x := 0; float64(x) < l; x++ {
			var newImages []*SatelliteImage
			for _, p := range images {
				for _, t := range tiles {
					if p.contains(t) {
						continue
					}
					for _, v := range t.variants() {
						img := p.clone()
						if _, ok := img.image[y]; !ok {
							img.image[y] = make(map[int]*tile)
						}
						img.image[y][x] = v
						if img.IsValid() {
							newImages = append(newImages, img)
						}
					}
				}
			}
			images = newImages
			fmt.Printf("calculating options for (%d/%d), possible combinations so far: %d\n", x, y, len(images))
		}
	}

	// there can be multiple valid combinations (e.g. all tiles rotated / flipped)
	img := images[0]

	for yt := 0; yt < len(img.image); yt++ {
		for yp := 1; yp < len(img.image[yt][0].pixels)-1; yp++ {
			var row []rune
			for xt := 0; xt < len(img.image[yt]); xt++ {
				for xp := 1; xp < len(img.image[yt][xt].pixels[yp])-1; xp++ {
					row = append(row, img.image[yt][xt].pixels[yp][xp])
				}
			}
			img.pixels = append(img.pixels, row)
		}
	}

	return img
}

func (si *SatelliteImage) MultiplyCorners() int64 {
	return int64(si.image[0][0].id) *
		int64(si.image[0][len(si.image[0])-1].id) *
		int64(si.image[len(si.image)-1][0].id) *
		int64(si.image[len(si.image)-1][len(si.image[len(si.image)-1])-1].id)
}

func (si *SatelliteImage) IsValid() bool {
	for y := 0; y < len(si.image); y++ {
		for x := 0; x < len(si.image[y]); x++ {
			if x != len(si.image[y])-1 {
				a, b := si.image[y][x], si.image[y][x+1]
				for i := 0; i < len(a.pixels); i++ {
					if a.pixels[i][len(a.pixels[i])-1] != b.pixels[i][0] {
						return false
					}
				}
			}
			if y != len(si.image)-1 {
				a, b := si.image[y][x], si.image[y+1][x]
				if a == nil || b == nil {
					continue
				}
				for i := 0; i < len(a.pixels[len(a.pixels)-1]); i++ {
					if a.pixels[len(a.pixels)-1][i] != b.pixels[0][i] {
						return false
					}
				}
			}
		}
	}
	return true
}

const seaMonster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

func (si *SatelliteImage) WaterRoughness() int {
	sm := tile{}
	for _, line := range strings.Split(seaMonster, "\n") {
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		sm.pixels = append(sm.pixels, row)
	}

	var numSeamonsters int
	for _, sm := range sm.variants() {
		numSeamonsters += si.countSeaMonsters(sm)
	}
	fmt.Printf("found %d sea monsters\n", numSeamonsters)

	return si.countPixels() - numSeamonsters*sm.countPixels()
}

func (si *SatelliteImage) String() string {
	var res string
	for _, r := range si.pixels {
		for _, p := range r {
			res += string(p)
		}
		res += "\n"
	}
	return res
}

func (si *SatelliteImage) countPixels() int {
	var numPixels int
	for _, r := range si.pixels {
		for _, p := range r {
			if p == '#' {
				numPixels++
			}
		}
	}
	return numPixels
}

func (si *SatelliteImage) countSeaMonsters(t *tile) int {
	var numSeamonsters int

	for yi := 0; yi < len(si.pixels)-len(t.pixels); yi++ {
	nextPos:
		for xi := 0; xi < len(si.pixels[yi])-len(t.pixels[0]); xi++ {
			for yt := 0; yt < len(t.pixels); yt++ {
				for xt := 0; xt < len(t.pixels[yt]); xt++ {
					if t.pixels[yt][xt] == '#' && si.pixels[yi+yt][xi+xt] != '#' {
						continue nextPos
					}
				}
			}
			numSeamonsters++
		}
	}
	return numSeamonsters
}

func (si *SatelliteImage) clone() *SatelliteImage {
	o := &SatelliteImage{image: make(map[int]map[int]*tile)}
	for y, _ := range si.image {
		for x, t := range si.image[y] {
			if _, ok := o.image[y]; !ok {
				o.image[y] = make(map[int]*tile)
			}
			o.image[y][x] = t.clone()
		}
	}
	return o
}

func (si *SatelliteImage) contains(o tile) bool {
	for _, r := range si.image {
		for _, t := range r {
			if t.id == o.id {
				return true
			}
		}
	}
	return false
}

func (t *tile) clone() *tile {
	o := &tile{}
	o.id = t.id
	for y, _ := range t.pixels {
		r := []rune{}
		for _, c := range t.pixels[y] {
			r = append(r, c)
		}
		o.pixels = append(o.pixels, r)
	}
	return o
}

func (t *tile) variants() []*tile {
	r90 := t.rotate()
	r180 := r90.rotate()
	r270 := r180.rotate()

	return []*tile{
		t, t.flipX(), t.flipY(),
		r90, r90.flipX(), r90.flipY(),
		r180,
		r270,
	}
}

func (t *tile) flipX() *tile {
	n := &tile{id: t.id}
	for y := len(t.pixels) - 1; y >= 0; y-- {
		row := []rune{}
		for x := 0; x < len(t.pixels[y]); x++ {
			row = append(row, t.pixels[y][x])
		}
		n.pixels = append(n.pixels, row)
	}
	return n
}

func (t *tile) flipY() *tile {
	n := &tile{id: t.id}
	for y := 0; y < len(t.pixels); y++ {
		row := []rune{}
		for x := len(t.pixels[y]) - 1; x >= 0; x-- {
			row = append(row, t.pixels[y][x])
		}
		n.pixels = append(n.pixels, row)
	}
	return n
}

func (t *tile) rotate() *tile {
	n := &tile{id: t.id}

	for x := 0; x < len(t.pixels[0]); x++ {
		var row []rune
		for y := len(t.pixels) - 1; y >= 0; y-- {
			row = append(row, t.pixels[y][x])
		}
		n.pixels = append(n.pixels, row)
	}
	return n
}

func (t *tile) String() string {
	var res string
	for _, r := range t.pixels {
		for _, p := range r {
			res += string(p)
		}
		res += "\n"
	}
	return res
}

func (t *tile) countPixels() int {
	var numPixels int
	for _, r := range t.pixels {
		for _, p := range r {
			if p == '#' {
				numPixels++
			}
		}
	}
	return numPixels
}
