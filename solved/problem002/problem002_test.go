package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOptimalPath(t *testing.T) {
	assert.Equal(t, 11, findMinCost(makeGrid([][]int{
		{1, 1, 1, 1, 1},
		{3, 100, 100, 100, 100},
		{1, 1, 1, 1, 1},
		{2, 2, 2, 2, 1},
		{1, 1, 1, 1, 1},
	})))

	assert.Equal(t, 8, findMinCost(makeGrid([][]int{
		{1},
		{3},
		{1},
		{2},
		{1},
	})))

	assert.Equal(t, 5, findMinCost(makeGrid([][]int{
		{1, 1, 1, 1, 1},
	})))
}

func makeGrid(weights [][]int) [][]cell {
	grid := make([][]cell, len(weights))
	for y := range weights {
		grid[y] = make([]cell, len(weights[y]))
		for x, weight := range weights[y] {
			grid[y][x].pathCost = -1
			grid[y][x].weight = weight
		}
	}
	return grid
}
