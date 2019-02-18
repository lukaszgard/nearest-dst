package main

import "sort"

type routes []Destination

func (r routes) Len() int {
	return len(r)
}

func (r routes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r routes) Less(i, j int) bool {
	if r[i].Duration < r[j].Duration {
		return true
	}
	if r[i].Duration > r[j].Duration {
		return false
	}
	return r[i].Distance < r[j].Distance
}

func sortDestinations(route []Destination) {
	sort.Sort(routes(route))
}
