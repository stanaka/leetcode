package main

import (
	"fmt"
	"sort"
)

// 1 2 3 4 5  6  7  8   9
// 1 2 4 8 16 32 64 128 256
var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

func displayBoard2(board [][]byte) {
	if production {
		return
	}

	for i := range board {
		for j := range board[i] {
			fmt.Printf("%s", string(board[i][j]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

var retPow []uint16

func pow(x, y uint16) uint16 {
	if x == 2 && y < 10 {
		return retPow[y]
	}
	z := x
	if y == 0 {
		return 1
	}
	for i := uint16(1); i < y; i++ {
		z = z * x
	}
	return z
}

func updatePotential(board [][]byte, potential [][]uint16, nums [][]uint8, i int, j int, num byte) bool {
	//debugf("updatePotential i:%d, j:%d, num:%d\n", i, j, num)

	v := pow(2, uint16(num-1))
	if potential[i][j] == pow(2, 9)-1 {
		if board[i][j] == byte(num)+'0' {
			//debugf("updatePotential false!! i:%d, j:%d, num:%d\n", i, j, num)
			return false
		}
		//debugf("updatePotential: true\n")
		return true
	}
	if potential[i][j]&v > 0 {
		return true
	}

	//debugf("updatePotential i:%d, j:%d, num:%d\n", i, j, num)
	potential[i][j] |= v
	nums[i][j]--
	if nums[i][j] == 1 {
		for k := 0; k < 9; k++ {
			if potential[i][j]&pow(2, uint16(k)) == 0 {
				//debugf("call set i:%d, j:%d, num:%d\n", i, j, k+1)
				return set(board, potential, nums, i, j, byte(k+1))
			}
		}
		return false
	}
	return true
}

func set(board [][]byte, potential [][]uint16, nums [][]uint8, i int, j int, num byte) bool {
	//debugf("set i:%d, j:%d, num:%d\n", i, j, num)
	if potential[i][j] == pow(2, 9)-1 {
		return true
	}
	board[i][j] = byte(num) + '0'
	nums[i][j] = 1
	potential[i][j] = pow(2, 9) - 1 // - pow(2, uint16(num))
	for k := range potential {
		if j != k && updatePotential(board, potential, nums, i, k, num) == false {
			return false
		}
		if i != k && updatePotential(board, potential, nums, k, j, num) == false {
			return false
		}
	}
	iS := 3 * (i / 3)
	jS := 3 * (j / 3)
	//debugf("set updatePotential i:%d, j:%d, iS:%d, jS:%d\n", i, j, iS, jS)
	for x := iS; x <= iS+2; x++ {
		for y := jS; y <= jS+2; y++ {
			if !(x == i && y == j) && updatePotential(board, potential, nums, x, y, num) == false {
				return false
			}
		}
	}
	return true
}

func checkCompletion(potential [][]uint16, board [][]byte) bool {
	for i := range potential {
		for j := range potential[i] {
			if potential[i][j] != pow(2, 9)-1 {
				debugf("Not completed!!\n")
				return false
			} else if board[i][j] == '.' {
				debugf("Not completed!!\n")
				return false
			}
		}
	}
	//fmt.Printf("Success!!\n")
	return true
}

func printPotential(p [][]uint16) {
	if production {
		return
	}
	potential := make([][]uint16, 9)
	for i := range p {
		potential[i] = make([]uint16, 9)
		copy(potential[i], p[i])
	}
	for i := range potential {
		for j := range potential[i] {
			potential[i][j] ^= pow(2, 9) - 1
		}
	}
	fmt.Printf("%v\n", potential)
}

func doAssume(board [][]byte, potential [][]uint16, nums [][]uint8, i int, j int, k uint16) bool {
	debugf("doAssume: %d, %d: %d\n", i+1, j+1, k+1)

	ret := set(board, potential, nums, i, j, byte(k+1))
	if ret == false {
		return false
	}

	if checkCompletion(potential, board) {
		return true
	} else {
		res := assume(board, potential, nums)
		if res {
			return true
		}
	}
	return false
}

type bc struct {
	i, j  int
	value int
}

type bcs []bc

func (b bcs) Len() int           { return len(b) }
func (b bcs) Less(i, j int) bool { return b[i].value < b[j].value }
func (b bcs) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func assume(board [][]byte, potential [][]uint16, nums [][]uint8) bool {
	debugf("assume\n")
	displayBoard2(board)
	//printPotential(potential)

	potentialCells := bcs{}
	for i := range potential {
		for j := range potential[i] {
			count := 0
			for k := uint16(0); k < uint16(9); k++ {
				if potential[i][j]&pow(2, k) == 0 {
					count++
				}
			}
			if count > 0 {
				potentialCells = append(potentialCells, bc{i, j, count})
			}
		}
	}
	sort.Sort(potentialCells)
	debugf("%v\n", potentialCells)

	for x := range potentialCells {
		//fmt.Printf("%v\n", potentialCells[x])

		for k := uint16(0); k < uint16(9); k++ {

			if potential[potentialCells[x].i][potentialCells[x].j]&pow(2, k) == 0 {
				board2 := make([][]byte, 9)
				for i := range board {
					board2[i] = make([]byte, 9)
					copy(board2[i], board[i])
				}
				potential2 := make([][]uint16, 9)
				for i := range potential {
					potential2[i] = make([]uint16, 9)
					copy(potential2[i], potential[i])
				}
				nums2 := make([][]uint8, 9)
				for i := range potential {
					nums2[i] = make([]uint8, 9)
					copy(nums2[i], nums[i])
				}

				res := doAssume(board, potential, nums, potentialCells[x].i, potentialCells[x].j, k)
				if res {
					debugf("assume: true!!\n")
					displayBoard2(board)
					return true
				}
				debugf("assume: rollback!!\n")

				for i := range board {
					copy(board[i], board2[i])
				}
				for i := range potential {
					copy(potential[i], potential2[i])
				}
				for i := range nums {
					copy(nums[i], nums2[i])
				}

			}
		}
		break
	}
	return false

}

func solveSudoku(board [][]byte) {
	production = true
	retPow = []uint16{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	//fmt.Printf("%v\n", board)

	potential := make([][]uint16, 9)
	nums := make([][]uint8, 9)
	for i := range potential {
		potential[i] = make([]uint16, 9)
		nums[i] = make([]uint8, 9)
		for j := range nums[i] {
			nums[i][j] = 9
		}
	}

	displayBoard2(board)
	for i := range potential {
		for j := range potential[i] {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				if set(board, potential, nums, i, j, num) == false {
					debugf("should not reach here!!\n")
				}
			}
		}
	}
	displayBoard2(board)

	if assume(board, potential, nums) == false {
		debugf("Cannot solve!!\n")
	}
	//fmt.Printf("lb: %v\n", board)

	//printPotential(potential)

}
