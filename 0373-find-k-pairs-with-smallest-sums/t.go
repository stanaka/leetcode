package main

import (
	"container/heap"
)

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var ans [][]int
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return ans
	}
	//fmt.Printf("nums1:%v, nums2:%v\n", nums1, nums2)
	h := &IntHeap{}
	heap.Init(h)
	for i := range nums1 {
		heap.Push(h, []int{nums1[i], nums2[0], 0})
	}
	for j := 0; j < k; j++ {
		if h.Len() == 0 {
			break
		}
		a := heap.Pop(h).([]int)
		ans = append(ans, []int{a[0], a[1]})
		if a[2]+1 < len(nums2) {
			heap.Push(h, []int{a[0], nums2[a[2]+1], a[2] + 1})
		}
	}

	return ans
}
