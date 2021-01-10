package main

import (
	"fmt"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	fmt.Println("Encryption Key:", CalcEncryptionKey(5099500, 7648211))
}

func CalcEncryptionKey(a, b int) int {
	lsA, lsB := CalcLoopSize(7, a), CalcLoopSize(7, b)

	encA, encB := transform(lsB, a), transform(lsA, b)
	if encA != encB {
		lib.Abort("calculated encryption keys do not match: %d != %d", encA, encB)
	}
	return encA
}

func CalcLoopSize(sn, pk int) int {
	for val, ls := 1, 1; ; ls++ {
		val = transformOnce(val, sn)
		if pk == val {
			return ls
		}
	}
}

func transform(ls, sn int) int {
	val := 1
	for i := 0; i < ls; i++ {
		val = transformOnce(val, sn)
	}
	return val
}

func transformOnce(val, sn int) int {
	return val * sn % 20201227
}
