package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortGraph(t *testing.T) {
	for i, tt := range []struct {
		inNodeCount int
		inEdges     []int
		outSorted   []int
	}{{
		inNodeCount: 6,
		inEdges: []int{
			1, 2,
			3, 2,
			4, 2,
			2, 5,
			6, 5,
			4, 6,
		},
		outSorted: []int{4, 6, 3, 1, 2, 5},
	}, {
		inNodeCount: 8,
		inEdges: []int{
			1, 4,
			1, 5,
			2, 4,
			3, 5,
			3, 8,
			4, 6,
			4, 7,
			4, 8,
			5, 7,
		},
		outSorted: []int{3, 2, 1, 5, 4, 8, 7, 6},
	}, {
		inNodeCount: 3,
		outSorted:   []int{3, 2, 1},
	}, {
		inNodeCount: 3,
		inEdges: []int{
			1, 2,
			1, 2,
		},
		outSorted: []int{3, 1, 2},
	}, {
		inNodeCount: 3,
		inEdges: []int{
			1, 2,
			2, 3,
			3, 1,
		},
		outSorted: nil,
	}, {
		inNodeCount: 3,
		inEdges: []int{
			1, 1,
		},
		outSorted: nil,
	}, {
		inNodeCount: 1,
		outSorted:   []int{1},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nodes := make([]node, tt.inNodeCount+1)
			for j := 0; j < len(tt.inEdges); j += 2 {
				fromNodeIdx := tt.inEdges[j]
				toNodeIdx := tt.inEdges[j+1]
				fromNode := &nodes[fromNodeIdx]
				fromNode.neighbours = append(fromNode.neighbours, toNodeIdx)
			}
			gotSorted := sortGraph(nodes)
			assert.Equal(t, tt.outSorted, gotSorted)
		})
	}
}
