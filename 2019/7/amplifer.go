package main

import "github.com/seiffert/advent-of-code/2019/lib/intcode"

func NewAmplifier(p []int, in chan int) (*Amplifier, chan int) {
	computer := intcode.NewComputer(p)
	receiver := intcode.ChanReceiver(in)
	computer.SetInputReceiver(receiver)

	out := make(chan int, 1)
	computer.SetOutputSender(intcode.ChanSender(out))

	return &Amplifier{
		computer: computer,

		In:  in,
		Out: out,
	}, out
}

type Amplifier struct {
	In  chan<- int
	Out <-chan int

	computer *intcode.Computer
}

func (a *Amplifier) Run() error {
	return a.computer.Calculate()
}
