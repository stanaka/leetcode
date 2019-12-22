package main

//import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func checkLength(head *ListNode, k int) bool {
	for i := 0; i < k-1; i++ {
		head = head.Next
		if head == nil {
			return false
		}
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	first := ListNode{0, head}
	var index *ListNode
	var next *ListNode
	var prev *ListNode
	var top *ListNode

	if k < 2 {
		return head
	}
	/*
		fmt.Printf("[0] ")
		for tmp := head; tmp != nil; tmp = tmp.Next {
			fmt.Printf("%v -> ", tmp.Val)
		}
		fmt.Printf("\n")
	*/

	next = head
	top = &first
	for next != nil {
		if checkLength(next, k) == false {
			break
		}

		index = next
		for i := 0; i < k; i++ {
			next = index.Next
			if i > 0 {
				index.Next = prev
			}
			prev = index
			index = next
		}
		//fmt.Printf("top:%v, prev:%v, index:%v , next:%v\n", top.Val, prev.Val, index.Val, next.Val)
		top.Next.Next = index
		top, top.Next = top.Next, prev
		//fmt.Printf("top:%v, prev:%v, index:%v , next:%v\n", top.Val, prev.Val, index.Val, next.Val)
		/*
			count := 0
			for tmp := first.Next; tmp != nil && count < 10; tmp = tmp.Next {
				fmt.Printf("%v -> ", tmp.Val)
				count++
			}
			fmt.Printf("\n")
		*/
	}

	return first.Next
}
