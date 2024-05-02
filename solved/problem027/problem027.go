// #27 Shortest sum of cubes
// https://coderun.yandex.ru/problem/sum-of-cubes/description
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	sum := FindShortestCubeSum(n)
	//check := 0
	//for i, cube := range sum {
	//	x := int(math.Cbrt(float64(cube)))
	//	if i > 0 {
	//		fmt.Print(" + ")
	//	}
	//	fmt.Printf("%d(%d^3)", cube, x)
	//	check += cube
	//}
	//fmt.Printf(" = %d(%t)", check, check == n)
	fmt.Print(len(sum))
}

func FindShortestCubeSum(n int) []int {
	if n < 1 {
		return nil
	}
	cubes := makeCubes(n)
	sum, _ := findShortestCubeSum(n, cubes, n)
	return sum
}

func findShortestCubeSum(n int, cubes []int, maxAllowedLen int) ([]int, bool) {
	var shortestCubeSum []int
	for i := len(cubes) - 1; i >= 0; i-- {
		cube := cubes[i]
		minShortestCubeSum := n / cube
		if minShortestCubeSum > maxAllowedLen {
			if shortestCubeSum != nil {
				return shortestCubeSum, true
			}
			return nil, false
		}
		if n%cube == 0 {
			return appendN(nil, cube, minShortestCubeSum), true
		}
		for j := minShortestCubeSum; j > 0; j-- {
			remainder := n - (cube * j)
			remainderCubeSum, ok := findShortestCubeSum(remainder, cubes[:i], maxAllowedLen-j)
			if !ok {
				continue
			}
			shortestCubeSum = appendN(remainderCubeSum, cube, j)
			maxAllowedLen = len(shortestCubeSum) - 1
		}
	}
	if shortestCubeSum == nil {
		return nil, false
	}
	return shortestCubeSum, true
}

func appendN(s []int, n, count int) []int {
	for i := 0; i < count; i++ {
		s = append(s, n)
	}
	return s
}

func makeCubes(n int) []int {
	cubes := make([]int, 0, int(math.Cbrt(3)))
	for i := 1; i <= n; i++ {
		cube := i * i * i
		if cube > n {
			break
		}
		cubes = append(cubes, cube)
	}
	return cubes
}
