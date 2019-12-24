package main

import (
	"testing"
)

type testCase struct {
	input  []int
	target int
	expect [][]int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				[]int{-1, 0, 0, 1},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
			},
		},
		{
			[]int{0, 0, 0, 0},
			0,
			[][]int{
				[]int{0, 0, 0, 0},
			},
		},
		{
			[]int{-3, -2, -1, 0, 0, 1, 2, 3},
			0,
			[][]int{
				[]int{-3, -2, 2, 3}, []int{-3, -1, 1, 3}, []int{-3, 0, 0, 3}, []int{-3, 0, 1, 2}, []int{-2, -1, 0, 3}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1},
			},
		},
	}

	for i, v := range testCases {
		output := fourSum(v.input, v.target)
		expect := v.expect
		for output != nil && expect != nil {
			if len(output) == len(expect) {
				break
			} else {
				t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
				break
			}
		}
	}
}

/*
func TestAnswer2(t *testing.T) {
	input := []int{}
	for i := 0; i < 40000; i++ {
		input = append(input, i)
	}
	expect := 399980000
	output := intToRoman(input)
	if output != expect {
		t.Errorf("case: expect %v, but %v", expect, output)
	}
}
*/
