package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	flag := map[int]map[int]map[int]map[int]bool{}
	if len(nums) < 4 {
		//fmt.Printf("under 4\n")
		return ans
	}
	sort.Ints(nums)
	for i1 := 0; i1 < len(nums)-3; i1++ {
		for i2 := i1 + 1; i2 < len(nums)-2; i2++ {
			i3 := i2 + 1
			i4 := len(nums) - 1
			for i3 < i4 {
				//fmt.Printf("1:[%d] %d, 2:[%d] %d, 3:[%d] %d, 4:[%d] %d\n", i1, nums[i1], i2, nums[i2], i3, nums[i3], i4, nums[i4])
				if nums[i1]+nums[i2]+nums[i3]+nums[i4] == target {
					if !flag[nums[i1]][nums[i2]][nums[i3]][nums[i4]] {
						ans = append(ans, []int{nums[i1], nums[i2], nums[i3], nums[i4]})
						if flag[nums[i1]] == nil {
							flag[nums[i1]] = map[int]map[int]map[int]bool{}
						}
						if flag[nums[i1]][nums[i2]] == nil {
							flag[nums[i1]][nums[i2]] = map[int]map[int]bool{}
						}
						if flag[nums[i1]][nums[i2]][nums[i3]] == nil {
							flag[nums[i1]][nums[i2]][nums[i3]] = map[int]bool{}
						}
						flag[nums[i1]][nums[i2]][nums[i3]][nums[i4]] = true
					}
					//fmt.Printf("append!!\n")
					i4--
				} else if nums[i1]+nums[i2]+nums[i3]+nums[i4] < target {
					i3++
				} else {
					i4--
				}
			}
		}
	}
	return ans
}
