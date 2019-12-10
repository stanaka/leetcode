package main

// l1,m1,h1
// l2,m2,h2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := 0, 0
	h1 := len(nums1)
	h2 := len(nums2)
	//t1, t2 := 0, 0
	m1, m2 := h1/2, h2/2
	total := h1 + h2
	cm1, cm2 := total/2, 0
	if total%2 == 0 {
		cm1 -= 1
		cm2 = cm1 + 1
	}
	ans := 0

	//fmt.Printf("1:%v 2:%v median %d, %d\n", nums1, nums2, cm1, cm2)
	//fmt.Printf("1:%v [%d,%d,%d] 2:%v [%d,%d,%d]\n", nums1, l1, m1, h1, nums2, l2, m2, h2)
	if h2 == 0 || h1 > 0 && nums1[h1-1] < nums2[l2] {
		if cm1 < h1 {
			ans = nums1[cm1]
			if cm2 > 0 {
				if cm2 < h1 {
					ans += nums1[cm2]
				} else {
					ans += nums2[l2]
				}
			} else {
				ans += ans
			}
		} else {
			ans = nums2[cm1-h1]
			if cm2 > 0 {
				ans += nums2[cm2-h1]
			} else {
				ans += ans
			}
		}
		//fmt.Printf("return %v\n", float64(ans)/2)
		return float64(ans) / 2
	}
	if h1 == 0 || h2 > 0 && nums2[h2-1] < nums1[l1] {
		if cm1 < h2 {
			ans = nums2[cm1]
			if cm2 > 0 {
				if cm2 < h2 {
					ans += nums2[cm2]
				} else {
					ans += nums1[l1]
				}
			} else {
				ans += ans
			}
		} else {
			ans = nums1[cm1-h2]
			if cm2 > 0 {
				ans += nums1[cm2-h2]
			} else {
				ans += ans
			}
		}
		//fmt.Printf("return %v\n", float64(ans)/2)
		return float64(ans) / 2
	}
	//count := 0
	for l1+1 != h1 || l2+1 != h2 {
		//fmt.Printf("cond1: %d v %d\n", nums1[m1], nums2[m2])
		if nums1[m1] < nums2[m2] {
			// [m1+l2, m1+m2)
			if cm1 < m1+m2 {
				//fmt.Printf("cond1.1\n")
				if h2 != l2+1 {
					h2 = m2
				} else {
					h1 = m1
				}
			} else if cm1 < h1+m2 {
				//fmt.Printf("cond1.2\n")
				if l1 != m1 {
					l1 = m1
				} else {
					h2 = m2
				}
			} else {
				//fmt.Printf("cond1.3\n")
				if l2 != m2 {
					l2 = m2
				} else {
					l1 = m1
				}
			}
		} else if nums2[m2] < nums1[m1] {
			// [m1+m2, m1+h2)
			if cm1 < m1+m2 {
				//fmt.Printf("cond2.1\n")
				if h1 != l1+1 {
					h1 = m1
				} else {
					h2 = m2
				}
			} else if cm1 < m1+h2 {
				//fmt.Printf("cond2.2\n")
				if l2 != m2 {
					l2 = m2
				} else {
					h1 = m1
				}
			} else {
				//fmt.Printf("cond2.3\n")
				if l1 != m1 {
					l1 = m1
				} else {
					l2 = m2
				}
			}
		} else {
			// [m1+m2, m1+h2)
			if cm1 < m1+m2 {
				//fmt.Printf("cond3.1\n")
				if h1 != l1+1 {
					h1 = m1
				} else {
					h2 = m2
				}
			} else {
				//fmt.Printf("cond3.2\n")
				if l1 != m1 {
					l1 = m1
				} else {
					l2 = m2
				}
			}
		}
		m1 = (l1 + h1) / 2
		m2 = (l2 + h2) / 2
		//fmt.Printf("1:%v [%d,%d,%d] 2:%v [%d,%d,%d]\n", nums1, l1, m1, h1, nums2, l2, m2, h2)
		//count++
		//if count > 5 {
		//	break
		//}
	}

	if m1+m2+1 == cm1 {
		if nums1[m1] < nums2[m2] {
			ans = nums2[m2]
		} else {
			ans = nums1[m1]
		}
		if cm2 > 0 {
			if m2+1 >= len(nums2) || m1+1 < len(nums1) && nums1[m1+1] < nums2[m2+1] {
				ans += nums1[m1+1]
			} else {
				ans += nums2[m2+1]
			}
		} else {
			ans += ans
		}
	} else {
		if nums1[m1] < nums2[m2] {
			ans = nums1[m1]
			m1++
		} else {
			ans = nums2[m2]
			m2++
		}
		if cm2 > 0 {
			if m2 == len(nums1) || nums1[m1] < nums2[m2] {
				ans += nums1[m1]
			} else {
				ans += nums2[m2]
			}
		} else {
			ans += ans
		}
	}
	//fmt.Printf("return %v\n", float64(ans)/2)

	return float64(ans) / 2
}
