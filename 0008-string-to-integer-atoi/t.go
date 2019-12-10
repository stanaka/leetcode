package main

import "math"

func myAtoi(str string) int {
	IntMax := int64(math.Pow(2, 31)) - 1
	IntMin := int64(math.Pow(-2, 31))
	ans := int64(0)
	flag := false
	i := 0
	for i < len(str) && str[i] == " "[0] {
		i++
	}
	if i == len(str) {
		return int(ans)
	}
	if str[i] == "+"[0] {
		i++
	} else if str[i] == "-"[0] {
		i++
		flag = true
	}
	for ; i < len(str); i++ {
		if str[i] >= "0"[0] && str[i] <= "9"[0] {
			ans = ans*10 + int64(str[i]-"0"[0])
			if ans > IntMax {
				break
			}
			continue
		}
		break
	}
	if flag {
		ans = -ans
	}
	if ans > IntMax {
		ans = IntMax
	}
	if ans < IntMin {
		ans = IntMin
	}
	return int(ans)
}
