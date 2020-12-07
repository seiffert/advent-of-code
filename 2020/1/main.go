package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	list, err := parseIntList(lib.MustReadFile("input.txt"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse list of numbers: %v\n", err)
		os.Exit(1)
	}

	ints, err := FindNIntsWithSum(2, 2020, list)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find 2 ints: %v", err)
		os.Exit(1)
	}
	fmt.Printf("the two ints are %v and their product %d\n", ints, mul(ints))

	ints, err = FindNIntsWithSum(3, 2020, list)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find 3 ints: %v", err)
		os.Exit(1)
	}
	fmt.Printf("the three ints are %v and their product %d\n", ints, mul(ints))
}

func FindNIntsWithSum(n, expectedSum int, list []int) (result []int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		rand.Seed(time.Now().Unix())
		for ctx.Err() == nil {
			result = pickN(n, list)
			if sum(result) == expectedSum {
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()

	// 🤞
	return result, ctx.Err()
}

func parseIntList(list string) ([]int, error) {
	var (
		numbers = strings.Split(list, "\n")
		ints    = make([]int, 0, len(numbers))
	)

	for _, num := range numbers {
		if len(num) == 0 {
			continue
		}

		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number %q: %v", num, err)
		}

		ints = append(ints, i)
	}

	return ints, nil
}

func mul(list []int) int {
	res := 1
	for _, i := range list {
		res *= i
	}
	return res
}

func sum(list []int) int {
	var res int
	for _, i := range list {
		res += i
	}
	return res
}

func pickN(n int, list []int) []int {
	ll := len(list)
	ints := make([]int, 0, n)
	for len(ints) < n {
		ints = append(ints, list[rand.Intn(ll)])
	}
	return ints
}
