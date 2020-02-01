package main

import (
	"testing"
)

type testCase struct {
	input  []string
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]string{"eca", "eab", "zac", "zba"},
			2,
		},
		{
			[]string{"eac", "eca", "zca", "zab"},
			2,
		},
		{
			[]string{"ca", "bb", "ac"},
			1,
		},
		{
			[]string{"xc", "yb", "za"},
			0,
		},
		{
			[]string{"ea", "eb", "za"},
			0,
		},
		{
			[]string{"ea", "eb", "za", "zb"},
			0,
		},
		{
			[]string{"ec", "eb", "ea"},
			1,
		},
		{
			[]string{"zyx", "wvu", "tsr"},
			3,
		},
		{
			[]string{},
			0,
		},
		{
			[]string{"doeeqiy", "yabhbqe", "twckqte"},
			3,
		},
	}

	for i, v := range testCases {
		output := minDeletionSize(v.input)
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
