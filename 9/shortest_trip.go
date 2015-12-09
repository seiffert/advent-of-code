package main

func shortestTripViaAllCities(t *Trip) *Trip {
	shortestTrip := &Trip{}

	for _, c := range cities {
		if !t.VisitedCity(c) {
			newTrip := t.Copy()
			newTrip.AddCity(c)
			newTrip = shortestTripViaAllCities(newTrip)

			if shortestTrip.Distance() == 0 || newTrip.IsShorterThan(shortestTrip) {
				shortestTrip.CopyFrom(newTrip)
			}
		}
	}
	if shortestTrip.Distance() != 0 {
		return shortestTrip
	}
	return t
}

func longestTripViaAllCities(t *Trip) *Trip {
	longestTrip := &Trip{}

	for _, c := range cities {
		if !t.VisitedCity(c) {
			newTrip := t.Copy()
			newTrip.AddCity(c)
			newTrip = longestTripViaAllCities(newTrip)

			if longestTrip.Distance() == 0 || longestTrip.IsShorterThan(newTrip) {
				longestTrip.CopyFrom(newTrip)
			}
		}
	}
	if longestTrip.Distance() != 0 {
		return longestTrip
	}
	return t
}
