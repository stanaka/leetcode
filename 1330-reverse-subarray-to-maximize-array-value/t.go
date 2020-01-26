package main

// 2,5,3,6,4
//  3,2,3,2 = 10

// [3,5,2],6,4
//  2,3,4,2 = 11

// 2,5,1,3,4
//  3,4,2,1 = 10

// 2,5,1,[4,3]
//  3,4,3,1 = 11

// 2,3,1,5,4
//  1,2,4,1 = 8

// 2,[3,1,5],4
// 2,5,1,3,4
//  3,4,2,1 = 10

// 2,3,5,1,4
//  1,2,4,3

// 2,4,9,24,2,1,10
//  2,5,15,22,1,9 = 52

// 2,4,9,2,24,1,10
//  2,5,7,22,23,9 = 68

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

func maxValueAfterReverse(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	//fmt.Printf("len:%d\n", len(nums))
	initial := 0
	for i := 1; i < len(nums); i++ {
		initial += abs(nums[i] - nums[i-1])
	}
	ans := initial
	diff := 0
	for l := 0; l < len(nums)-1; l++ {
		diff = max(diff,
			abs(nums[0]-nums[l+1])-abs(nums[l]-nums[l+1]),
			abs(nums[l]-nums[len(nums)-1])-abs(nums[l]-nums[l+1]),
		)
	}
	high := 0
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	low := MaxInt
	for l := 0; l < len(nums)-1; l++ {
		high = max(high, min(nums[l], nums[l+1]))
		low = min(low, max(nums[l], nums[l+1]))
	}
	ans += max(diff, 2*(high-low))
	return ans
}

func maxValueAfterReverseStraight(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	initial := 0
	for i := 1; i < len(nums); i++ {
		initial += abs(nums[i] - nums[i-1])
	}
	ans := initial
	for l := 0; l < len(nums)-1; l++ {
		for r := l + 1; r < len(nums); r++ {
			cur := initial
			if 0 < l {
				cur -= abs(nums[l] - nums[l-1])
				cur += abs(nums[r] - nums[l-1])
			}
			if r < len(nums)-1 {
				cur -= abs(nums[r+1] - nums[r])
				cur += abs(nums[r+1] - nums[l])
			}
			ans = max(ans, cur)
		}
	}

	return ans
}
