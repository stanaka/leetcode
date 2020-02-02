package main

// 20 - 9 = 11
// 20 - 27 = -7

// 19 - 9 = 10
// 19 - 27 = -8

// 3*3*3*3*3*3 = 729
// 3*3*3*3*3 = 243
// 3*3*3*3 = 81
// 3*3*3 = 27
// 3*3 = 9

// 178 = 3^5 - 3^4 + 3^3 - 3^2 - 3 + 3/3
// 178 = 3^4 + 3^4

// 200 = 3^4 + 3^4 + 3^3 + 3 - 3/3
//        3  1  3  1  2  1   1  1 = 13
// 200 = 3^5 - 3^3 - 3^2 - 3 - 3 - 3/3
//        4  1  2  1  1  1   1   1  1 = 13

// 200 = 3^5 - 43
// 200 = 3^4 - 119

// 2 = 3/3 + 3/3
// 2 = 3-3/3

// x: 3, target: 929
//diff: 200, ops: 5, ans:5
//diff: 43, ops: 4, ans:10
//diff: 16, ops: 2, ans:13
//diff: 7, ops: 1, ans:15
//diff: 4, ops: 0, ans:16
//diff: 1, ops: 0, ans:17
//diff: 0, ops: 1, ans:19

// x: 3, target: 929
//diff: 200, ops: 5, ans:5
//diff: 119, ops: 3, ans:9
//diff: 38, ops: 3, ans:13
//diff: 11, ops: 2, ans:16
//diff: 2, ops: 1, ans:18
//diff: 1, ops: 0, ans:19
//diff: 0, ops: 1, ans:21

//x: 3, target: 907
//diff: 178, ops: 5, ans:5
//diff: 65, ops: 4, ans:10
//diff: 16, ops: 3, ans:14
//diff: 7, ops: 1, ans:16
//diff: 4, ops: 0, ans:17
//diff: 1, ops: 0, ans:18
//diff: 0, ops: 1, ans:20

//x: 3, target: 907
//diff: 178, ops: 5, ans:5
//diff: 97, ops: 3, ans:9
//diff: 16, ops: 3, ans:13
//diff: 7, ops: 1, ans:15
//diff: 4, ops: 0, ans:16
//diff: 1, ops: 0, ans:17
//diff: 0, ops: 1, ans:19

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findNext(x, diff int) int {
	//fmt.Printf("f: x: %d, diff: %d\n", x, diff)
	if x > diff {
		return min(diff*2-1, (x-diff)*2)
	}

	if x == diff {
		return 0
	}

	i := 0
	m := x
	for m < diff {
		i++
		m *= x
	}

	if m == diff {
		return i
	}

	l, r := MaxInt, MaxInt

	if m-diff < diff {
		l = findNext(x, m-diff) + i
	}
	r = findNext(x, diff-(m/x)) + i - 1

	//fmt.Printf("f: x: %d, diff: %d, ans: %d, l:%d, r:%d\n", x, diff, min(l, r)+1, l, r)
	return min(l, r) + 1
}

func leastOpsExpressTarget(x int, target int) int {
	//fmt.Printf("x: %d, target: %d\n", x, target)
	return findNext(x, target)
}
