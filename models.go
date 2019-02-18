package main

import "github.com/julienschmidt/httprouter"

// OSRM struct represent output from OSRM API
type OSRM struct {
	Routes []struct {
		Legs []struct {
			Summary  string        `json:"summary"`
			Weight   float64       `json:"weight"`
			Duration float64       `json:"duration"`
			Steps    []interface{} `json:"steps"`
			Distance float64       `json:"distance"`
		} `json:"legs"`
		WeightName string  `json:"weight_name"`
		Weight     float64 `json:"weight"`
		Duration   float64 `json:"duration"`
		Distance   float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string    `json:"hint"`
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
	Code string `json:"code"`
}

// OSRMErr struct #TODO: still needed ?
type OSRMErr struct {
	Message string `json:"message"`
}

// Destination represent a route from Source to Destination
// which is obtain by 'getRoute'
type Destination struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

// DestinationRoute struct is used to provide nearest/closest destination
type DestinationRoute struct {
	Source             string        `json:"source"`
	SortedDestinations []Destination `json:"routes"`
}

// HTTPRoute used to represent a handler struct
type HTTPRoute struct {
	HadleFunc  httprouter.Handle
	HTTPMethod string
	Pattern    string
}

// HTTPRoutes represent slice of HTTPRoute
type HTTPRoutes []HTTPRoute

var httpRoutes = HTTPRoutes{
	HTTPRoute{
		index,
		"GET",
		"/",
	},
	HTTPRoute{
		nearestDest,
		"GET",
		"/routes",
	},
}
