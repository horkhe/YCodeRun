package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinPathLen(t *testing.T) {
	for i, tt := range []struct {
		inNodeCount int
		inMatrix    []int
		inFrom      int
		inTo        int
		outPathLen  int
	}{{
		inNodeCount: 10,
		inMatrix: []int{
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 0, 1, 1, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 1, 0, 0, 0, 1, 0,
			0, 1, 0, 0, 0, 0, 1, 0, 0, 0,
			0, 1, 1, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 1, 0, 0, 1,
			0, 1, 0, 1, 0, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 1, 0, 0, 0, 0, 1, 0, 0,
			0, 0, 0, 0, 1, 1, 0, 0, 0, 0,
		},
		inFrom:     5,
		inTo:       4,
		outPathLen: 2,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			0, 1, 0, 0, 1,
			1, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		inFrom:     3,
		inTo:       5,
		outPathLen: 3,
	}, {
		inNodeCount: 2,
		inMatrix: []int{
			0, 1,
			1, 0,
		},
		inFrom:     1,
		inTo:       1,
		outPathLen: 0,
	}, {
		inNodeCount: 2,
		inMatrix: []int{
			1, 1,
			1, 0,
		},
		inFrom:     1,
		inTo:       1,
		outPathLen: 0,
	}, {
		inNodeCount: 10,
		inMatrix: []int{
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 0, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 1, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 0, 1, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 1, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
		},
		inFrom:     1,
		inTo:       10,
		outPathLen: 9,
	}, {
		inNodeCount: 10,
		inMatrix: []int{
			0, 1, 0, 0, 1, 0, 0, 0, 0, 0,
			1, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 1, 0, 0, 0, 0, 0,
			1, 0, 0, 1, 0, 1, 0, 1, 0, 0,
			0, 0, 0, 0, 1, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 0, 1, 0, 0,
			0, 0, 0, 0, 1, 0, 1, 0, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 0,
		},
		inFrom:     1,
		inTo:       10,
		outPathLen: 3,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
		inFrom:     1,
		inTo:       5,
		outPathLen: 1,
	}, {
		inNodeCount: 1,
		inMatrix: []int{
			1,
		},
		inFrom:     1,
		inTo:       1,
		outPathLen: 0,
	}, {
		inNodeCount: 1,
		inMatrix: []int{
			0,
		},
		inFrom:     1,
		inTo:       1,
		outPathLen: 0,
	}, {
		inNodeCount: 3,
		inMatrix: []int{
			1, 1, 0,
			0, 1, 1,
			0, 0, 0,
		},
		inFrom:     1,
		inTo:       3,
		outPathLen: 2,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			1, 1, 0, 0, 0,
			1, 1, 1, 0, 0,
			1, 1, 1, 1, 0,
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 0,
		},
		inFrom:     1,
		inTo:       5,
		outPathLen: 4,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			0, 0, 1, 0, 1,
			1, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 0,
		},
		inFrom:     2,
		inTo:       5,
		outPathLen: 2,
	}, {
		inNodeCount: 6,
		inMatrix: []int{
			1, 1, 0, 0, 1, 0,
			1, 1, 0, 0, 0, 0,
			1, 0, 0, 0, 1, 1,
			0, 0, 1, 0, 1, 1,
			1, 0, 1, 1, 1, 0,
			0, 0, 0, 1, 1, 1,
		},
		inFrom:     6,
		inTo:       2,
		outPathLen: 3,
	}, {
		inNodeCount: 4,
		inMatrix: []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 1, 0,
			0, 1, 0, 0,
		},
		inFrom:     2,
		inTo:       3,
		outPathLen: -1,
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			0, 0, 0, 1, 0,
			1, 0, 0, 1, 0,
			1, 1, 1, 1, 1,
			0, 0, 0, 1, 1,
			0, 1, 1, 1, 1,
		},
		inFrom:     2,
		inTo:       3,
		outPathLen: 3,
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nodes := buildGraph(tt.inNodeCount, tt.inMatrix)
			gotPathLen := findMinPathLen(nodes, tt.inFrom, tt.inTo)
			assert.Equal(t, tt.outPathLen, gotPathLen)
		})
	}
}

func buildGraph(nodeCount int, matrix []int) []node {
	nodes := make([]node, nodeCount+1)
	for i := 0; i < nodeCount; i++ {
		for j := 0; j < nodeCount; j++ {
			if matrix[i*nodeCount+j] == 0 {
				continue
			}
			nodeTo := &nodes[j+1]
			nodeTo.arcsFrom = append(nodeTo.arcsFrom, i+1)
		}
	}
	return nodes
}
