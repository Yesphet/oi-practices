package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			switch board[i][j] {
			case '.':
				board[i][j] = 0
			default:
				board[i][j] = board[i][j] - 48
			}
		}
	}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == 0 {
				continue
			}
			for i := 0; i < 9; i++ {
				if i != y && board[x][i] == board[x][y] {
					return false
				}
				if i != x && board[i][y] == board[x][y] {
					return false
				}
			}
			for i := x / 3 * 3; i < (x/3+1)*3; i++ {
				for j := y / 3 * 3; j < (y/3+1)*3; j++ {
					if (i != x || j != y) && board[i][j] == board[x][y] {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {

	input := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	fmt.Printf("%v", isValidSudoku(input))
}
