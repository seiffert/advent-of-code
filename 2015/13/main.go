package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	InputLineRegexp = regexp.MustCompile("^([^ ]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([^.]+).$")

	personsByName = map[string]*Person{}
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		l := s.Text()
		m := InputLineRegexp.FindAllStringSubmatch(l, -1)

		diff, _ := strconv.Atoi(m[0][3])
		if m[0][2] == "lose" {
			diff = -1 * diff
		}

		person(m[0][1]).SetPreference(m[0][4], diff)
	}
	for _, p := range personsByName {
		p.SetPreference("Santa", 0)
	}

	t := bestTable(&Table{})
	fmt.Printf("Best table: %s\n", t)
	fmt.Printf("Total happiness diff: %d\n", t.HappinessDifference())
}

type Person struct {
	Name        string
	Preferences []Preference
}
type Preference struct {
	Person    *Person
	Happiness int
}

func person(name string) *Person {
	if person, ok := personsByName[name]; ok {
		return person
	}
	personsByName[name] = &Person{
		Name:        name,
		Preferences: []Preference{},
	}
	return personsByName[name]
}

func (p *Person) SetPreference(name string, happiness int) {
	p.Preferences = append(p.Preferences, Preference{
		Person:    person(name),
		Happiness: happiness,
	})
}

func (p *Person) HappinessDiff(other *Person) int {
	for _, p := range p.Preferences {
		if p.Person == other {
			return p.Happiness
		}
	}
	return 0
}

type Table struct {
	persons []*Person
	byName  map[string]*Person
}

func (t *Table) AddPerson(p *Person) {
	if t.byName == nil {
		t.reset()
	}

	t.persons = append(t.persons, p)
	t.byName[p.Name] = p
}

func (t *Table) HasPerson(p *Person) bool {
	if t.byName == nil {
		t.reset()
	}
	if _, ok := t.byName[p.Name]; ok {
		return true
	}
	return false
}

func (t *Table) HappinessDifference() int {
	if len(t.persons) == 0 {
		return 0
	}

	diff := 0
	prev := t.persons[len(t.persons)-1]
	for _, p := range t.persons {
		diff += prev.HappinessDiff(p)
		diff += p.HappinessDiff(prev)
		prev = p
	}
	return diff
}

func (t *Table) String() string {
	names := []string{}
	for _, p := range t.persons {
		names = append(names, p.Name)
	}
	return strings.Join(names, " - ")
}

func (t *Table) CopyFrom(other *Table) {
	t.reset()
	for _, p := range other.persons {
		t.AddPerson(p)
	}
}

func (t *Table) reset() {
	t.byName = map[string]*Person{}
	t.persons = []*Person{}
}

func bestTable(t *Table) *Table {
	best := &Table{}
	for _, p := range personsByName {
		if !t.HasPerson(p) {
			n := &Table{}
			n.CopyFrom(t)
			n.AddPerson(p)

			n = bestTable(n)
			if n.HappinessDifference() > best.HappinessDifference() {
				best.CopyFrom(n)
			}
		}
	}
	if len(best.persons) > 0 {
		return best
	}
	return t
}
