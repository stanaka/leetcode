package main

func firstMissingPositive(nums []int) int {
	//fmt.Printf("%v\n", nums)
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			//fmt.Printf("%d: %v\n", i, nums)
		}
	}
	//fmt.Printf("%v\n", nums)
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
