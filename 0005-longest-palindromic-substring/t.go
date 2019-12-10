package main

func longestPalindrome(s string) string {
	max := 1
	pos := 0
	if len(s) <= 1 {
		return s
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			t := 2
			j := 1
			for i-1-j >= 0 && i+j < len(s) && s[i+j] == s[i-1-j] {
				t += 2
				j++
			}
			if max < t {
				max = t
				pos = i - j
			}
		}
		if i+1 < len(s) && s[i+1] == s[i-1] {
			t := 3
			j := 1
			for i-j-1 >= 0 && i+j+1 < len(s) && s[i+1+j] == s[i-1-j] {
				t += 2
				j++
			}
			if max < t {
				max = t
				pos = i - j
			}
		}
	}
	return s[pos : pos+max]
}
