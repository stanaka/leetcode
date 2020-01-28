package main

import (
	"sort"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func numMovesStonesII(stones []int) []int {
	t := len(stones)
	sort.Ints(stones)
	//fmt.Printf("%v\n", stones)
	if t < 3 {
		return []int{0, 0}
	}
	r := 1
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	min := MaxInt
	v1 := max(stones[t-1]-stones[1]-t+2,
		stones[t-2]-stones[0]-t+2)
	for l := 0; l < t; l++ {
		for r < t && stones[r]-stones[l] < t {
			r++
		}
		v2 := t - (r - 1 - l + 1)
		//fmt.Printf("l: %d [%d], r-1: %d [%d], v1:%d, v2:%d\n",
		//	l, stones[l], r-1, stones[r-1], v1, v2)
		if (l == 0 || r == t) && v2 == 1 && t-1 == (stones[r-1]-stones[l]+1) {
			continue
		}
		if v2 < min {
			//fmt.Printf("min %d->%d: l: %d [%d], r-1: %d [%d], v:%d\n",
			//	min, v2, l, stones[l], r-1, stones[r-1], v2)
			min = v2
		}
		if r == t && t >= (stones[r-1]-stones[l]+1) {
			break
		}
	}

	return []int{min, v1}
}
