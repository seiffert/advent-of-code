package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var (
	roomRegexp = regexp.MustCompile(`([-a-z]+)-([0-9]+)\[([^\]]*)\]$`)
)

type checkSum []*character
type character struct {
	char  rune
	count int
}

type byCharCount checkSum

func (s byCharCount) Len() int      { return len(s) }
func (s byCharCount) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byCharCount) Less(i, j int) bool {
	return s[i].count > s[j].count || s[i].count == s[j].count && s[i].char < s[j].char
}

func (cs *checkSum) add(char rune) {
	for _, c := range *cs {
		if c.char == char {
			c.count++
			return
		}
	}
	*cs = append(*cs, &character{char, 1})
}
func (cs *checkSum) calc() string {
	sort.Sort(byCharCount(*cs))
	var result string
	for _, c := range *cs {
		result += string(c.char)
		if len(result) == 5 {
			break
		}
	}
	return result
}

func parseRoom(room string) (string, int, string) {
	matches := roomRegexp.FindAllStringSubmatch(room, -1)
	sectorID, _ := strconv.Atoi(matches[0][2])
	return matches[0][1], sectorID, matches[0][3]
}

func SectorID(room string) int {
	_, sectorID, _ := parseRoom(room)
	return sectorID
}

func IsValid(room string) bool {
	room, _, checkSum := parseRoom(room)

	return check(room) == checkSum
}

func check(room string) string {
	sum := &checkSum{}
	for _, char := range room {
		if char != '-' {
			sum.add(char)
		}
	}
	return sum.calc()
}

func Decrypt(room string) string {
	room, sectorID, _ := parseRoom(room)

	var result string
	for _, r := range room {
		if r == '-' {
			result += " "
			continue
		}
		result += string(rune((int(r) - 97 + sectorID) % 26 + 97))
	}
	return result
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var sum int
	for s.Scan() {
		room := s.Text()
		if IsValid(room) {
			sum += SectorID(room)
			decryptedRoom := Decrypt(room)
			log.Printf("found room %s (%s)", decryptedRoom, room)
		}
	}
	fmt.Printf("Sum of sector IDs: %d", sum)
}
