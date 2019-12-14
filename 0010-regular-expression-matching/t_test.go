package main

import (
	"testing"
)

type testCase struct {
	input   string
	pattern string
	expect  bool
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"ab", "ab.*", true},
		{"aaaab", "a*ab", true},
		{"ab", "a*ab", true},
		{"b", "a*ab", false},
		{"b", "a*b", true},
		{"aaa", "ab*a*c*a", true},
		{"a", "", false},
		{"", "a", false},
		{"", "", true},
	}

	for i, v := range testCases {
		output := isMatch(v.input, v.pattern)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

/*
func TestAnswer2(t *testing.T) {
	input := []int{}
	for i := 0; i < 40000; i++ {
		input = append(input, i)
	}
	expect := 399980000
	output := maxArea(input)
	if output != expect {
		t.Errorf("case: expect %v, but %v", expect, output)
	}
}
*/
