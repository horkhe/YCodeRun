package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLoop(t *testing.T) {
	for i, tt := range []struct {
		inNodeCount int
		inMatrix    []int
		outLoop     []int
	}{{
		inNodeCount: 3,
		inMatrix: []int{
			0, 1, 1,
			1, 0, 1,
			1, 1, 0,
		},
		outLoop: []int{3, 2, 1},
	}, {
		inNodeCount: 4,
		inMatrix: []int{
			0, 0, 1, 0,
			0, 0, 0, 1,
			1, 0, 0, 0,
			0, 1, 0, 0,
		},
		outLoop: nil,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			0, 1, 0, 0, 0,
			1, 0, 0, 0, 0,
			0, 0, 0, 1, 1,
			0, 0, 1, 0, 1,
			0, 0, 1, 1, 0,
		},
		outLoop: []int{5, 4, 3},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nodes := buildGraph(tt.inNodeCount, tt.inMatrix)
			loop := findLoop(nodes)
			assert.Equal(t, tt.outLoop, loop)
		})
	}
}
