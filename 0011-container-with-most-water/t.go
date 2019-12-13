package main

//import "fmt"

func maxArea(height []int) int {
	ans := 0
	l := 0
	r := len(height) -1
	for l != r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		area := h * (r-l)
		if ans < area {
			ans = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

type line struct {
	x int
	y int
}

func maxArea_bruteforce(height []int) int {
	ans := 0
	l := []line{}
	for i, h := range height {
		if len(l) == 0 {
			l = append(l, line{i, h})
			continue
		}
		flag := true
		for _, h2 := range l {
			w := h
			if h2.y < w {
				w = h2.y
			}
			if h2.y >= h {
				flag = false
			}
			a := (i-h2.x) * w
			if a > ans {
				ans = a
			}
		}
		if flag {
			l = append(l, line{i, h})
			/*
			// xn > (x2y2-x1y1)/(y2-y1)
			l2 := []line{}
			if len(l) > 1 {
				for j := 0; j < len(l)-1; j++ {
					r := (l[j+1].y*l[j+1].x - l[j].y*l[j].x) / (l[j+1].y - l[j].y)
					//fmt.Printf("h1{%d,%d}, h2{%d,%d}, r: %v\n", l[j].x,l[j].y,l[j+1].x,l[j+1].y,r)
					if r >= i {
						l2 = append(l2, line{l[j].x, l[j].y})
					}
				}
			}
			l = append(l2, line{i, h})
			*/
		}
		//fmt.Printf("%v\n", l)

	}
	//fmt.Printf("%v\n", ans)
	return ans
}