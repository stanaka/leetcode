package main

import (
	"testing"
)

type testCase struct {
	input  string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`babad`, `bab`},
		{`cbbd`, `bb`},
		{`abcdedcba`, `abcdedcba`},
		{`cbbbbbbd`, `bbbbbb`},
		{``, ``},
	}

	for i, v := range testCases {
		output := longestPalindrome(v.input)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

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
