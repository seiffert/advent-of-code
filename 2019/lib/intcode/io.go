package intcode

import (
	"fmt"

	"github.com/seiffert/advent-of-code/2019/lib"
)

type InputReceiver func() int
type OutputSender func(int)

func StdinReceiver() int {
	var (
		valid bool
		input int
	)

	for !valid {
		fmt.Println("Expecting input:")
		if _, err := fmt.Scanf("%d", &input); err != nil {
			lib.LogError("invalid input: %w", err)
		} else {
			valid = true
		}
	}
	return input
}

func ChanReceiver(inChan <-chan int) InputReceiver {
	return func() int {
		return <-inChan
	}
}

func StdoutSender(val int) {
	fmt.Printf("Output: %d\n", val)
}

func ChanSender(outChan chan<- int) OutputSender {
	return func(val int) {
		outChan <- val
	}
}
