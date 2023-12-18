// #4. Knight move.
// https://coderun.yandex.ru/problem/knight-move?currentPage=1&pageSize=10&rowNumber=4
package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	pathCount := countPaths(n, m)
	fmt.Print(pathCount)
}

func countPaths(n, m int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = -1
		}
	}
	maxI := n - 1
	maxJ := m - 1
	grid[maxI][maxJ] = 1
	pathCount := countPathsRec(grid, maxI, maxJ, 0, 0)

	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		fmt.Printf("%03d ", grid[i][j])
	//	}
	//	fmt.Println()
	//}

	return pathCount
}

func countPathsRec(grid [][]int, maxI, maxJ, i, j int) int {
	c := grid[i][j]
	if c != -1 {
		return c
	}
	c = 0
	i2 := i + 2
	j2 := j + 1
	if i2 <= maxI && j2 <= maxJ {
		c += countPathsRec(grid, maxI, maxJ, i2, j2)
	}
	i3 := i + 1
	j3 := j + 2
	if i3 <= maxI && j3 <= maxJ {
		c += countPathsRec(grid, maxI, maxJ, i3, j3)
	}
	grid[i][j] = c
	return c
}
