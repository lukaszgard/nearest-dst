package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDestinations(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		desc         string
		destinations []Destination
		expected     []Destination
	}{
		{
			desc: "compare two destinations, first one win based on duration - first criterion",
			destinations: []Destination{
				Destination{Duration: 1, Distance: 100},
				Destination{Duration: 2, Distance: 10},
			},
			expected: []Destination{
				Destination{Duration: 1, Distance: 100},
				Destination{Duration: 2, Distance: 10},
			},
		},
		{
			desc: "compare two destinations, second one win based on distance - second criterion (if time is equal)",
			destinations: []Destination{
				Destination{Duration: 1, Distance: 2},
				Destination{Duration: 1, Distance: 1},
			},
			expected: []Destination{
				Destination{Duration: 1, Distance: 1},
				Destination{Duration: 1, Distance: 2},
			},
		},
		{
			desc: "compare three destinations, last one win based on duration and compare two first based on distance",
			destinations: []Destination{
				Destination{Duration: 1.123, Distance: 2},
				Destination{Duration: 1.123, Distance: 3},
				Destination{Duration: 1.111, Distance: 1},
			},
			expected: []Destination{
				Destination{Duration: 1.111, Distance: 1},
				Destination{Duration: 1.123, Distance: 2},
				Destination{Duration: 1.123, Distance: 3},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sortDestinations(tC.destinations)
			assert.Equal(tC.expected, tC.destinations)
		})
	}
}
