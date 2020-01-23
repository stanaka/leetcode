package main

func longestMountain(a []int) int {

	if len(a) < 3 {
		return 0
	}
	//fmt.Printf("%v\n", a)
	pos := -1
	flag := false
	ans := 0
	for i := 1; i < len(a); i++ {
		if pos == -1 {
			if a[i-1] < a[i] {
				pos = i - 1
				//fmt.Printf("%d: mountain start\n", i-1)
			}
		} else {
			if flag == false {
				if a[i-1] == a[i] {
					pos = -1
				} else if a[i-1] > a[i] {
					//fmt.Printf("%d: mountain peak\n", i)
					flag = true
				}
			} else {
				if a[i-1] <= a[i] {
					flag = false
					//fmt.Printf("%d: mountain end len: %d\n", i, i-pos)
					if ans < i-pos {
						ans = i - pos
					}
					pos = -1
					if a[i-1] < a[i] {
						pos = i - 1
						//fmt.Printf("%d: another mountain start\n", i-1)
					}
				}
			}
		}
	}

	if flag == true {
		if pos >= 0 && a[len(a)-2] > a[len(a)-1] {
			//fmt.Printf("%d: mountain end len:%d\n", len(a), len(a)-pos)
			if ans < len(a)-pos {
				ans = len(a) - pos
			}
		}
	}
	return ans
}
