package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	idx := head
	var prev *ListNode
	offset := 0
	for idx != nil {
		idx = idx.Next
		if offset < n {
			offset++
		} else {
			if prev == nil {
				prev = head
			} else {
				prev = prev.Next
			}
		}
	}
	if prev != nil {
		if prev.Next != nil && prev.Next.Next != nil {
			prev.Next = prev.Next.Next
		} else {
			prev.Next = nil
		}
	} else {
		if offset == n {
			head = head.Next
		}
	}
	return head
}
