package main

import (
	"fmt"
	"strings"
)

func ExampleErrorCorrect() {
	input := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`
	fmt.Println(ErrorCorrect(strings.Split(input, "\n")))
	// Output: advent
}
