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
	numberCount, _ := strconv.Atoi(scanner.Text())
	numbers := make([]int, numberCount)
	for i := 0; i < numberCount; i++ {
		scanner.Scan()
		numbers[i], _ = strconv.Atoi(scanner.Text())
	}

	cost := addCheap(numbers)

	fmt.Printf("%.2f\n", cost)
}

func addCheap(numbers []int) float64 {
	sorted := make(intHeap, 0, len(numbers))
	for _, num := range numbers {
		heap.Push(&sorted, num)
	}
	var cost float64
	for sorted.Len() > 1 {
		n1 := heap.Pop(&sorted).(int)
		n2 := heap.Pop(&sorted).(int)
		n3 := n1 + n2
		heap.Push(&sorted, n3)
		c := float64(n3) * 0.05
		cost += c
	}
	return cost
}

type intHeap []int

func (ih *intHeap) Len() int           { return len(*ih) }
func (ih *intHeap) Less(i, j int) bool { return (*ih)[i] < (*ih)[j] }
func (ih *intHeap) Swap(i, j int)      { (*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i] }
func (ih *intHeap) Push(v any)         { (*ih) = append(*ih, v.(int)) }
func (ih *intHeap) Pop() any {
	n := len(*ih) - 1
	v := (*ih)[n]
	*ih = (*ih)[:n]
	return v
}
