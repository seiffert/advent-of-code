package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/seiffert/advent-of-code/2019/lib"
	"github.com/seiffert/advent-of-code/2019/lib/intcode"
)

func main() {
	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		fileInput, err := ioutil.ReadFile("input.txt")
		if err != nil {
			lib.Abort("error reading input file: %w", err)
		}
		input = string(fileInput)
	}

	p := intcode.Parse(strings.TrimSpace(input))

	var maxResult int
	for _, phaseSetting := range permutations([]int{0, 1, 2, 3, 4}) {
		result := calculate(p, phaseSetting)
		if result > maxResult {
			maxResult = result
		}
	}
	fmt.Println("max result (step 1): ", maxResult)

	maxResult = 0
	for _, phaseSetting := range permutations([]int{5, 6, 7, 8, 9}) {
		result := calculate(p, phaseSetting)
		if result > maxResult {
			maxResult = result
		}
	}
	fmt.Println("max result (step 2): ", maxResult)
}

func permutations(in []int) [][]int {
	var generator func([]int, int)
	var res [][]int

	generator = func(in []int, n int) {
		if n == 1 {
			tmp := make([]int, len(in))
			copy(tmp, in)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			generator(in, n-1)
			if n%2 == 1 {
				tmp := in[i]
				in[i] = in[n-1]
				in[n-1] = tmp
			} else {
				tmp := in[0]
				in[0] = in[n-1]
				in[n-1] = tmp
			}
		}
	}
	generator(in, len(in))
	return res
}

func calculate(p, phases []int) int {
	var wg sync.WaitGroup

	// set up all amps
	in := make(chan int, 1)
	firstIn := in
	for i := 0; i < 5; i++ {
		var amp *Amplifier
		// connect the amp to the output of the previous one
		amp, in = NewAmplifier(p, in)
		amp.In <- phases[i]

		wg.Add(1)
		go func(amp *Amplifier, i int) {
			defer wg.Done()
			if err := amp.Run(); err != nil {
				lib.Abort("error calculating result: %w", err)
			}
		}(amp, i)
	}

	// send initial input
	firstIn <- 0

	// connect last output to first input
	go func() {
		for i := range in {
			firstIn <- i
		}
	}()

	// wait until all amps halted
	wg.Wait()

	return <-firstIn
}
