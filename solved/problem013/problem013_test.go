package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestPath(t *testing.T) {
	for i, tt := range []struct {
		inNodeCount     int
		inMatrix        []int
		inFromIdx       int
		inToIdx         int
		outShortestPath []int
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
		inFromIdx:       5,
		inToIdx:         4,
		outShortestPath: []int{5, 2, 4},
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			0, 1, 0, 0, 1,
			1, 0, 1, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 0, 0, 0, 0,
		},
		inFromIdx:       3,
		inToIdx:         5,
		outShortestPath: []int{3, 2, 1, 5},
	}, {
		inNodeCount: 2,
		inMatrix: []int{
			0, 1,
			1, 0,
		},
		inFromIdx:       1,
		inToIdx:         1,
		outShortestPath: []int{1},
	}, {
		inNodeCount: 2,
		inMatrix: []int{
			1, 1,
			1, 0,
		},
		inFromIdx:       1,
		inToIdx:         1,
		outShortestPath: []int{1},
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
		inFromIdx:       1,
		inToIdx:         10,
		outShortestPath: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
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
		inFromIdx:       1,
		inToIdx:         10,
		outShortestPath: []int{1, 5, 8, 10},
	}, {
		inNodeCount: 5,
		inMatrix: []int{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
		},
		inFromIdx:       1,
		inToIdx:         5,
		outShortestPath: []int{1, 5},
	}, {
		inNodeCount: 1,
		inMatrix: []int{
			1,
		},
		inFromIdx:       1,
		inToIdx:         1,
		outShortestPath: []int{1},
	}, {
		inNodeCount: 1,
		inMatrix: []int{
			0,
		},
		inFromIdx:       1,
		inToIdx:         1,
		outShortestPath: []int{1},
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
		inFromIdx:       6,
		inToIdx:         2,
		outShortestPath: []int{6, 5, 1, 2},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := findShortestPath(tt.inNodeCount, tt.inMatrix, tt.inFromIdx, tt.inToIdx)
			assert.Equal(t, tt.outShortestPath, got)
		})
	}
}
