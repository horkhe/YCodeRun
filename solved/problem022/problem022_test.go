package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimums(t *testing.T) {
	for i, tt := range []struct {
		inSeq        []int
		inWindowSize int
		outMinimums  []int
	}{{
		inSeq:        []int{1, 3, 2, 4, 5, 3, 1},
		inWindowSize: 3,
		outMinimums:  []int{1, 2, 2, 3, 1},
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := findMinimums(tt.inSeq, tt.inWindowSize)
			assert.Equal(t, tt.outMinimums, got)
		})
	}
}
