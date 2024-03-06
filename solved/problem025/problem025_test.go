package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCheap(t *testing.T) {
	for i, tt := range []struct {
		inNumbers []int
		outCost   float64
	}{{
		inNumbers: []int{10, 11, 12, 13},
		outCost:   4.6,
	}, {
		inNumbers: []int{1, 1},
		outCost:   0.10,
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := addCheap(tt.inNumbers)
			assert.Equal(t, tt.outCost, got)
		})
	}
}
