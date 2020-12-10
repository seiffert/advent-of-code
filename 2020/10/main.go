package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	a := NewAdapterList(lib.MustReadFile("input.txt"))

	d1, _, d3 := a.Distribution()
	fmt.Printf("%d * %d = %d\n", d1, d3, d1*d3)

	fmt.Printf("possible combinations: %d\n", a.PossibleCombinations())
}

type AdapterList []int

func NewAdapterList(input string) AdapterList {
	rawAdapters := strings.Split(input, "\n")
	list := make(AdapterList, 0, len(rawAdapters))
	for _, rawAdapter := range rawAdapters {
		adapter, _ := strconv.Atoi(rawAdapter)
		list = append(list, adapter)
	}
	sort.Ints(list)
	return list
}

func (al AdapterList) Distribution() (int, int, int) {
	dist := map[int]int{}
	var prev int
	for _, a := range al {
		d := a - prev
		dist[d] = dist[d] + 1
		prev = a
	}
	return dist[1], dist[2], dist[3] + 1
}

func (al AdapterList) PossibleCombinations() int64 {
	cache := map[string]int64{}
	return possibleCombinations(AdapterList{}, al, al[len(al)-1], cache)
}

func possibleCombinations(prefix AdapterList, adapters AdapterList, end int, cache map[string]int64) int64 {
	cacheKey := fmt.Sprintf("%#v", adapters)
	if cachedResult, ok := cache[cacheKey]; ok {
		return cachedResult
	}

	if len(adapters) == 0 {
		return 1
	}

	var last int
	if len(prefix) > 0 {
		last = prefix[len(prefix)-1]
	}
	var num int64
	for i := 0; i < len(adapters); i++ {
		if adapters[i]-last > 3 {
			break
		}

		np := make(AdapterList, len(prefix)+1)
		copy(np, prefix)
		np[len(np)-1] = adapters[i]

		num += possibleCombinations(np, adapters[i+1:], end, cache)
	}

	cache[cacheKey] = num
	return num
}
