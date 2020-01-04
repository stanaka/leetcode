package main

var dp [][]int

func subIsMatch(s string, p string) bool {
	switch dp[len(s)][len(p)] {
	case -1:
		return false
	case 1:
		return true
	}
	//fmt.Printf("s:%s, p:%s\n", s, p)

	/*if (s == "" && p != "" && p != "*") || s != "" && p == "" {
		dp[len(s)][len(p)] = -1
		return false
	}
	if s == "" && p == "" {
		dp[len(s)][len(p)] = 1
		return true
	}*/

	pos := len(s) - 1
	j := len(p) - 1
	for j >= 0 {
		//if pos >= 0 {
		//	fmt.Printf("pos[%d]:%v, j[%d]:%v\n", pos, string(s[pos]), j, string(p[j]))
		//} else {
		//	fmt.Printf("pos[%d]:, j[%d]:%v\n", pos, j, string(p[j]))
		//}
		if pos >= 0 && (p[j] == '?' || p[j] != '*' && p[j] == s[pos]) {
			// match!
			j--
			pos--
		} else if p[j] == '*' {
			if j == 0 {
				dp[len(s)][len(p)] = 1
				return true
			}
			j--
			for {
				//fmt.Printf("pos:%d, j:%d\n", pos, j)
				if subIsMatch(s[:pos+1], p[:j+1]) {
					dp[len(s)][len(p)] = 1
					return true
				} else {
					if pos <= 0 && j >= 0 {
						//fmt.Printf("false return\n")
						dp[len(s)][len(p)] = -1
						return false
					}
					pos--
				}
			}
		} else {
			//fmt.Printf("false return 2 pos:%d, j:%d\n", pos, j)
			dp[len(s)][len(p)] = -1
			return false
		}
	}
	//fmt.Printf("last check: pos:%d, j:%d\n", pos, j)
	if pos < 0 && j < 0 {
		dp[len(s)][len(p)] = 1
		return true
	}
	dp[len(s)][len(p)] = -1
	return false
}

func isMatch(s string, p string) bool {
	/*	flag := false
		p := ""
		for i := range p2 {
			if flag == false {
				p += string(p2[i])
				if p2[i] == '*' {
					flag = true
				}
			} else {
				flag = false
				if p2[i] != '*' {
					p += string(p2[i])
				}
			}
		}
	*/
	dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(p)+1)
	}
	//fmt.Printf("s:%s, p2:%s, p:%s\n", s, p2, p)

	return subIsMatch(s, p)
}
