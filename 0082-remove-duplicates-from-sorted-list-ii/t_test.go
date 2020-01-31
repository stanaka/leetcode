package main

import (
	"testing"
)

type testCase struct {
	input  *ListNode
	expect *ListNode
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			&ListNode{1, &ListNode{2,
				&ListNode{3, &ListNode{3,
					&ListNode{4, &ListNode{4,
						&ListNode{5, nil}}}}}}},
			&ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			&ListNode{1, &ListNode{1,
				&ListNode{1, &ListNode{2,
					&ListNode{3, nil}}}}},
			&ListNode{2, &ListNode{3, nil}},
		},
		{
			nil,
			nil,
		},
	}

	for i, v := range testCases {
		output := deleteDuplicates(v.input)
		expect := v.expect
		if (output != nil && expect == nil) || (output == nil && expect != nil) {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
			break
		}
		j := 1
		for output != nil && expect != nil {
			if output.Val == expect.Val {
				output = output.Next
				expect = expect.Next
			} else {
				t.Errorf("case[%d][%d]: expect %v, but %v", i+1, j, expect, output)
				break
			}
			j++
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
