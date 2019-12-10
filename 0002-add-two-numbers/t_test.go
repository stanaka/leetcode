package main

import (
	"testing"
)


func TestCase1(t *testing.T) {
	input1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	input2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	output := addTwoNumbers(input1, input2)
	if output.Val != 7 || output.Next.Val != 0 || output.Next.Next.Val != 8 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase2(t *testing.T) {
	input1 := &ListNode{5, nil}
	input2 := &ListNode{5, nil}
	output := addTwoNumbers(input1, input2)
	if output.Val != 0 || output.Next.Val != 1 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase3(t *testing.T) {
	input1 := &ListNode{1, &ListNode{8, nil}}
	input2 := &ListNode{0, nil}
	output := addTwoNumbers(input1, input2)
	if output.Val != 1 || output.Next.Val != 8 {
		t.Errorf("Error!! %v", output)
	}
}
