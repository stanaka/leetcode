package main

import (
	"testing"
)

type testCase struct {
	input1 string
	input2 string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			"cba",
			"abcd",
			"cbad",
		},
		{
			"",
			"aa",
			"aa",
		},
		{
			"stuvz",
			"syztuthtfv",
			"stttuvzyhf",
		},
	}

	for i, v := range testCases {
		output := customSortString(v.input1, v.input2)
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
