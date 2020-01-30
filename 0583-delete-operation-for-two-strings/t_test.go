package main

import (
	"testing"
)

type testCase struct {
	input1 string
	input2 string
	expect int
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{ /*
			{
				"sea",
				"eat",
				2,
			},
			{
				"seax",
				"eayt",
				4,
			},
			{
				"",
				"",
				0,
			},
			{
				"b",
				"",
				1,
			},
			{
				"",
				"b",
				1,
			},
			{
				"sea",
				"ate",
				4,
			},
			{
				"algorithm",
				"altruistic",
				9,
			},
			{
				"dinitrophenylhydrazine",
				"acetylphenylhydrazine",
				11,
			},*/
		{
			"intention",
			"execution",
			8,
		},
	}

	for i, v := range testCases {
		output := minDistance(v.input1, v.input2)
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
