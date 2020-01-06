package main

import (
	"testing"
)

type testCase struct {
	input  [][]int
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			12,
		},
		{
			[][]int{},
			0,
		},
		{
			[][]int{[]int{-100}},
			-100,
		},
	}

	for i, v := range testCases {
		output := minFallingPathSum(v.input)
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
