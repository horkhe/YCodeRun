// #20. A histogram and a rectangle
// https://coderun.yandex.ru/problem/histogram-and-rectangle?currentPage=2&pageSize=10&rowNumber=20
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	hist := make([]int, n)
	for i := range hist {
		_ = scanner.Scan()
		hist[i], _ = strconv.Atoi(scanner.Text())
	}
	res := maxRectArea(hist)
	fmt.Println(res.area)
}

func maxRectArea(hist []int) rect {
	histLen := len(hist)
	var largestRect rect
	stack := make([]rect, 0, histLen)
	for i := 0; i < histLen; i++ {
		currCand := newRect(hist[i], i, histLen)
		// See how the current candidate compares to those in the stack.
		for j := len(stack) - 1; j >= 0; j-- {
			prevCand := stack[j]
			// If the current candidate height is greater than the previous one,
			// then it can potentially be greater than the max area, so push it
			// to the stack and proceed with the next candidate.
			if currCand.height > prevCand.height {
				break
			}
			// Remove the prev candidate from the stack.
			stack = stack[:j]
			// Since the height has reduced the prev candidate area limit
			// projection is not longer valid. The prev height needs to be
			// reduced to the current level. But first lets see is the prev
			// area broke the record.
			prevCand = newRect(prevCand.height, prevCand.begin, i)
			if prevCand.area > largestRect.area {
				largestRect = prevCand
			}
			// Update the current candidate begin position and see how it holds
			// against the new stack head.
			currCand = newRect(currCand.height, prevCand.begin, histLen)
		}
		// If the curr candidate area limit is smaller, than the max area found
		// so far, then drop it.
		if currCand.area <= largestRect.area {
			continue
		}
		// If we reached this point the candidate has proved to be valid and
		// should be pushed to stack.
		stack = append(stack, currCand)
	}
	for i := 0; i < len(stack); i++ {
		curr := stack[i]
		if curr.area > largestRect.area {
			largestRect = curr
		}
	}
	return largestRect
}

type rect struct {
	height int
	begin  int
	end    int
	area   int
}

func newRect(height, begin, end int) rect {
	return rect{height: height, begin: begin, end: end, area: (end - begin) * height}
}
