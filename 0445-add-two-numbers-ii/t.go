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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lenL1 := 0
	lenL2 := 0
	pL1 := l1
	for pL1 != nil {
		lenL1++
		pL1 = pL1.Next
	}
	pL2 := l2
	for pL2 != nil {
		lenL2++
		pL2 = pL2.Next
	}

	pL1 = l1
	pL2 = l2
	base := &ListNode{0, nil}
	ans := base

	flag := false
	pF := base

	for lenL1 > 0 || lenL2 > 0 {
		//fmt.Printf("lenL1: %d, lenL2: %d, ans: %v\n", lenL1, lenL2, ans)
		if lenL1 > lenL2 {
			ans.Next = &ListNode{pL1.Val, nil}
			pL1 = pL1.Next
			lenL1--
		} else if lenL1 < lenL2 {
			ans.Next = &ListNode{pL2.Val, nil}
			pL2 = pL2.Next
			lenL2--
		} else {
			ans.Next = &ListNode{pL1.Val + pL2.Val, nil}
			pL1 = pL1.Next
			pL2 = pL2.Next
			lenL1--
			lenL2--
		}
		if ans.Next.Val > 9 {
			//fmt.Printf("flag!! %d\n", ans.Next.Val)
			ans.Next.Val -= 10
			for pF != nil && pF != ans.Next {
				pF.Val++
				if pF.Val == 10 {
					pF.Val = 0
				}
				pF = pF.Next
			}
		} else if ans.Next.Val == 9 && flag == false {
			flag = true
			pF = ans
		} else if ans.Next.Val < 9 {
			if flag == true {
				flag = false
			}
			pF = ans.Next
		}

		ans = ans.Next
	}
	if base.Val == 0 {
		return base.Next
	}
	return base
}
