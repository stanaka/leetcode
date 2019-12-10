package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	output := &ListNode{}
	index := output
	flag := false
	init := true
	for l1 != nil || l2 != nil || flag == true {
		if init != true {
			index.Next = &ListNode{}
			index = index.Next
		}
		init = false
		v := 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if flag == true {
			v++
			flag = false
		}
		if v > 9 {
			v -= 10
			flag = true
		}
		index.Val = v
	}
	return output
}
