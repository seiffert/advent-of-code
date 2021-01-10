package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

var pwdLineRegexp = regexp.MustCompile(`^([0-9]+)-([0-9]+) ([a-z]): (\w+)$`)

func main() {
	input := lib.MustReadFile("input.txt")

	pwds, err := ParsePasswordList(strings.Split(input, "\n"))
	if err != nil {
		lib.Abort("failed to parse passwords: %v", err)
	}

	valid, valid2 := CountValidPasswords(pwds)

	fmt.Printf("the list contains %d / %d valid passwords\n", valid, valid2)
}

func CountValidPasswords(pwds []*Password) (int, int) {
	var valid, valid2 int
	for _, pwd := range pwds {
		if pwd.IsValid() {
			valid++
		}
		if pwd.IsValid2() {
			valid2++
		}
	}
	return valid, valid2
}

func ParsePasswordList(input []string) ([]*Password, error) {
	var result []*Password
	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			continue
		}

		pwd, err := NewPassword(line)
		if err != nil {
			return nil, fmt.Errorf("could not parse password line %q: %v", line, err)
		}

		result = append(result, pwd)
	}
	return result, nil
}

type Password struct {
	policy   Policy
	password string
}
type Policy struct {
	Letter rune
	Min    int
	Max    int
}

func NewPassword(input string) (*Password, error) {
	parsed := pwdLineRegexp.FindStringSubmatch(input)
	if len(parsed) == 0 {
		return nil, fmt.Errorf("did not match expected password line")
	}

	min, err := strconv.Atoi(parsed[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse policy minimum %q: %v", parsed[1], err)
	}
	max, err := strconv.Atoi(parsed[2])
	if err != nil {
		return nil, fmt.Errorf("could not parse policy maximum %q: %v", parsed[2], err)
	}
	letter := parsed[3][0]

	return &Password{
		policy: Policy{
			Letter: rune(letter),
			Min:    min,
			Max:    max,
		},
		password: parsed[4],
	}, nil
}

func (p *Password) IsValid() bool {
	count := strings.Count(p.password, string(p.policy.Letter))
	return count >= p.policy.Min && count <= p.policy.Max
}

func (p *Password) IsValid2() bool {
	return (rune(p.password[p.policy.Min-1]) == p.policy.Letter ||
		rune(p.password[p.policy.Max-1]) == p.policy.Letter) &&
		!(rune(p.password[p.policy.Min-1]) == p.policy.Letter &&
			rune(p.password[p.policy.Max-1]) == p.policy.Letter)
}
