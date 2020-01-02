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

func subTrap(h []int, l, r int) int {
	//fmt.Printf("l:%d, r:%d\n", l, r)
	if r-l == 1 {
		return 0
	}
	max := 0
	pos := -1
	for i := l + 1; i < r; i++ {
		if h[i] > max {
			pos = i
			max = h[i]
		}
	}
	vol := 0
	height := min(h[l], h[r])

	if pos == -1 {
		return height * (r - l - 1)
	}

	if h[pos] < height {
		vol += (height - h[pos]) * (r - l - 1)
	}
	if l != pos {
		vol += subTrap(h, l, pos)
	}
	if r != pos {
		vol += subTrap(h, pos, r)
	}
	return vol
}

func trap(h []int) int {
	//fmt.Printf("h:%v\n", h)

	if len(h) < 2 {
		return 0
	}

	return subTrap(h, 0, len(h)-1)
}
