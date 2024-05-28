package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOperations(t *testing.T) {
	for i, tt := range []struct {
		inPlaygroundSize  int
		inGameSeq         []int
		outOperationCount int
		outRemoves        []int
	}{{
		inPlaygroundSize:  5,
		inGameSeq:         []int{1, 2, 3, 4, 5},
		outOperationCount: 5,
		outRemoves:        []int{0, 0, 0, 0, 0},
	}, {
		inPlaygroundSize:  5,
		inGameSeq:         []int{1, 2, 3, 2, 5},
		outOperationCount: 4,
		outRemoves:        []int{0, 0, 0, 0, 0},
	}, {
		inPlaygroundSize:  5,
		inGameSeq:         []int{1, 2, 3, 2},
		outOperationCount: 3,
		outRemoves:        []int{0, 0, 0, 0},
	}, {
		inPlaygroundSize:  2,
		inGameSeq:         []int{1, 2, 3, 1, 3, 1, 2},
		outOperationCount: 4,
		outRemoves:        []int{0, 0, 2, 0, 0, 0, 3},
	}, {
		inPlaygroundSize:  2,
		inGameSeq:         []int{1, 1, 1, 1, 1, 1, 1},
		outOperationCount: 1,
		outRemoves:        []int{0, 0, 0, 0, 0, 0, 0},
	}, {
		inPlaygroundSize:  2,
		inGameSeq:         []int{1, 2, 3, 1, 1, 2, 3},
		outOperationCount: 4,
		outRemoves:        []int{0, 0, 2, 0, 0, 1, 0},
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			gotCount, gotRemoves := countOperations(10, tt.inPlaygroundSize, tt.inGameSeq)
			assert.Equal(t, tt.outOperationCount, gotCount)
			assert.Equal(t, tt.outRemoves, gotRemoves)
		})
	}
}

func TestBuildNextToyIndices(t *testing.T) {
	for i, tt := range []struct {
		in, out []int
	}{{
		in:  []int{1, 2},
		out: []int{2, 2},
	}, {
		in:  []int{1, 2, 1},
		out: []int{2, 3, 3},
	}, {
		in:  []int{1, 2, 1, 1, 1, 2, 2, 1, 2},
		out: []int{2, 5, 3, 4, 7, 6, 8, 9, 9},
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := buildNextUseIndices(10, tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}
