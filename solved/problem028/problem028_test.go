package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInsertPos(t *testing.T) {
	for i, tt := range []struct {
		inSeq  []int
		inNum  int
		outPos int
	}{
		{inSeq: []int{}, inNum: 1, outPos: 0},
		{inSeq: []int{2}, inNum: 1, outPos: 0},
		{inSeq: []int{2}, inNum: 2, outPos: 0},
		{inSeq: []int{2}, inNum: 3, outPos: 1},
		{inSeq: []int{2, 4}, inNum: 1, outPos: 0},
		{inSeq: []int{2, 4}, inNum: 2, outPos: 0},
		{inSeq: []int{2, 4}, inNum: 3, outPos: 1},
		{inSeq: []int{2, 4}, inNum: 4, outPos: 1},
		{inSeq: []int{2, 4}, inNum: 5, outPos: 2},
		{inSeq: []int{2, 4, 6}, inNum: 1, outPos: 0},
		{inSeq: []int{2, 4, 6}, inNum: 2, outPos: 0},
		{inSeq: []int{2, 4, 6}, inNum: 3, outPos: 1},
		{inSeq: []int{2, 4, 6}, inNum: 4, outPos: 1},
		{inSeq: []int{2, 4, 6}, inNum: 5, outPos: 2},
		{inSeq: []int{2, 4, 6}, inNum: 6, outPos: 2},
		{inSeq: []int{2, 4, 6}, inNum: 7, outPos: 3},
		{inSeq: []int{2, 4, 6, 8}, inNum: 1, outPos: 0},
		{inSeq: []int{2, 4, 6, 8}, inNum: 2, outPos: 0},
		{inSeq: []int{2, 4, 6, 8}, inNum: 3, outPos: 1},
		{inSeq: []int{2, 4, 6, 8}, inNum: 4, outPos: 1},
		{inSeq: []int{2, 4, 6, 8}, inNum: 5, outPos: 2},
		{inSeq: []int{2, 4, 6, 8}, inNum: 6, outPos: 2},
		{inSeq: []int{2, 4, 6, 8}, inNum: 7, outPos: 3},
		{inSeq: []int{2, 4, 6, 8}, inNum: 8, outPos: 3},
		{inSeq: []int{2, 4, 6, 8}, inNum: 9, outPos: 4},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := findInsertPos(tt.inSeq, tt.inNum)
			assert.Equal(t, tt.outPos, got)
		})
	}
}

func TestFindLongestAscendingSubseq(t *testing.T) {
	for i, tt := range []struct {
		in  []int
		out []int
	}{{
		in:  []int{3, 29, 5, 5, 28, 6},
		out: []int{3, 5, 28},
	}, {
		in:  []int{4, 8, 2, 6, 2, 10, 6, 29, 58, 9},
		out: []int{4, 8, 10, 29, 58},
	}, {
		in:  []int{1, 2, 3, 4, 5, 6},
		out: []int{1, 2, 3, 4, 5, 6},
	}, {
		in:  []int{6, 5, 4, 3, 2, 1},
		out: []int{6},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := findLongestAscendingSubseq(tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}
