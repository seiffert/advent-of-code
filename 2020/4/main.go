package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

var (
	kvMapRegexp      = regexp.MustCompile("([a-z]+):([^ \n]+)[ \n]*")
	hexColorRegexp   = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	passportIDRegexp = regexp.MustCompile(`^[0-9]{9}$`)
)

func main() {
	input := lib.MustReadFile("input.txt")

	list := NewPassportList(input)

	fmt.Printf("there are %d valid passwords (simple)\n", list.CountValid(false))
	fmt.Printf("there are %d valid passwords (extended)\n", list.CountValid(true))
}

type PassportList []Passport
type Passport map[string]string

func NewPassportList(input []byte) PassportList {
	rawPassports := strings.Split(string(input), "\n\n")
	list := make(PassportList, 0, len(rawPassports))
	for _, raw := range rawPassports {
		list = append(list, NewPassport(raw))
	}
	return list
}

func (pl PassportList) CountValid(extended bool) int {
	var valid int
	for _, p := range pl {
		if p.IsValid(extended) {
			valid++
		}
	}
	return valid
}

func NewPassport(input string) Passport {
	matches := kvMapRegexp.FindAllStringSubmatch(input, -1)
	p := Passport{}
	for _, match := range matches {
		p[match[1]] = match[2]
	}
	return p
}

func (p Passport) IsValid(extended bool) bool {
	if !extended {
		return p["byr"] != "" &&
			p["iyr"] != "" &&
			p["eyr"] != "" &&
			p["hgt"] != "" &&
			p["hcl"] != "" &&
			p["ecl"] != "" &&
			p["pid"] != ""
	}

	return isBetween(p["byr"], 1920, 2002) &&
		isBetween(p["iyr"], 2010, 2020) &&
		isBetween(p["eyr"], 2020, 2030) &&
		(strings.HasSuffix(p["hgt"], "cm") && isBetween(p["hgt"][:len(p["hgt"])-2], 150, 193) ||
			strings.HasSuffix(p["hgt"], "in") && isBetween(p["hgt"][:len(p["hgt"])-2], 59, 76)) &&
		hexColorRegexp.MatchString(p["hcl"]) &&
		isOneOf(p["ecl"], []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) &&
		passportIDRegexp.MatchString(p["pid"])
}

func isBetween(in string, min, max int) bool {
	num, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return num <= max && num >= min
}

func isOneOf(in string, list []string) bool {
	for _, li := range list {
		if li == in {
			return true
		}
	}
	return false
}
