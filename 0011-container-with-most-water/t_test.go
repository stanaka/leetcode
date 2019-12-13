package main

import (
	"testing"
)

type testCase struct {
	input  []int
	expect int
}

// xn,yn
// x1,y1 -> (xn-x1) y1
// x2,y2 -> (xn-x2) y2

// (xn-x1)y1 < (xn-x2)y2
// xny1-x1y1 < xny2 - x2y2
// xn(y2-y1) > x2y2 + x1y1

// (xn-x2)y2 > (xn-x1)y1
// xn(y2-y1) > x2y2 - x1y1
// xn > (x2y2-x1y1)/(y2-y1)

// 1: {3,4}, 2: {4,5}
// xn > (4*5-3*4)/1 = 8
// 1: {2,3}, 2: {3,4}
// xn > (3*4-2*3)/1 = 6

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
		{[]int{1,2,3,4,5,6,7,8,9}, 20},
	}

	for i, v := range testCases {
		output := maxArea(v.input)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

func TestAnswer2(t *testing.T) {
	input := []int{}
	for i := 0; i < 40000; i++ {
		input = append(input, i)
	}
	expect := 399980000
	output := maxArea(input)
	if output != expect {
		t.Errorf("case: expect %v, but %v", expect, output)
	}
}
