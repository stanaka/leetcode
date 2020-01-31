package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	f := map[int]int{}
	for i := head; i != nil; i = i.Next {
		f[i.Val]++
	}
	ans := &ListNode{0, nil}
	j := ans
	for i := head; i != nil; i = i.Next {
		if f[i.Val] > 1 {
			continue
		}
		j.Next = &ListNode{i.Val, nil}
		j = j.Next
	}
	return ans.Next
}
