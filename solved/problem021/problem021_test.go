package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseError(t *testing.T) {
	for _, tt := range []struct {
		in, error string
	}{{
		in:    "<abc>x",
		error: "tokenize error at 5: want=<, got=x",
	}, {
		in:    "<abc>/",
		error: "tokenize error at 5: want=<, got=/",
	}, {
		in:    "<abc>>",
		error: "tokenize error at 5: want=<, got=>",
	}, {
		in:    "<abc><<",
		error: "tokenize error at 6: want=/|[a-z], got=<",
	}, {
		in:    "<abc><>",
		error: "tokenize error at 6: want=/|[a-z], got=>",
	}, {
		in:    "<abc><//",
		error: "tokenize error at 7: want=[a-z], got=/",
	}, {
		in:    "<abc></>",
		error: "tokenize error at 7: want=[a-z], got=>",
	}, {
		in:    "<abc></x<",
		error: "tokenize error at 8: want=[a-z]|>, got=<",
	}, {
		in:    "<abc></x/",
		error: "tokenize error at 8: want=[a-z]|>, got=/",
	}, {
		in:    "<abc></xy<",
		error: "tokenize error at 9: want=[a-z]|>, got=<",
	}, {
		in:    "<abc></xy/",
		error: "tokenize error at 9: want=[a-z]|>, got=/",
	}, {
		in:    "<abc></xyz>q",
		error: "tokenize error at 11: want=<, got=q",
		//}, {
		//	in:    "<abc></xyz><abc>",
		//	error: "odd tag count",
		//}, {
		//	in:    "<a><b><b></a>",
		//	error: "closing tag flip",
		//}, {
		//	in:    "<a></b></b></a>",
		//	error: "opening tag flip",
		//}, {
		//	in:    "<a><b><b><b><b></a>",
		//	error: "opening/closing tag count mismatch",
		//}, {
		//	in:    "<a></b></b></b></b></a>",
		//	error: "opening/closing tag count mismatch",
		//}, {
		//	in:    "<a></b></b></b></b></a>",
		//	error: "opening/closing tag count mismatch",
		//}, {
		//	in:    "<a><b><c></c><d></d></c></a>",
		//	error: "reduce error: opening=[1]<b>, closing=[6]</c>",
		//}, {
		//	in:    "<a></a><b></b></c><c><a></a>",
		//	error: "reduce error: closing=[4]</c>",
	}} {
		t.Run(tt.in, func(t *testing.T) {
			_, _, err := parseFix([]byte(tt.in))
			assert.Equal(t, tt.error, err.Error())
		})
	}
}

func TestGenerateParse(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := generateXML("abc", 3, 5, 2)
		_, _, err := parseFix(s)
		require.NoError(t, err, string(s))
	}
}

func TestParseFix(t *testing.T) {
	for _, tt := range []struct {
		in      string
		outPos  int
		outChar byte
	}{{
		in:      "</b></ab>",
		outPos:  1,
		outChar: 'a',
	}, {
		in:      "<ab><ab></ab></b></ab></ab>",
		outPos:  14,
		outChar: 'a',
	}, {
		in:      "<a><ba>",
		outPos:  4,
		outChar: '/',
	}, {
		in:      "<ab><ab><cab><ab></ab></ab>",
		outPos:  9,
		outChar: '/',
	}, {
		in:      "<a></b>",
		outPos:  1,
		outChar: 'b',
	}, {
		in:      "<a><b></b></b><a></a>",
		outPos:  1,
		outChar: 'b',
	}, {
		in:      "<aa><bb>c/bb></aa>",
		outPos:  8,
		outChar: '<',
	}, {
		in:      "<aa><bb>>/bb></aa>",
		outPos:  8,
		outChar: '<',
	}, {
		in:      "<aa><b>></bb></aa>",
		outPos:  6,
		outChar: 'b',
	}, {
		in:      "<aa><>b></bb></aa>",
		outPos:  5,
		outChar: 'b',
	}, {
		in:      "<aa><bb><>bb></aa>",
		outPos:  9,
		outChar: '/',
	}, {
		in:      "<aa><bb></>b></aa>",
		outPos:  10,
		outChar: 'b',
	}, {
		in:      "<aa><bb></bbc</aa>",
		outPos:  12,
		outChar: '>',
	}, {
		in:      "<a><>a>",
		outPos:  4,
		outChar: '/',
	}, {
		in:      "<a/</a>",
		outPos:  2,
		outChar: '>',
	}} {
		t.Run(tt.in, func(t *testing.T) {
			gotPos, gotChar, err := parseFix([]byte(tt.in))
			require.NoError(t, err)
			assert.Equal(t, tt.outPos, gotPos)
			assert.Equal(t, tt.outChar, gotChar)
		})
	}
}

func TestParseFixAll(t *testing.T) {
	s := generateXML("abc", 3, 5, 2)
	fmt.Println(string(s))
	for i := range s {
		orig := s[i]
		for _, flip := range "</abc>" {
			if byte(flip) == orig {
				continue
			}
			s[i] = byte(flip)
			gotPos, gotFix, err := parseFix(s)
			fmt.Printf("%d:%c->%c\n", i, orig, byte(flip))
			require.NoError(t, err)
			assert.Equal(t, i, gotPos)
			assert.Equal(t, orig, gotFix)
		}
		s[i] = orig
	}
}
