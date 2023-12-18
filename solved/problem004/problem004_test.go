package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPathCount(t *testing.T) {
	assert.Equal(t, 1, countPaths(3, 2))
	assert.Equal(t, 0, countPaths(5, 5))
	assert.Equal(t, 2, countPaths(4, 4))
	assert.Equal(t, 293930, countPaths(31, 34))
}
