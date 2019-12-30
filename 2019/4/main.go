package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func main() {
	if len(os.Args) != 3 {
		lib.Abort("expects the min and max value as arguments")
	}

	min := os.Args[1]
	max := os.Args[2]

	minPwd, err := strconv.Atoi(min)
	if err != nil {
		lib.Abort("error parsing min password: %w", err)
	}
	maxPwd, err := strconv.Atoi(max)
	if err != nil {
		lib.Abort("error parsing max password: %w", err)
	}

	var validPasswords1, validPasswords2 int
	for pwd := minPwd; pwd < maxPwd; pwd++ {
		if IsValidPassword1(pwd) {
			validPasswords1++
		}
		if IsValidPassword2(pwd) {
			validPasswords2++
		}
	}

	fmt.Println("number of valid passwords (step 1):", validPasswords1)
	fmt.Println("number of valid passwords (step 2):", validPasswords2)
}
