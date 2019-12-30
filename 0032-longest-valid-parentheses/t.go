package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {

	dp := make([]int, len(s)+1)
	ans := 0
	max := 2

	for i := range s {
		if s[i] == '(' {
			max++
			if i > 0 {
				for j := min(max, len(s)); j > 1; j-- {
					dp[j] = dp[j-1]
				}
			}

			if dp[0] > 0 {
				dp[1] = dp[0]
			} else {
				dp[1] = i + 1
			}
			dp[0] = 0
		} else if s[i] == ')' {
			max--
			if max < 2 {
				max = 2
			}
			if i > 0 {
				for j := 0; j < min(max, len(s)); j++ {
					dp[j] = dp[j+1]
				}
			}
			if dp[0] > 0 && i-dp[0]+2 > ans {
				ans = i - dp[0] + 2
			}
		}
		//fmt.Printf("%v\n", dp)
	}
	return ans
}
