package main

// O(N)
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	preSum := make(map[int]int)
	preSum[0] = 1
	sum := 0
	for i := range nums {
		sum += nums[i]
		ans += preSum[sum-k]
		preSum[sum] += 1
	}
	return ans
}

// O(N^2)
func subarraySum_simple(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
