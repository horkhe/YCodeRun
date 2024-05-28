// #26. Toys on the playground
// https://coderun.yandex.ru/problem/avto?currentPage=3&pageSize=10&rowNumber=26
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	toyCount, _ := strconv.Atoi(s.Text())
	s.Scan()
	playgroundSize, _ := strconv.Atoi(s.Text())
	s.Scan()
	toyUseSequenceLen, _ := strconv.Atoi(s.Text())
	toyUseSequence := make([]int, toyUseSequenceLen)
	for i := range toyUseSequence {
		s.Scan()
		toyUseSequence[i], _ = strconv.Atoi(s.Text())
	}
	operationCount, _ := countOperations(toyCount, playgroundSize, toyUseSequence)
	fmt.Println(operationCount)
}

func countOperations(toyCount, playgroundSize int, toyUseSequence []int) (int, []int) {
	nextToyUseIndices := buildNextUseIndices(toyCount, toyUseSequence)
	pg := newPlayground(playgroundSize)
	operationCount := 0
	removes := make([]int, len(toyUseSequence))
	for i, toy := range toyUseSequence {
		nextUseIdx := nextToyUseIndices[i]
		inserted, removed := pg.upsert(toy, nextUseIdx)
		if inserted {
			operationCount++
		}
		removes[i] = removed
	}
	return operationCount, removes
}

// buildNextUseIndices makes a slice such that an element with index i contains
// index of the next usage of toySeq[i] in toySeq after i, if it is the
// last usage of toySeq[i] then it contains len(toySeq).
func buildNextUseIndices(toyCount int, toySeq []int) []int {
	nextUseIndices := make([]int, len(toySeq))
	lastSeenIndices := make([]int, toyCount+1)
	for i, toy := range toySeq {
		lastSeenIdx := lastSeenIndices[toy]
		lastSeenIndices[toy] = i
		if lastSeenIdx == 0 && toySeq[0] != toy {
			continue
		}
		nextUseIndices[lastSeenIdx] = i
	}
	for toy, lastSeenIdx := range lastSeenIndices {
		if lastSeenIdx == 0 && toySeq[0] != toy {
			continue
		}
		nextUseIndices[lastSeenIdx] = len(toySeq)
	}
	return nextUseIndices
}

type playground struct {
	size     int
	itemMap  map[int]*playgroundItem
	itemHeap []*playgroundItem
}

func newPlayground(size int) *playground {
	pg := playground{
		size:     size,
		itemMap:  make(map[int]*playgroundItem, size),
		itemHeap: make([]*playgroundItem, 0, size),
	}
	return &pg
}

func (pg *playground) upsert(toy, nextUseIdx int) (bool, int) {
	// If the toy is already on the playground, then update its next usage
	// and fix the heap structure.
	pgi := pg.itemMap[toy]
	if pgi != nil {
		pgi.nextUseIdx = nextUseIdx
		heap.Fix(pg, pgi.heapIdx)
		return false, 0
	}
	// If the playground is full then we need to remove a toy that is gonna be
	// used next the latest of all the other toys on the playground.
	var removedToy int
	if len(pg.itemHeap) == pg.size {
		removedItem := heap.Pop(pg).(*playgroundItem)
		delete(pg.itemMap, removedItem.toy)
		removedToy = removedItem.toy
	}
	// Create a new item and add it to the map and heap.
	pgi = &playgroundItem{
		toy:        toy,
		nextUseIdx: nextUseIdx,
	}
	pg.itemMap[toy] = pgi
	heap.Push(pg, pgi)
	return true, removedToy
}

func (pg *playground) Len() int {
	return len(pg.itemHeap)
}

func (pg *playground) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, nextUseIdx so we use
	// greater than here.
	return pg.itemHeap[i].nextUseIdx > pg.itemHeap[j].nextUseIdx
}

func (pg *playground) Swap(i, j int) {
	pg.itemHeap[i], pg.itemHeap[j] = pg.itemHeap[j], pg.itemHeap[i]
	pg.itemHeap[i].heapIdx = i
	pg.itemHeap[j].heapIdx = j
}

func (pg *playground) Push(x any) {
	n := len(pg.itemHeap)
	pgi := x.(*playgroundItem)
	pgi.heapIdx = n
	pg.itemHeap = append(pg.itemHeap, pgi)
}

func (pg *playground) Pop() any {
	old := pg.itemHeap
	n := len(old)
	pgi := old[n-1]
	old[n-1] = nil   // avoid memory leak
	pgi.heapIdx = -1 // for safety
	pg.itemHeap = old[0 : n-1]
	return pgi
}

type playgroundItem struct {
	toy        int
	nextUseIdx int
	heapIdx    int
}
