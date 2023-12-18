// #10. Topological sort.
// https://coderun.yandex.ru/problem/topological-sorting?currentPage=1&pageSize=10&rowNumber=10
package main

import "fmt"

func main() {
	var nodeCount, edgeCount int
	fmt.Scan(&nodeCount, &edgeCount)
	nodes := make([]node, nodeCount+1)
	for i := 0; i < edgeCount; i++ {
		var fromNodeIdx, toNodeIdx int
		fmt.Scan(&fromNodeIdx, &toNodeIdx)
		fromNode := &nodes[fromNodeIdx]
		fromNode.neighbours = append(fromNode.neighbours, toNodeIdx)
	}

	sorted := sortGraph(nodes)

	if len(sorted) == 0 {
		fmt.Println(-1)
		return
	}
	for _, nodeIdx := range sorted {
		fmt.Print(nodeIdx)
		fmt.Print(" ")
	}
	fmt.Println()
}

func sortGraph(nodes []node) []int {
	sorted := make([]int, len(nodes)-1)
	sortedPos := len(sorted) - 1
	for nodeIdx := 1; nodeIdx < len(nodes); nodeIdx++ {
		if nodes[nodeIdx].mark != markWhite {
			continue
		}
		var ok bool
		sortedPos, ok = sortGraphFrom(nodes, nodeIdx, sorted, sortedPos)
		if !ok {
			return nil
		}
	}
	return sorted
}

func sortGraphFrom(nodes []node, currIdx int, sorted []int, sortedPos int) (int, bool) {
	curr := &nodes[currIdx]
	if curr.mark == markGray {
		return -1, false
	}
	if curr.mark == markBlack {
		return sortedPos, true
	}
	curr.mark = markGray
	for _, nextIdx := range curr.neighbours {
		var ok bool
		sortedPos, ok = sortGraphFrom(nodes, nextIdx, sorted, sortedPos)
		if !ok {
			return -1, false
		}
	}
	curr.mark = markBlack
	sorted[sortedPos] = currIdx
	sortedPos--
	return sortedPos, true
}

type node struct {
	neighbours []int
	mark       int
}

const (
	markWhite = 0
	markGray  = 1
	markBlack = 2
)
