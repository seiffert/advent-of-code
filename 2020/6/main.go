package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	list := NewAnswerList(lib.MustReadFile("input.txt"))

	result := list.SumGroups(func(g GroupAnswers) int {
		return g.PositivelyAnsweredQuestions()
	})
	fmt.Printf("The sum of positively answered questions across groups is %d\n", result)

	result = list.SumGroups(func(g GroupAnswers) int {
		return g.UnanimousPositivelyAnsweredQuestions()
	})
	fmt.Printf("The sum of questions which were answered positively by all is %d\n", result)
}

type AnswerList []GroupAnswers

func NewAnswerList(input string) AnswerList {
	al := AnswerList{}
	for _, group := range strings.Split(input, "\n\n") {
		al = append(al, GroupAnswers(strings.Split(group, "\n")))
	}
	return al
}

func (al AnswerList) SumGroups(counter func(GroupAnswers) int) int {
	var result int
	for _, ga := range al {
		result += counter(ga)
	}
	return result
}

type GroupAnswers []string

func (ga GroupAnswers) PositivelyAnsweredQuestions() int {
	qs := map[byte]struct{}{}
	for _, a := range ga {
		for i := 0; i < len(a); i++ {
			qs[a[i]] = struct{}{}
		}
	}
	return len(qs)
}

func (ga GroupAnswers) UnanimousPositivelyAnsweredQuestions() int {
	qs := []byte(ga[0])
	for _, a := range ga {
		for i := len(qs) - 1; i >= 0; i-- {
			if !strings.Contains(a, string(qs[i])) {
				qs = append(qs[:i], qs[i+1:]...)
			}
		}
	}
	return len(qs)
}
