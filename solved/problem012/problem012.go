// #12. Find the shortest path.
// https://coderun.yandex.ru/problem/shortest-path-length?currentPage=2&pageSize=10&rowNumber=12
package main

import (
	"fmt"
)

func main() {
	var nodeCount int
	fmt.Scan(&nodeCount)
	nodes := make([]node, nodeCount+1)
	for i := 0; i < nodeCount; i++ {
		for j := 0; j < nodeCount; j++ {
			var connected int
			fmt.Scan(&connected)
			if connected == 0 {
				continue
			}
			nodeTo := &nodes[j+1]
			nodeTo.arcsFrom = append(nodeTo.arcsFrom, i+1)
		}
	}
	var fromNodeIdx, toNodeIdx int
	fmt.Scan(&fromNodeIdx, &toNodeIdx)
	minPathLen := findMinPathLen(nodes, fromNodeIdx, toNodeIdx)
	fmt.Println(minPathLen)
}

func findMinPathLen(nodes []node, fromIdx, toIdx int) int {
	if fromIdx == toIdx {
		return 0
	}
	distance := 1
	reachable := nodes[toIdx].arcsFrom
	for len(reachable) > 0 {
		for _, currIdx := range reachable {
			curr := &nodes[currIdx]
			if currIdx == fromIdx {
				return distance
			}
			curr.visited = true
		}
		var nextReachable []int
		for _, currIdx := range reachable {
			curr := &nodes[currIdx]
			for _, nextIdx := range curr.arcsFrom {
				next := &nodes[nextIdx]
				if next.visited {
					continue
				}
				nextReachable = append(nextReachable, nextIdx)
			}
		}
		reachable = nextReachable
		distance++
	}
	return -1
}

type node struct {
	arcsFrom []int
	visited  bool
}
