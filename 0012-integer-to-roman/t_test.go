package main

import (
	"testing"
)

type testCase struct {
	input  int
	expect string
}


func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{3, `III`},
		{4, `IV` },
		{9,`IX` },
		{40,`XL` },
		{58, `LVIII` },
		{1994, `MCMXCIV`},
		{1999,`MCMXCIX`},
	}

	for i, v := range testCases {
		output := intToRoman(v.input)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
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