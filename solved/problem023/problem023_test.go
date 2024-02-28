package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{{
		in: "7\n" +
			"+ 1\n" +
			"+ 2\n" +
			"-\n" +
			"+ 3\n" +
			"+ 4\n" +
			"-\n" +
			"-\n",
		out: "" +
			"1\n" +
			"2\n" +
			"3\n",
	}, {
		in: "8\n" +
			"+ 1\n" +
			"+ 2\n" +
			"+ 3\n" +
			"* 4\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n",
		out: "" +
			"1\n" +
			"2\n" +
			"4\n" +
			"3\n",
	}, {
		in: "18\n" +
			"+ 1\n" +
			"* 2\n" +
			"+ 3\n" +
			"* 4\n" +
			"+ 5\n" +
			"* 6\n" +
			"+ 7\n" +
			"* 8\n" +
			"+ 9\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n" +
			"-\n",
		out: "" +
			"1\n" +
			"2\n" +
			"4\n" +
			"6\n" +
			"8\n" +
			"3\n" +
			"5\n" +
			"7\n" +
			"9\n",
	}} {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			var buf bytes.Buffer
			process(strings.NewReader(tt.in), &buf)
			assert.Equal(t, tt.out, buf.String())
		})
	}
}
