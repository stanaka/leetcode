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
			[]int{2, 1, 4, 7, 3, 2, 5},
			5,
		},
		{
			[]int{2, 2, 2},
			0,
		},
		{
			[]int{2, 4, 2},
			3,
		},
		{
			[]int{2, 4, 2, 2},
			3,
		},
		{
			[]int{875, 884, 239, 731, 723, 685},
			4,
		},
		{
			[]int{},
			0,
		},
	}

	for i, v := range testCases {
		output := longestMountain(v.input)
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
