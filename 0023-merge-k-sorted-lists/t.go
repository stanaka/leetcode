package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LNHeap []*ListNode

func (h LNHeap) Len() int           { return len(h) }
func (h LNHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h LNHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LNHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *LNHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	var last *ListNode

	h := &LNHeap{}
	heap.Init(h)
	for _, ln := range lists {
		if ln != nil {
			heap.Push(h, ln)
		}
	}
	for h.Len() > 0 {
		ln := heap.Pop(h).(*ListNode)
		if ln == nil {
			break
		}
		//fmt.Printf("%d ", ln.Val)

		if ans == nil {
			ans = &ListNode{}
			ans.Val = ln.Val
			last = ans
		} else {
			last.Next = &ListNode{ln.Val, nil}
			last = last.Next
		}

		if ln.Next != nil {
			heap.Push(h, ln.Next)
		}
	}

	return ans
}
