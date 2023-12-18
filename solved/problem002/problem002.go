// #2. The cheapest path.
// https://coderun.yandex.ru/problem/cheapest-way?currentPage=1&pageSize=10&rowNumber=2
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&m)

	grid := make([][]cell, n)
	for y := range grid {
		grid[y] = make([]cell, m)
		for x := range grid[y] {
			grid[y][x].pathCost = -1
			fmt.Scan(&grid[y][x].weight)
		}
	}

	minCost := findMinCost(grid)

	fmt.Print(minCost)
}

func findMinCost(grid [][]cell) int {
	lastY := len(grid) - 1
	lastX := len(grid[lastY]) - 1
	grid[lastY][lastX].pathCost = grid[lastY][lastX].weight
	return findMinCostRec(grid, 0, 0)
}

func findMinCostRec(grid [][]cell, x, y int) int {
	c := grid[y][x]
	if c.pathCost >= 0 {
		return c.pathCost
	}
	var cost = math.MaxInt64
	if x2 := x + 1; x2 < len(grid[y]) {
		cost = c.weight + findMinCostRec(grid, x2, y)
	}
	if y2 := y + 1; y2 < len(grid) {
		cost2 := c.weight + findMinCostRec(grid, x, y2)
		if cost2 < cost {
			cost = cost2
		}
	}
	grid[y][x].pathCost = cost
	return cost
}

type cell struct {
	weight   int
	pathCost int
}
