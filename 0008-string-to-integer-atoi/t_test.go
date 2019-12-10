package main

import (
	"testing"
)

type testCase struct {
	input  string
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`42`, 42},
		{`   -42`, -42},
		{`4193 with words`, 4193},
		{`words and 987`, 0},
		{`-91283472332`, -2147483648},
		{``, 0},
		{` `, 0},
		{`+-2`, 0},
		{`0-1`, 0},
		{`9223372036854775808`, 2147483647},
		{`-2147483648`, -2147483648},
	}

	for i, v := range testCases {
		output := myAtoi(v.input)
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
