package main

func min(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func minFallingPathSum(a [][]int) int {
	if len(a) == 0 {
		return 0
	}
	b := make([][]int, len(a))
	ans := 1000000
	for i := len(a) - 1; i >= 0; i-- {
		b[i] = make([]int, len(a[i]))
		for j := range a[i] {
			minV := a[i][j]
			if i < len(a)-1 {
				minV += b[i+1][j]
			}
			if j > 0 {
				if i < len(a)-1 {
					minV = min(minV, a[i][j-1]+b[i+1][j-1])
				} else {
					minV = min(minV, a[i][j-1])
				}
			}
			if j < len(a[i])-1 {
				if i < len(a)-1 {
					minV = min(minV, a[i][j+1]+b[i+1][j+1])
				} else {
					minV = min(minV, a[i][j+1])
				}
			}
			b[i][j] = minV
			if i == 0 {
				ans = min(ans, b[i][j])
			}
		}
	}
	return ans
}
