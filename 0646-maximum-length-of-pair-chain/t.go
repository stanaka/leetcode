package main

import (
	"sort"
)

type bcs [][]int

func (b bcs) Len() int           { return len(b) }
func (b bcs) Less(i, j int) bool { return b[i][1] < b[j][1] }
func (b bcs) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	sort.Sort(bcs(pairs))

	ans := 1
	pos := pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if pos < pairs[i][0] {
			pos = pairs[i][1]
			ans++
		}
	}

	return ans
}
