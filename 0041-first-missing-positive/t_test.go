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
			[]int{1, 2, 0},
			3,
		},

		{
			[]int{3, 4, -1, 1},
			2,
		},
		{
			[]int{7, 8, 9, 11, 12},
			1,
		},
		{
			[]int{},
			1,
		},
		{
			[]int{2, 4, 6, 5, 3, 1},
			7,
		},
		{
			[]int{2, 2, 1, 4, 3},
			5,
		},
	}

	for i, v := range testCases {
		output := firstMissingPositive(v.input)
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
