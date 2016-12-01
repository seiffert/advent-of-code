package main

var (
	dists = make(map[*City]map[*City]int)
)

func dist(a, b *City) int {
	return dists[a][b]
}

func setDist(a, b *City, dist int) {
	if dists[a] == nil {
		dists[a] = make(map[*City]int)
	}
	dists[a][b] = dist

	if dists[b] == nil {
		dists[b] = make(map[*City]int)
	}
	dists[b][a] = dist
}
