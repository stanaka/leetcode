package main

import (
	"testing"
)

type testCase struct {
	input  []int
	expect []int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]int{7, 4, 9},
			[]int{1, 2},
		},
		{
			[]int{6, 5, 4, 3, 10},
			[]int{2, 3},
		},
		{
			[]int{100, 101, 104, 102, 103},
			[]int{0, 0},
		},
		{
			[]int{8, 7, 6, 5, 2},
			[]int{2, 2},
		},
		{
			[]int{8, 7, 6, 5, 1000000000},
			[]int{2, 999999991},
		},
		{
			[]int{},
			[]int{0, 0},
		},
	}

	for i, v := range testCases {
		output := numMovesStonesII(v.input)
		if output[0] != v.expect[0] {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
		if output[1] != v.expect[1] {
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
