// #17. Conveyer
// https://coderun.yandex.ru/problem/conveyor/description?currentPage=2&pageSize=10&rowNumber=17
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	testCount, _ := strconv.Atoi(scanner.Text())
	results := make([]int, testCount)
	for i := 0; i < testCount; i++ {
		scanner.Scan()
		seqLen, _ := strconv.Atoi(scanner.Text())
		seq := make([]float64, seqLen)
		for j := 0; j < seqLen; j++ {
			scanner.Scan()
			seq[j], _ = strconv.ParseFloat(scanner.Text(), 64)
		}
		results[i] = isSortPossible(seq)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func isSortPossible(seq []float64) int {
	currPriority := -math.MaxFloat64
	stack := make([]float64, 0, len(seq))
	for seqIdx := 0; seqIdx < len(seq); seqIdx++ {
		nextPriority := seq[seqIdx]
		if nextPriority < currPriority {
			return 0
		}
		// Move priorities lower than the next priority out of stack.
		stackIdx := 0
		for ; stackIdx < len(stack); stackIdx++ {
			if stack[stackIdx] < nextPriority {
				break
			}
		}
		if stackIdx < len(stack) {
			currPriority = stack[stackIdx]
			stack = stack[:stackIdx]
		}
		// Move the next priority to the stack.
		stack = append(stack, nextPriority)
	}
	return 1
}
