package main

//import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			//match := i < len(s) && (p[j] == s[i] || string(p[j]) == ".")
			match := false
			if i < len(s) && (p[j] == s[i] || string(p[j]) == ".") {
				match = true
			}
			if j+1 < len(p) && string(p[j+1]) == "*" {
				dp[i][j] = dp[i][j+2] || match && dp[i+1][j]
			} else {
				dp[i][j] = match && dp[i+1][j+1]
			}
		}
	}
	//fmt.Printf("%v\n", dp)
	return dp[0][0]
}

/*
func isMatch_botsu(s string, p string) bool {
	sPos := 0
	pPos := 0
	bPos := -1
	fmt.Printf("s: %s, p %s\n", s,p)

	for sPos < len(s) && pPos < len(p) {
		fmt.Printf("s[%d]: %s, p[%d]: %s\n", sPos, string(s[sPos]), pPos, string(p[pPos]))
		if string(p[pPos]) == "." {
			// case .*:
			if pPos +1 < len(p) && string(p[pPos+1]) == "*" {
				bPos = pPos
				pPos++
				if pPos + 1 < len(p) {
					for sPos < len(s) {
						fmt.Printf("s[%d]: %s, p[%d]: %s\n", sPos, string(s[sPos]), pPos, string(p[pPos]))
						if string(p[pPos+1]) != string(s[sPos]) {
							sPos++
						} else {
							pPos++
							break
						}
					}
				} else {
					return true
				}
			} else {
				pPos++
				sPos++
			}
		} else {
			if string(p[pPos]) == string(s[sPos]) {
				count := 0
				if pPos +1 < len(p) && string(p[pPos+1]) == "*" {
					fmt.Printf("star check after %s\n", string(p[pPos]))
					idx := 2
					for pPos + idx < len(p) {
						if string(p[pPos+idx]) == string(p[pPos]) {
							count++
							idx++
						} else if string(p[pPos+idx]) == "*" {
							count--
							idx++
						} else {
							break
						}
					}
					fmt.Printf("star check after %s count:%d\n", string(p[pPos]), count)
					for count > 0 && sPos < len(s) {
						if string(p[pPos]) == string(s[sPos]) {
							count--
							sPos++
						} else {
							fmt.Printf("s[%d/%d], p[%d/%d]: false\n", sPos, len(s), pPos, len(p))
							return false
						}
					}
					for sPos < len(s) && string(p[pPos]) == string(s[sPos]) {
						sPos++
					}
					fmt.Printf("p[%d+%d]\n", pPos, idx)
					pPos += idx
				} else {
					pPos++
					sPos++
				}
			} else {
				if pPos +1 < len(p) && string(p[pPos+1]) == "*" {
					pPos+= 2
				} else {
					if bPos >= 0 {
						fmt.Printf("backtrack!!\n")
						pPos = bPos
					} else {
						return false
					}
				}
			}
		}
	}
	if pPos == len(p) && sPos == len(s) {
		return true
	}
	if sPos == len(s) {
		for pPos + 1 < len(p) && string(p[pPos+1]) == "*" {
			pPos += 2
		}
		if pPos == len(p) {
			return true
		}
	}

	fmt.Printf("s[%d/%d], p[%d/%d]: false\n", sPos, len(s), pPos, len(p))
	return false
}
*/
