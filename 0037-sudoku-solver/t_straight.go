package main

import "fmt"

func displayBoard(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%s", string(board[i][j]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func isValid(board [][]byte, x, y int, k byte) bool {
	for i := range board {
		if board[i][y] != '.' && board[i][y] == k {
			return false
		}
		if board[x][i] != '.' && board[x][i] == k {
			return false
		}
		cx := 3*(x/3) + i%3
		cy := 3*(y/3) + i/3
		if board[cx][cy] != '.' && board[cx][cy] == k {
			return false
		}
	}
	return true
}

func solve(board [][]byte) bool {
	//displayBoard(board)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				for k := byte('1'); k <= '9'; k++ {
					if isValid(board, i, j, k) {
						board[i][j] = k
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func solveSudoku2(board [][]byte) {
	solve(board)
}
