package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkDistancesFrom(t *testing.T) {
	for i, tt := range []struct {
		inLineCount  int
		inColCount   int
		inFromPos    position
		outDistances []int
	}{{
		inLineCount: 8,
		inColCount:  8,
		inFromPos:   position{3, 4},
		outDistances: []int{
			3, 2, 3, 2, 3, 2, 3, 2,
			2, 3, 4, 1, 2, 1, 4, 3,
			3, 2, 1, 2, 3, 2, 1, 2,
			2, 3, 2, 3, -1, 3, 2, 3,
			3, 2, 1, 2, 3, 2, 1, 2,
			2, 3, 4, 1, 2, 1, 4, 3,
			3, 2, 3, 2, 3, 2, 3, 2,
			4, 3, 2, 3, 2, 3, 2, 3},
	}, {
		inLineCount: 8,
		inColCount:  8,
		inFromPos:   position{0, 0},
		outDistances: []int{
			-1, 3, 2, 3, 2, 3, 4, 5,
			3, 4, 1, 2, 3, 4, 3, 4,
			2, 1, 4, 3, 2, 3, 4, 5,
			3, 2, 3, 2, 3, 4, 3, 4,
			2, 3, 2, 3, 4, 3, 4, 5,
			3, 4, 3, 4, 3, 4, 5, 4,
			4, 3, 4, 3, 4, 5, 4, 5,
			5, 4, 5, 4, 5, 4, 5, 6},
	}, {
		inLineCount: 2,
		inColCount:  2,
		inFromPos:   position{0, 0},
		outDistances: []int{
			-1, 0,
			0, 0,
		},
	}, {
		inLineCount: 2,
		inColCount:  3,
		inFromPos:   position{0, 0},
		outDistances: []int{
			-1, 0, 0,
			0, 0, 1,
		},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			b := newBoard(tt.inLineCount, tt.inColCount)
			b.markDistancesFrom(tt.inFromPos)
			assert.Equal(t, tt.outDistances, b.squares)
		})
	}
}
