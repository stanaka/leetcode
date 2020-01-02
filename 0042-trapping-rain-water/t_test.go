package main

import (
	"testing"
)

type testCase struct {
	input  []int
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{4, 3, 2, 1, 0, 1, 2, 3},
			9,
		},
		{
			[]int{4, 3, 2, 1, 0, 3, 0, 1, 0, 2},
			11,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{2, 1, 1, 3},
			2,
		},
	}

	for i, v := range testCases {
		output := trap(v.input)
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
