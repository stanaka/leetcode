package main

import "strings"

func convert(s string, numRows int) string {
	a := make([]string, numRows)
	if numRows == 1 {
		return s
	}
	steps := numRows*2 - 2
	for i := 0; i < len(s); i++ {
		idx := i % steps
		if idx >= numRows {
			idx = numRows*2 - idx - 2
		}
		a[idx] += string(s[i])
	}
	return strings.Join(a, "")
}
