package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{{
		in:  "42",
		out: 42,
	}, {
		in:  "-42",
		out: -42,
	}, {
		in:  "-((-(-(42))))",
		out: -42,
	}, {
		in:  "28-42",
		out: -14,
	}, {
		in:  "(  (-( -  28) +(-(42  ) ) ) )",
		out: -14,
	}, {
		in:  "1+2-3+4-5+6-7",
		out: -2,
	}, {
		in:  "1+2-3+4-5+6-7",
		out: -2,
	}, {
		in:  "1+(2*2 - 3)",
		out: 2,
	}, {
		in:  "1+(2*-2 - 3)",
		out: -6,
	}, {
		in:  "(1+(-(-(2))  ))*-((2 - 5))",
		out: 9,
	}, {
		in:  "2 * 2 * 3 + 4 - 5 + 6 * 5 - 4 - 3",
		out: 34,
	}, {
		in:  "-2 * -2",
		out: 4,
	}, {
		in:  "2 -3*4 + 5",
		out: -5,
	}, {
		in:  "2 + 3 * 4 + 5",
		out: 19,
	}, {
		in:  "(2 + +3) * +(+4 + 5)",
		out: 45,
	}, {
		in:  "",
		out: 0,
	}, {
		in:  "-1-2*-3--4*-5",
		out: -15,
	}, {
		in:  "-1-2*-3-4-5+-6",
		out: -10,
	}, {
		in:  "-1-2*-(3-4)-5",
		out: -8,
	}, {
		in:  "-2(3)",
		out: -6,
	}, {
		in:  "+(-3(7))",
		out: -21,
	}, {
		in:  "+((5)(6))",
		out: 30,
	}, {
		in:  "+-+2(3)+(+-+3(7))+(5)(6)",
		out: 3,
	}, {
		in:  "(-2(3))",
		out: -6,
	}, {
		in:  "(+(+-+3(7)))",
		out: -21,
	}, {
		in:  "(+((5)(6)))",
		out: 30,
	}, {
		in:  "1+2(3)",
		out: 7,
	}, {
		in:  "1*2(3)",
		out: 6,
	}, {
		in:  "(1+2(3))",
		out: 7,
	}, {
		in:  "(1*2(3))",
		out: 6,
	}} {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Eval(tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestEvalErr(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out string
	}{{
		in:  "1+a+1",
		out: "evalAddSub: right operand missing at 1",
	}, {
		in:  "1 1 + 2",
		out: "invalid token '1' at 2",
	}, {
		in:  "1 ** 1",
		out: "evalMult: invalid token '*' at 3",
	}, {
		in:  "1+2-",
		out: "evalAddSub: right operand missing at 3",
	}, {
		in:  "1+2+",
		out: "evalAddSub: right operand missing at 3",
	}, {
		in:  "1+2*",
		out: "evalMult: right operand missing at 3",
	}, {
		in:  "1+2(",
		out: "evalPar: empty",
	}, {
		in:  "1+2)",
		out: "invalid token ')' at 3",
	}, {
		in:  "(1+2))",
		out: "invalid token ')' at 5",
	}, {
		in:  "(1+2)(   ",
		out: "evalPar: empty",
	}} {
		t.Run(tt.in, func(t *testing.T) {
			_, err := Eval(tt.in)
			assert.EqualError(t, err, tt.out)
		})
	}
}
