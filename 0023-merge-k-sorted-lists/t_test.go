package main

import (
	"testing"
)

type testCase struct {
	input  []*ListNode
	expect *ListNode
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			[]*ListNode{&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
				&ListNode{2, &ListNode{6, nil}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
		},
		{
			[]*ListNode{&ListNode{1, nil}},
			&ListNode{1, nil},
		},
		{
			[]*ListNode{},
			nil,
		},
	}

	for i, v := range testCases {
		output := mergeKLists(v.input)
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
