// #3. Print out the most expensive path.
// https://coderun.yandex.ru/problem/print-the-route-of-the-maximum-cost?currentPage=1&pageSize=10&rowNumber=3
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n < 1 || m < 1 {
		fmt.Println(0)
		fmt.Println("")
		return
	}
	grid := make([][]cell, n)
	for y := 0; y < n; y++ {
		grid[y] = make([]cell, m)
		for x := 0; x < m; x++ {
			fmt.Scan(&grid[y][x].cost)
		}
	}
	cost, path := findPath(grid)
	fmt.Println(cost)
	fmt.Println(path)
}

func findPath(grid [][]cell) (int, string) {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	grid[maxY][maxX].pathCost = grid[maxY][maxX].cost
	grid[maxY][maxX].pathFound = true
	cost := findPathRec(grid, maxX, maxY, 0, 0)

	//for y := range grid {
	//	for _, c := range grid[y] {
	//		fmt.Printf("%d-%c-%02d  ", c.cost, c.pathDir, c.pathCost)
	//	}
	//	fmt.Println("")
	//}

	var b strings.Builder
	for x, y := 0, 0; x < maxX || y < maxY; {
		if x != 0 || y != 0 {
			b.WriteString(" ")
		}
		dir := grid[y][x].pathDir
		b.WriteRune(dir)
		if dir == 'D' {
			y += 1
		} else {
			x += 1
		}
	}
	return cost, b.String()
}

func findPathRec(grid [][]cell, dimX, dimY, x, y int) int {
	c := &grid[y][x]
	if c.pathFound {
		return c.pathCost
	}
	downPathCost := -1
	if y < dimY {
		downPathCost = findPathRec(grid, dimX, dimY, x, y+1)
	}
	rightPathCost := -1
	if x < dimX {
		rightPathCost = findPathRec(grid, dimX, dimY, x+1, y)
	}
	c.pathFound = true
	if downPathCost > rightPathCost {
		c.pathCost = downPathCost + c.cost
		c.pathDir = 'D'
	} else {
		c.pathCost = rightPathCost + c.cost
		c.pathDir = 'R'
	}
	return c.pathCost
}

type cell struct {
	cost      int
	pathFound bool
	pathCost  int
	pathDir   rune
}
