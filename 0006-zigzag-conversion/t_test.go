package main

import (
	"testing"
)

type testCase struct {
	input1 string
	input2 int
	expect string
}

// P   A   H   N
// A P L S I I G
// Y   I   R

// P  I  N
// A LS IG
// YA HR
// P  I

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`PAYPALISHIRING`, 3, `PAHNAPLSIIGYIR`},
		{`PAYPALISHIRING`, 4, `PINALSIGYAHRPI`},
		{`A`, 1, `A`},
		{``, 4, ``},
	}

	for i, v := range testCases {
		output := convert(v.input1, v.input2)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

/*
func TestAnswer2(t *testing.T) {
	input := ``
	for i := 0; i < 1000; i++ {
		input += "b"
	}
	expect := input
	output := longestPalindrome(input)
	if output != expect {
		t.Errorf("case: expect %v, but %v", expect, output)
	}
}
*/
