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
		/*{
			[]int{2, 2, 3, 2},
			3,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{0, 1, 0, 1, 0, 1, 99},
			99,
		},*/
		{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2},
			-4,
		},
	}

	for i, v := range testCases {
		output := singleNumber(v.input)
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
