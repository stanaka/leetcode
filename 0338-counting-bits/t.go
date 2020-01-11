package main

func countBits(num int) []int {
	ans := make([]int, num+1)
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0, 1}
	}
	ans[1] = 1
	idx := 0
	base := 2
	for i := 2; i < num+1; i++ {
		ans[i] = 1 + ans[idx]
		//fmt.Printf("i: %d, ans[i]:%d, base: %d, idx: %d\n", i, ans[i], base, idx)
		idx++
		if base == idx {
			idx = 0
			base = i + 1
		}
	}
	return ans
}
