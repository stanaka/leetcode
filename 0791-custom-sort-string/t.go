package main

func customSortString(s string, t string) string {
	f := make([]bool, 26)
	c := make([]int, 26)
	ans := ""
	for i := range s {
		f[s[i]-'a'] = true
	}
	for i := range t {
		if f[t[i]-'a'] {
			c[t[i]-'a']++
		} else {
			ans += string(t[i])
		}
	}
	u := ""
	for i := range s {
		for j := 0; j < c[s[i]-'a']; j++ {
			u += string(s[i])
		}
	}

	return u + ans
}
