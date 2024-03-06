package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrangeTrains(t *testing.T) {
	for i, tt := range []struct {
		inDeadEndCount  int
		inSchedule      []train
		outArrangements []int
		outFailedOn     int
	}{{
		inDeadEndCount:  1,
		inSchedule:      []train{{2, 5}},
		outArrangements: []int{1},
	}, {
		inDeadEndCount: 1,
		inSchedule:     []train{{2, 5}, {5, 6}},
		outFailedOn:    2,
	}, {
		inDeadEndCount:  1,
		inSchedule:      []train{{2, 4}, {5, 6}},
		outArrangements: []int{1, 1},
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			gotArrangements, gotFailedOn := arrangeTrains(tt.inDeadEndCount, tt.inSchedule)
			assert.Equal(t, tt.outArrangements, gotArrangements)
			assert.Equal(t, tt.outFailedOn, gotFailedOn)
		})
	}
}
