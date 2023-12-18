package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxRectArea(t *testing.T) {
	for i, tt := range []struct {
		in  []int
		out rect
	}{{
		in:  []int{7, 2, 1, 4, 5, 1, 3, 3},
		out: rect{4, 3, 5, 8},
	}, {
		in:  []int{7},
		out: rect{7, 0, 1, 7},
	}, {
		in:  []int{1, 2, 3, 4, 5, 6, 7},
		out: rect{4, 3, 7, 16},
	}, {
		in:  []int{7, 6, 5, 4, 3, 2, 1},
		out: rect{4, 0, 4, 16},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := maxRectArea(tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestPerformance(t *testing.T) {
	hist := make([]int, 1000000)
	for i := range hist {
		hist[i] = i
	}

	begin := time.Now()
	rect := maxRectArea(hist)
	took := time.Since(begin)

	assert.Less(t, took, 20*time.Millisecond)
	assert.Equal(t, 250000000000, rect.area)
}
