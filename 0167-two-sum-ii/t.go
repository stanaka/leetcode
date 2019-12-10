package main

func twoSum(nums []int, target int) []int {
	i := len(nums)-1
	j := 0
	for {
		if nums[i]+nums[j] == target {
			return []int{j+1, i+1}
		} else if nums[i]+nums[j] > target {
			i--
		} else if nums[i]+nums[j] < target {
			j++
		}
	}
	// should not be here
	return []int{}
}
