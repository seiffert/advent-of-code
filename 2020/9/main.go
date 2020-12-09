package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	s := NewEncryptedStream(lib.MustReadFile("input.txt"))

	invalidNumber := s.FindFirstInvalidNumber(25)
	fmt.Printf("first invalid number: %d\n", invalidNumber)

	min, max := s.FindContiguousSetWithSum(invalidNumber)
	fmt.Printf("contiguous set found: %d + %d = %d\n", min, max, min+max)
}

type EncryptedStream []int64

func NewEncryptedStream(input string) EncryptedStream {
	lines := strings.Split(input, "\n")
	es := make(EncryptedStream, 0, len(lines))
	for _, line := range lines {
		num, _ := strconv.ParseInt(line, 10, 64)
		es = append(es, num)
	}
	return es
}

func (es EncryptedStream) FindFirstInvalidNumber(l int) int64 {
	for i := l; i < len(es); i++ {
		n := es[i]
		var valid bool
		for _, o := range es[i-l : i] {
			for _, p := range es[i-l : i] {
				if o+p == n && o != p {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid {
			return n
		}
	}
	return -1
}

func (es EncryptedStream) FindContiguousSetWithSum(s int64) (int64, int64) {
	for start := 0; start < len(es)-1; start++ {
		for end := start + 1; end < len(es); end++ {
			var sum, min, max int64
			for _, n := range es[start : end+1] {
				sum += n
				if min == 0 || n < min {
					min = n
				}
				if max == 0 || n > max {
					max = n
				}
			}

			if sum == s {
				return min, max
			}
		}
	}
	return 0, 0
}
