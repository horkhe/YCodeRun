package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestWayOut(t *testing.T) {
	for i, tt := range []struct {
		inDimension int
		inLines     []string
		outPathLen  int
	}{{
		inDimension: 3,
		inLines: []string{
			"###",
			"###",
			".##",

			".#.",
			".#S",
			".#.",

			"###",
			"...",
			"###",
		},
		outPathLen: 6,
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			c := newCave(tt.inDimension)
			for z := 0; z < tt.inDimension; z++ {
				for y := 0; y < tt.inDimension; y++ {
					c.setLine(z, y, tt.inLines[z*tt.inDimension+y])
				}
			}
			got := c.findShortestWayOut()
			assert.Equal(t, tt.outPathLen, got)
		})
	}
}
