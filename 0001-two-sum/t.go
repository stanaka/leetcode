package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if _, ok := m[c]; ok {
			return []int{m[c], i}
		}
		m[nums[i]] = i
	}
	// should not be here
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				return []int{j, i}
			}
		}
	}
	// should not be here
	return []int{}
}
