package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsSortPossible(t *testing.T) {
	for i, tt := range []struct {
		inSeq            []float64
		outUnsortedCount int
	}{{
		inSeq:            []float64{},
		outUnsortedCount: 1,
	}, {
		inSeq:            []float64{2.9, 2.1},
		outUnsortedCount: 1,
	}, {
		inSeq:            []float64{5.6, 9.0, 2.0},
		outUnsortedCount: 0,
	}, {
		inSeq:            []float64{3, 0, 0, 2, 0, 3, 0, 3, 3, 2, 0, 1, 0, 0, 0, 2, 0, 0, 1, 0, 0, 0, 3, 1, 3, 2, 1, 2, 2, 2, 2, 3, 3},
		outUnsortedCount: 0,
	}, {
		inSeq:            []float64{3, 0, 0, 2, 0, 3, 0, 3, 3, 2, 0, 1, 0, 0, 0, 2, 0, 0, 1, 0, 0, 0, 3, 1, 3, 2, 1, 2, 2, 2, 2, 3, 3, 1},
		outUnsortedCount: 0,
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isSortPossible(tt.inSeq)
			assert.Equal(t, tt.outUnsortedCount, got)
		})
	}
}

func TestIsSortPossiblePerf(t *testing.T) {
	seqLen := 10000
	for j := 1; j < 10; j++ {
		seq := make([]float64, seqLen)
		for i := range seq {
			seq[i] = float64(rand.Intn(2))
		}
		//fmt.Printf("Input: %v\n", seq)

		begin := time.Now()
		isSortPossible(seq)
		took := time.Since(begin)
		if took > time.Second {
			fmt.Printf("%4d: Sort took %s\n", j, took)
		}
	}
}
