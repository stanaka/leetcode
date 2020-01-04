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
		{
			"aa",
			"a",
			false,
		},
		{
			"aa",
			"*",
			true,
		},
		{
			"cb",
			"?a",
			false,
		},
		{
			"adceb",
			"*a*b",
			true,
		},
		{
			"acdcb",
			"a*c?b",
			false,
		},
		{
			"",
			"",
			true,
		},
		{
			"",
			"a",
			false,
		},
		{
			"aab",
			"c*a*b",
			false,
		},
		{
			"aa",
			"aa",
			true,
		},
		{
			"",
			"*",
			true,
		},
		{
			"b",
			"??",
			false,
		},
		{
			"ab",
			"***a?",
			true,
		},
		{
			"bbbababbabbbbabbbbaabaaabbbbabbbababbbbababaabbbab",
			"a******b*",
			false,
		},
		{
			"abaabaaaabbabbaaabaabababbaabaabbabaaaaabababbababaabbabaabbbbaabbbbbbbabaaabbaaaaabbaabbbaaaaabbbabb",
			"ab*aaba**abbaaaa**b*b****aa***a*b**ba*a**ba*baaa*b*ab*",
			false,
		},
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
	input := ""
	words := []string{}
	for i := 0; i < 5000; i++ {
		input = input + "a"
		words = append(words, "a")
	}
	//expect := 399980000
	_ = findSubstring(input, words)
	//if output != expect {
	//	t.Errorf("case: expect %v, but %v", expect, output)
	//}
}
*/
