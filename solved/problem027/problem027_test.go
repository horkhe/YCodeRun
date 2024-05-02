package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestCubeSum(t *testing.T) {
	for i, tt := range []struct {
		in  int
		out []int
	}{
		{in: 0, out: nil},
		{in: 1, out: []int{1}},
		{in: 9, out: []int{1, 8}},
		{in: 50, out: []int{1, 1, 8, 8, 8, 8, 8, 8}},
		{in: 100, out: []int{1, 8, 27, 64}},
		{in: 999, out: []int{27, 27, 216, 729}},
		{in: 1000, out: []int{1000}},
		{in: 1001, out: []int{1, 1000}},
		{in: 1009, out: []int{1, 8, 1000}},
		{in: 1073, out: []int{1, 343, 729}},
		{in: 1075, out: []int{1, 1, 1, 343, 729}},
		{in: 1085, out: []int{8, 8, 125, 216, 216, 512}},
		{in: 1086, out: []int{8, 8, 125, 216, 729}},
		{in: 1087, out: []int{8, 8, 216, 343, 512}},
		{in: 1088, out: []int{64, 512, 512}},
		{in: 1089, out: []int{1, 64, 512, 512}},
		{in: 1090, out: []int{1, 1, 64, 512, 512}},
		{in: 9999, out: []int{1, 8, 729, 9261}},
		{in: 10000, out: []int{1000, 1000, 8000}},
		{in: 999993, out: []int{5832, 262144, 343000, 389017}},
		{in: 999994, out: []int{343, 3375, 4096, 79507, 912673}},
		{in: 999995, out: []int{2744, 8000, 54872, 205379, 729000}},
		{in: 999996, out: []int{13824, 42875, 238328, 704969}},
		{in: 999997, out: []int{125, 110592, 110592, 778688}},
		{in: 999998, out: []int{21952, 24389, 68921, 884736}},
		{in: 999999, out: []int{24389, 27000, 35937, 912673}},
		{in: 1000000, out: []int{1000000}},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := FindShortestCubeSum(tt.in)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestMakeCubes(t *testing.T) {
	cubes := makeCubes(1000000)
	assert.Equal(t, 100, len(cubes))
}
