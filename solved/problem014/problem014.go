// #14. Flees.
// https://coderun.yandex.ru/problem/fleas/description?currentPage=2&pageSize=10&rowNumber=14
package main

import "fmt"

func main() {
	var lineCount, colCount int
	fmt.Scan(&lineCount, &colCount)
	var feederPos position
	fmt.Scan(&feederPos.line, &feederPos.col)
	feederPos.line--
	feederPos.col--
	var fleeCount int
	fmt.Scan(&fleeCount)
	fleePositions := make([]position, fleeCount)
	for i := 0; i < fleeCount; i++ {
		fleePos := &fleePositions[i]
		fmt.Scan(&fleePos.line, &fleePos.col)
		fleePos.line--
		fleePos.col--
	}

	totalDistance := evalFleePathDistance(lineCount, colCount, feederPos, fleePositions)

	fmt.Println(totalDistance)
}

func evalFleePathDistance(lineCount, colCount int, feederPos position, fleePositions []position) int {
	b := newBoard(lineCount, colCount)
	b.markDistancesFrom(feederPos)
	totalDistance := 0
	for _, fleePos := range fleePositions {
		fleeDistance := b.get(fleePos)
		if fleeDistance == 0 {
			return -1
		}
		if fleeDistance == -1 {
			continue
		}
		totalDistance += fleeDistance
	}
	return totalDistance
}

type board struct {
	lineCount int
	colCount  int
	squares   []int
}

func newBoard(lineCount, colCount int) *board {
	b := board{
		lineCount: lineCount,
		colCount:  colCount,
		squares:   make([]int, lineCount*colCount),
	}
	return &b
}

func (b *board) markDistancesFrom(fromPos position) {
	b.set(fromPos, -1)
	positions := b.appendAndSetReachableFrom(nil, fromPos, 1)
	distance := 2
	for len(positions) > 0 {
		var reachable []position
		for _, curr := range positions {
			reachable = b.appendAndSetReachableFrom(reachable, curr, distance)
		}
		distance++
		positions = reachable
	}
}

func (b *board) appendAndSetReachableFrom(positions []position, fromPos position, distance int) []position {
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line - 2, fromPos.col + 1}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line - 1, fromPos.col + 2}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line + 1, fromPos.col + 2}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line + 2, fromPos.col + 1}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line + 2, fromPos.col - 1}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line + 1, fromPos.col - 2}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line - 1, fromPos.col - 2}, distance)
	positions = b.appendAndSetIfInBounds(positions, position{fromPos.line - 2, fromPos.col - 1}, distance)
	return positions
}

func (b *board) appendAndSetIfInBounds(positions []position, pos position, distance int) []position {
	if b.isOutOfBounds(pos) {
		return positions
	}
	if b.get(pos) != 0 {
		return positions
	}
	b.set(pos, distance)
	positions = append(positions, pos)
	return positions
}

func (b *board) get(pos position) int {
	return b.squares[pos.line*b.colCount+pos.col]
}

func (b *board) set(pos position, val int) {
	b.squares[pos.line*b.colCount+pos.col] = val
}

func (b *board) isOutOfBounds(pos position) bool {
	if pos.line < 0 {
		return true
	}
	if pos.line >= b.lineCount {
		return true
	}
	if pos.col < 0 {
		return true
	}
	if pos.col >= b.colCount {
		return true
	}
	return false
}

type position struct {
	line, col int
}
