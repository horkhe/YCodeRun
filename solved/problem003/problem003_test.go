package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPath(t *testing.T) {
	grid := makeGrid([][]int{
		{9, 9, 9, 9, 9},
		{3, 0, 0, 0, 0},
		{9, 9, 9, 9, 9},
		{6, 6, 6, 6, 8},
		{9, 9, 9, 9, 9},
	})
	cost, path := findPath(grid)
	assert.Equal(t, 74, cost)
	assert.Equal(t, "D D R R R R D D", path)

	grid = makeGrid([][]int{
		{9},
		{3},
		{9},
		{6},
		{9},
	})
	cost, path = findPath(grid)
	assert.Equal(t, 36, cost)
	assert.Equal(t, "D D D D", path)

	grid = makeGrid([][]int{
		{9, 9, 9, 9, 9},
	})
	cost, path = findPath(grid)
	assert.Equal(t, 45, cost)
	assert.Equal(t, "R R R R", path)

	grid = makeGrid([][]int{
		{9},
	})
	cost, path = findPath(grid)
	assert.Equal(t, 9, cost)
	assert.Equal(t, "", path)

	grid = makeGrid([][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})
	cost, path = findPath(grid)
	assert.Equal(t, 74, cost)
	assert.Equal(t, "D D R R R R D D", path)

}

func makeGrid(costGrid [][]int) [][]cell {
	cellGrid := make([][]cell, len(costGrid))
	for y := range costGrid {
		cellGrid[y] = make([]cell, len(costGrid[y]))
		for x := range costGrid[y] {
			cellGrid[y][x].cost = costGrid[y][x]
		}
	}
	return cellGrid
}
