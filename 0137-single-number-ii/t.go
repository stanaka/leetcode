package main

func singleNumber(nums []int) int {
	ans := int32(0)
	mask := 1
	for i := uint32(0); i < 32; i++ {
		v := int32(0)
		for _, n := range nums {
			v += int32(n&mask) >> i
		}
		v %= 3
		v = v << i
		ans |= v
		//fmt.Printf("i: %d, v: %d\n", i, v)
		mask = mask << 1
	}
	return int(ans)
}
