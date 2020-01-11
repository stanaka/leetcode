package main

import (
	"testing"
)

type testCase struct {
	input  int
	expect []int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			2,
			[]int{0, 1, 1},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
		{
			10,
			[]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
		{
			0,
			[]int{},
		},
	}

	for i, v := range testCases {
		output := countBits(v.input)
		if len(output) != len(v.expect) {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		} else {
			for i := range output {
				if output[i] != v.expect[i] {
					t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
					break
				}
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
