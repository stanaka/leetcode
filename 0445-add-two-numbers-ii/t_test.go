package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input1 *ListNode
	input2 *ListNode
	expect *ListNode
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{
			&ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{8, &ListNode{0, &ListNode{7, nil}}}},
		},
		{
			nil,
			nil,
			nil,
		},
		{
			&ListNode{0, nil},
			&ListNode{0, nil},
			&ListNode{0, nil},
		},
		{
			&ListNode{9, nil},
			&ListNode{9, nil},
			&ListNode{1, &ListNode{8, nil}},
		},
		{
			&ListNode{9, &ListNode{5, nil}},
			&ListNode{5, nil},
			&ListNode{1, &ListNode{0, &ListNode{0, nil}}},
		},
		{
			&ListNode{1, &ListNode{9, &ListNode{5, nil}}},
			&ListNode{5, nil},
			&ListNode{2, &ListNode{0, &ListNode{0, nil}}},
		},
		{
			&ListNode{9, &ListNode{5, nil}},
			&ListNode{9, &ListNode{5, nil}},
			&ListNode{1, &ListNode{9, &ListNode{0, nil}}},
		},
		{
			&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			&ListNode{9, nil},
			&ListNode{1, &ListNode{0, &ListNode{0, &ListNode{8, nil}}}},
		},
	}

	for i, v := range testCases {
		output := addTwoNumbers(v.input1, v.input2)
		e := v.expect
		o := output
		var se, so string
		for e != nil {
			se += fmt.Sprintf("%d,", e.Val)
			e = e.Next
		}
		for o != nil {
			so += fmt.Sprintf("%d,", o.Val)
			o = o.Next
		}
		e = v.expect
		o = output

		if v.expect == nil {
			if output != nil {
				t.Errorf("case[%d]: expect %v, but %v", i+1, se, so)
			} else {
				continue
			}
		}
		for {
			if e == nil {
				if o != nil {
					t.Errorf("case[%d]: expect %v, but %v", i+1, se, so)
				} else {
					break
				}
			}
			if o == nil || o.Val != e.Val {
				t.Errorf("case[%d]: expect %v, but %v", i+1, se, so)
				t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
				break
			}
			o = o.Next
			e = e.Next
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
