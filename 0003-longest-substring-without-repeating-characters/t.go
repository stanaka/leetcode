package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	m := make([]int, 256)
	b := 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] > b {
			if ans < i-b {
				ans = i - b
			}
			b = m[s[i]]
		}
		m[s[i]] = i + 1
	}
	if ans < len(s)-b {
		ans = len(s) - b
	}
	return ans
}

/*
func lengthOfLongestSubstring(s string) int {
	ans := 0
	m := map[uint8]int{}
	b := 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] > b {
			if ans < i-b {
				ans = i - b
			}
			b = m[s[i]]
		}
		m[s[i]] = i + 1
	}
	if ans < len(s)-b {
		ans = len(s) - b
	}
	return ans
}
*/
