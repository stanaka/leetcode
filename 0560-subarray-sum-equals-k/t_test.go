package main

import (
	"testing"
)

type testCase struct {
	input  []int
	target int
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{},
			0,
			0,
		},
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{1, 1, 3, 1, 3, 1},
			5,
			3,
		},
		{
			[]int{-1, -1, 1},
			0,
			1,
		},
		{
			[]int{-1, -1, 1, 0, -1},
			0,
			4,
		},
		{
			[]int{0, 0, 0},
			0,
			6,
		},
	}

	for i, v := range testCases {
		output := subarraySum(v.input, v.target)
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
