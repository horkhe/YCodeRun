package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	for _, tt := range []struct {
		exp string
		res int
	}{
		{"0", 0},
		{"1", 1},
		{"!1", 0},
		{"!!1", 1},
		{"!!!1", 0},
		{"!0", 1},
		{"!!0", 0},
		{"!!!0", 1},
		{"0&0", 0},
		{"0&1", 0},
		{"1&0", 0},
		{"1&1", 1},
		{"!0&1", 1},
		{"1&!0", 1},
		{"!0&!0", 1},
		{"!0&!0", 1},
		{"!!!0&!!1", 1},
		{"(0)", 0},
		{"!(1&!1)&!!(!(!(1&1&1)))", 1},
		{"0|0", 0},
		{"0|1", 1},
		{"1|0", 1},
		{"1|1", 1},
		{"0^0", 0},
		{"0^1", 1},
		{"1^0", 1},
		{"1^1", 0},
		{"1^1", 0},
		{"1^0|!0", 1},
		{"0|0|!(1^1)", 1},
		{"1|(0&0^1)", 1},
	} {
		t.Run(tt.exp, func(t *testing.T) {
			got, err := Eval(tt.exp)
			require.NoError(t, err)
			assert.Equal(t, tt.res, got)
		})
	}
}
