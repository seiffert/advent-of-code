package main

import (
	"strconv"
)

func IsValidPassword1(pwd int) bool {
	str := strconv.Itoa(pwd)

	var (
		double    = false
		monotonic = true
		last      = ' '
	)
	for _, c := range str {
		if last != ' ' && c == last {
			double = true
		}
		if last != ' ' && c < last {
			monotonic = false
		}
		last = c
	}
	return double && monotonic
}

func IsValidPassword2(pwd int) bool {
	str := strconv.Itoa(pwd)

	var (
		double    = false
		streak    = 0
		monotonic = true
		last      = ' '
	)
	for _, c := range str {
		if last != ' ' {
			if c == last {
				streak++
			} else {
				if streak == 2 {
					double = true
				}
				streak = 1
			}
			if c < last {
				monotonic = false
			}
		} else {
			streak = 1
		}
		last = c
	}
	if streak == 2 {
		double = true
	}
	return double && monotonic
}
