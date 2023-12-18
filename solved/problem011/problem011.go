// #11. Loop search.
// https://coderun.yandex.ru/problem/cycle-search?currentPage=2&pageSize=10&rowNumber=11
package main

import "fmt"

func main() {
	var nodeCount int
	fmt.Scan(&nodeCount)
	nodes := make([]node, nodeCount+1)
	for nodeIdx1 := 1; nodeIdx1 <= nodeCount; nodeIdx1++ {
		for nodeIdx2 := 1; nodeIdx2 <= nodeCount; nodeIdx2++ {
			var connected int
			fmt.Scan(&connected)
			if connected == 0 {
				continue
			}
			node1 := &nodes[nodeIdx1]
			node1.neighbours = append(node1.neighbours, nodeIdx2)
		}
	}

	loop := findLoop(nodes)
	if len(loop) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(len(loop))
	for _, nodeIdx := range loop {
		fmt.Print(nodeIdx, " ")
	}
}

func buildGraph(nodeCount int, matrix []int) []node {
	nodes := make([]node, nodeCount+1)
	for i := 0; i < nodeCount; i++ {
		for j := 0; j < nodeCount; j++ {
			if matrix[i*nodeCount+j] == 0 {
				continue
			}
			nodeIdx1 := i + 1
			nodeIdx2 := j + 1
			nodes[nodeIdx1].neighbours = append(nodes[nodeIdx1].neighbours, nodeIdx2)
		}
	}
	return nodes
}

func findLoop(nodes []node) []int {
	for i := 1; i < len(nodes); i++ {
		if nodes[i].visited {
			continue
		}
		path := findLoopFrom(nodes, -1, i)
		if len(path) > 0 {
			loopStartIdx := path[0]
			for j := 1; j < len(path); j++ {
				if path[j] == loopStartIdx {
					return path[1 : j+1]
				}
			}
			continue
		}
	}
	return nil
}

func findLoopFrom(nodes []node, fromIdx, currIdx int) []int {
	curr := &nodes[currIdx]
	if curr.visited {
		return []int{currIdx}
	}
	curr.visited = true
	for _, nextIdx := range curr.neighbours {
		if nextIdx == fromIdx {
			continue
		}
		path := findLoopFrom(nodes, currIdx, nextIdx)
		if len(path) > 0 {
			path = append(path, currIdx)
			return path
		}
	}
	return nil
}

type node struct {
	neighbours []int
	visited    bool
}
