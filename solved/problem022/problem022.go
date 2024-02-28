// #22. Minimum of the segment
// https://coderun.yandex.ru/problem/minimum-of-the-segment?currentPage=3&pageSize=10&rowNumber=22
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	seqSize, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	windowSize, _ := strconv.Atoi(scanner.Text())
	seq := make([]int, seqSize)
	for i := range seq {
		scanner.Scan()
		seq[i], _ = strconv.Atoi(scanner.Text())
	}

	minimums := findMinimums(seq, windowSize)

	for i := range minimums {
		fmt.Println(minimums[i])
	}
}

func findMinimums(seq []int, windowSize int) []int {
	w := newWindow(windowSize)
	minimums := make([]int, 0, len(seq)-windowSize+1)
	var minVal int
	for i := 0; i < windowSize; i++ {
		minVal = w.push(seq[i])
	}
	minimums = append(minimums, minVal)
	for i := windowSize; i < len(seq); i++ {
		minVal = w.push(seq[i])
		minimums = append(minimums, minVal)
	}
	return minimums
}

func newWindow(size int) *window {
	w := window{
		size:  size,
		heap:  make([]item, 0, size),
		order: make([]int, size),
	}
	return &w
}

type window struct {
	size         int
	heap         []item
	order        []int
	nextOrderPos int
}

type item struct {
	val      int
	orderPos int
}

func (w *window) push(val int) int {
	if len(w.heap) < w.size {
		heap.Push(w, val)
	} else {
		heapIdx := w.order[w.nextOrderPos]
		w.heap[heapIdx].val = val
		heap.Fix(w, heapIdx)
	}
	w.nextOrderPos++
	if w.nextOrderPos == w.size {
		w.nextOrderPos = 0
	}
	return w.heap[0].val
}

func (w *window) Push(x any) {
	heapIdx := w.Len()
	item := item{x.(int), w.nextOrderPos}
	w.order[w.nextOrderPos] = heapIdx
	w.heap = append(w.heap, item)
}

func (w *window) Pop() any {
	lastIdx := len(w.heap) - 1
	item := w.heap[lastIdx]
	w.heap = w.heap[:lastIdx]
	return item.val
}

func (w *window) Len() int {
	return len(w.heap)
}

func (w *window) Less(i, j int) bool {
	return w.heap[i].val < w.heap[j].val
}

func (w *window) Swap(i, j int) {
	w.heap[i], w.heap[j] = w.heap[j], w.heap[i]
	w.order[w.heap[i].orderPos] = i
	w.order[w.heap[j].orderPos] = j
}
