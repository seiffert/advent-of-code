package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

var errInfiniteLoopDetected = fmt.Errorf("infinite loop detected")

func main() {
	b := NewBootSequence(lib.MustReadFile("input.txt"))

	err := b.Run()
	switch err {
	case errInfiniteLoopDetected:
		fmt.Printf("found infinite loop. last acc: %d\n", b.Acc)
	case nil:
		lib.Abort("boot sequence terminated without a loop")
	default:
		lib.Abort("boot sequence terminated with a different error: %v", err)
	}

	bs := AllPossibleBootSequences(lib.MustReadFile("input.txt"))
	for _, b := range bs {
		if err := b.Run(); err == nil {
			fmt.Printf("boot sequence terminated without an error. acc: %d\n", b.Acc)
		}
	}
}

type BootSequence struct {
	Cmds    []*Command
	Acc, Pc int
}

func AllPossibleBootSequences(input string) []*BootSequence {
	var result []*BootSequence
	rawCommands := strings.Split(input, "\n")
	r := strings.NewReplacer("nop", "jmp", "jmp", "nop")
	for i := 0; i < len(rawCommands); i++ {
		newCmd := r.Replace(rawCommands[i])
		if newCmd != rawCommands[i] {
			modifiedCommands := make([]string, len(rawCommands))
			copy(modifiedCommands, rawCommands)
			modifiedCommands[i] = newCmd

			result = append(result, NewBootSequence(strings.Join(modifiedCommands, "\n")))
		}
	}
	return result
}

func NewBootSequence(input string) *BootSequence {
	rawCommands := strings.Split(input, "\n")
	b := &BootSequence{}
	for _, rawCmd := range rawCommands {
		b.Cmds = append(b.Cmds, NewCommand(rawCmd))
	}
	return b
}

func (b *BootSequence) Run() error {
	for b.Pc < len(b.Cmds) {
		cmd := b.Cmds[b.Pc]
		if cmd.Rc > 0 {
			return errInfiniteLoopDetected
		}

		cmd.Rc++
		switch cmd.Op {
		case "acc":
			b.Acc += cmd.Arg
			b.Pc++
		case "jmp":
			b.Pc += cmd.Arg
		case "nop":
			b.Pc++
		}
	}
	return nil
}

type Command struct {
	Op  string
	Arg int
	Rc  int
}

func NewCommand(input string) *Command {
	fields := strings.Fields(input)
	arg, _ := strconv.Atoi(fields[1])
	return &Command{Op: fields[0], Arg: arg}
}
