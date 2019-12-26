package main

import (
	"testing"
)

type testCase struct {
	input  *ListNode
	target int
	expect *ListNode
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			5,
			&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}},
		},
		{
			&ListNode{1, nil},
			0,
			&ListNode{1, nil},
		},
		{
			&ListNode{1, nil},
			1,
			nil,
		},
		{
			nil,
			2,
			nil,
		},
	}

	for i, v := range testCases {
		output := removeNthFromEnd(v.input, v.target)
		expect := v.expect
		for output != nil && expect != nil {
			if output.Val == expect.Val {
				output = output.Next
				expect = expect.Next
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
