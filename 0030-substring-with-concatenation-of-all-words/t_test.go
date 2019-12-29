package main

import (
	"testing"
)

type testCase struct {
	input  string
	words  []string
	expect []int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
		{
			"",
			[]string{},
			[]int{},
		},
		{
			"barfoobarfoobarbarman",
			[]string{"foo", "bar", "bar"},
			[]int{0, 6, 9},
		},
		{
			"a",
			[]string{"a"},
			[]int{0},
		},
		{
			"aaa",
			[]string{"a", "a"},
			[]int{0, 1},
		},
	}

	for i, v := range testCases {
		output := findSubstring(v.input, v.words)
		expect := v.expect
		if len(output) != len(expect) {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

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
