// #13. Path in a graph.
// https://coderun.yandex.ru/problem/the-path-in-the-graph?currentPage=2&pageSize=10&rowNumber=13
package main

import "fmt"

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
			curr := &nodes[i+1]
			curr.arcsFrom = append(curr.arcsFrom, j+1)
		}
	}
	var fromIdx, toIdx int
	fmt.Scan(&fromIdx, &toIdx)

	distance := findShortestPathIter(nodes, fromIdx, toIdx)

	fmt.Println(distance)
	if distance < 1 {
		return
	}
	currIdx := fromIdx
	for i := 0; i < distance; i++ {
		fmt.Print(currIdx, " ")
		currIdx = nodes[currIdx].pathTo
	}
	fmt.Print(toIdx)
}

func findShortestPath(nodeCount int, matrix []int, fromIdx, toIdx int) []int {
	nodes := make([]node, nodeCount+1)
	for i := 0; i < nodeCount; i++ {
		for j := 0; j < nodeCount; j++ {
			if matrix[i*nodeCount+j] == 0 {
				continue
			}
			toNode := &nodes[j+1]
			toNode.arcsFrom = append(toNode.arcsFrom, i+1)
		}
	}
	distance := findShortestPathIter(nodes, fromIdx, toIdx)
	if distance < 0 {
		return nil
	}
	var path []int
	currIdx := fromIdx
	for i := 0; i < distance; i++ {
		path = append(path, currIdx)
		currIdx = nodes[currIdx].pathTo
	}
	path = append(path, toIdx)
	return path
}

func findShortestPathIter(nodes []node, fromIdx, toIdx int) int {
	if fromIdx == toIdx {
		return 0
	}
	nodes[toIdx].visited = true
	reachable := nodes[toIdx].arcsFrom
	distance := 1
	for len(reachable) > 0 {
		for _, currIdx := range reachable {
			if currIdx == fromIdx {
				return distance
			}
			curr := &nodes[currIdx]
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
				next.pathTo = currIdx
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
	pathTo   int
	visited  bool
}
