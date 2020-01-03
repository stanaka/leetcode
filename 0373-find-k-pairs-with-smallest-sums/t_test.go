package main

import (
	"testing"
)

type testCase struct {
	input1 []int
	input2 []int
	target int
	expect [][]int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 7, 11},
			[]int{2, 4, 6},
			3,
			[][]int{[]int{1, 2}, []int{1, 4}, []int{1, 6}},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2, 3},
			2,
			[][]int{[]int{1, 1}, []int{1, 1}},
		},
		{
			[]int{1, 2},
			[]int{3},
			2,
			[][]int{[]int{1, 3}, []int{2, 3}},
		},
		{
			[]int{1, 1},
			[]int{1, 1},
			4,
			[][]int{[]int{1, 1}, []int{1, 1}, []int{1, 1}, []int{1, 1}},
		},
		{
			[]int{},
			[]int{3, 5, 7, 9},
			1,
			[][]int{},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2, 3},
			10,
			[][]int{[]int{1, 1}, []int{1, 1}, []int{1, 2}, []int{1, 2}, []int{2, 1}, []int{1, 3}, []int{2, 2}, []int{1, 3}, []int{2, 3}},
		},
		{
			[]int{1, 3, 5},
			[]int{1, 2, 4},
			10,
			[][]int{[]int{1, 1}, []int{1, 2}, []int{3, 1}, []int{1, 4}, []int{3, 2}, []int{5, 1}, []int{3, 4}, []int{5, 2}, []int{5, 4}},
		},
		{[]int{1, 2, 4, 5, 6},
			[]int{3, 5, 7, 9},
			20,
			[][]int{[]int{1, 3}, []int{2, 3}, []int{1, 5}, []int{2, 5}, []int{4, 3}, []int{5, 3}, []int{1, 7},
				[]int{2, 7}, []int{6, 3}, []int{4, 5}, []int{1, 9}, []int{5, 5}, []int{6, 5}, []int{4, 7}, []int{2, 9},
				[]int{5, 7}, []int{6, 7}, []int{4, 9}, []int{5, 9}, []int{6, 9}},
		},
		{
			[]int{},
			[]int{},
			0,
			[][]int{},
		},
	}

	for i, v := range testCases {
		output := kSmallestPairs(v.input1, v.input2, v.target)
		if len(output) != len(v.expect) {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
			continue
		}
		for j := range output {
			if output[j][0] != v.expect[j][0] || output[j][1] != v.expect[j][1] {
				t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
				break
			}
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
