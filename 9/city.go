package main

var (
	cities = make(map[string]*City)
)

type City struct {
	Name string
}

func city(name string) *City {
	if city, ok := cities[name]; ok {
		return city
	}

	cities[name] = &City{Name: name}
	return cities[name]
}
