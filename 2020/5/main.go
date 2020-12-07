package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

var (
	seatLocationBinarizer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
)

func main() {
	list := NewSeatList(string(lib.MustReadFile("input.txt")))
	fmt.Printf("the highest set ID is %d\n", list.HighestSeatID())
	fmt.Printf("my seat is %d\n", list.MySeatID())
}

type SeatList []Seat

func NewSeatList(input string) SeatList {
	lines := strings.Split(input, "\n")
	sl := make(SeatList, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			sl = append(sl, NewSeat(line))
		}
	}
	return sl
}

func (sl SeatList) HighestSeatID() int {
	var max int
	for _, s := range sl {
		max = int(math.Max(float64(max), float64(s.ID())))
	}
	return max
}

func (sl SeatList) MySeatID() int {
	sort.Slice(sl, func(i, j int) bool { return sl[i].ID() < sl[j].ID() })

	var prev int
	for i := 0; i < len(sl); i++ {
		if prev != 0 && sl[i].ID() != prev+1 {
			return prev + 1
		}
		prev = sl[i].ID()
	}
	return -1
}

type Seat struct{ row, column int }

func NewSeat(input string) Seat {
	binarySeatLog := seatLocationBinarizer.Replace(input)

	row, _ := strconv.ParseInt(binarySeatLog[:7], 2, 64)
	column, _ := strconv.ParseInt(binarySeatLog[7:], 2, 64)
	return Seat{int(row), int(column)}
}

func (s Seat) ID() int {
	return s.column + s.row*8
}
