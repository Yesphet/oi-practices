package main

import (
	"fmt"
	"math"
)

var bitCountTable = []uint8{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 5, 6, 6, 7, 6, 7, 7, 8, 6, 7, 7, 8, 7, 8, 8, 9}

type Node struct {
	x, y int
}

func solveSudoku(board [][]byte) {
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

	availMap := make([][]uint16, 9, 9)
	for i := 0; i < 9; i++ {
		availMap[i] = make([]uint16, 9, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				continue
			}
			availMap[i][j] = calAvailable(board, i, j)
		}
	}

	trySolveSudoku(board, availMap)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] += 48
		}
	}
}

func trySolveSudoku(board [][]byte, availMap [][]uint16) bool {
	var oneAvailList = make([]Node, 1000, 1000)
	ph := 0
	pt := -1

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				continue
			}
			if availMap[i][j] == 0 {
				return false
			}
			if bitCountTable[availMap[i][j]] == 1 {
				pt++
				oneAvailList[pt] = Node{x: i, y: j}
			}
		}
	}

	for ; ph <= pt; {
		x := oneAvailList[ph].x
		y := oneAvailList[ph].y
		if bitCountTable[availMap[x][y]] == 0 {
			ph++
			continue
		}
		v := byte(math.Log2(float64(availMap[x][y])) + 1)
		board[x][y] = v
		availMap[x][y] = 0
		//fmt.Printf("%d : %d = %d\n", x+1, y+1, v)

		for i := 0; i < 9; i++ {
			availMap[x][i] = availMap[x][i] &^ (1 << (v - 1))
			availMap[i][y] = availMap[i][y] &^ (1 << (v - 1))

			if bitCountTable[availMap[x][i]] == 1 {
				pt++
				oneAvailList[pt] = Node{x: x, y: i}
			}
			if bitCountTable[availMap[i][y]] == 1 {
				pt++
				oneAvailList[pt] = Node{x: i, y: y}
			}
		}
		for i := x / 3 * 3; i < (x/3+1)*3; i++ {
			for j := y / 3 * 3; j < (y/3+1)*3; j++ {
				availMap[i][j] = availMap[i][j] &^ (1 << (v - 1))
				if bitCountTable[availMap[i][j]] == 1 {
					pt++
					oneAvailList[pt] = Node{x: i, y: j}
				}
			}
		}
		ph++
	}

	x, y := -1, -1
	availNum := uint8(10)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				continue
			}
			if bitCountTable[availMap[i][j]] < availNum {
				availNum = bitCountTable[availMap[i][j]]
				x = i
				y = j
			}
		}
	}
	if x == -1 && y == -1 {
		return true
	}

	for i := 0; i < 9; i++ {
		if availMap[x][y]&(1<<i) == (1 << i) {
			newBoard, newAvailMap := dupBoardAndAvailMap(board, availMap)
			newBoard[x][y] = byte(i) + 1
			newAvailMap[x][y] = 0
			updateAvailable(newAvailMap, x, y, newBoard[x][y])

			complete := trySolveSudoku(newBoard, newAvailMap)
			if complete {
				rewriteOldBoard(board, newBoard)
				return true
			}
		}
	}
	return false
}

func rewriteOldBoard(oldBoard [][]byte, newBoard [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			oldBoard[i][j] = newBoard[i][j]
		}
	}
}

func dupBoardAndAvailMap(board [][]byte, availMap [][]uint16) ([][]byte, [][]uint16) {
	newBoard := make([][]byte, 9, 9)
	newAvailMap := make([][]uint16, 9, 9)
	for i := 0; i < 9; i++ {
		newBoard[i] = make([]byte, 9, 9)
		copy(newBoard[i], board[i])
		newAvailMap[i] = make([]uint16, 9, 9)
		copy(newAvailMap[i], availMap[i])
	}
	return newBoard, newAvailMap
}

func updateAvailable(availMap [][]uint16, x, y int, v byte) {
	for i := 0; i < 9; i++ {
		availMap[x][i] = availMap[x][i] &^ (1 << (v - 1))
		availMap[i][y] = availMap[i][y] &^ (1 << (v - 1))
	}
	for i := x / 3 * 3; i < (x/3+1)*3; i++ {
		for j := y / 3 * 3; j < (y/3+1)*3; j++ {
			availMap[i][j] = availMap[i][j] &^ (1 << (v - 1))
		}
	}
}

func calAvailable(board [][]byte, x, y int) uint16 {
	availBinary := uint16(1<<9 - 1)
	for i := 0; i < 9; i++ {
		if board[x][i] != 0 {
			availBinary = availBinary &^ (1 << (board[x][i] - 1))
		}
		if board[i][y] != 0 {
			availBinary = availBinary &^ (1 << (board[i][y] - 1))
		}
	}
	for i := x / 3 * 3; i < (x/3+1)*3; i++ {
		for j := y / 3 * 3; j < (y/3+1)*3; j++ {
			if board[i][j] != 0 {
				availBinary = availBinary &^ (1 << (board[i][j] - 1))
			}
		}
	}
	return availBinary
}

func main() {
	//input := [][]byte{
	//	{'5','3','.','.','7','.','.','.','.'},
	//	{'6','.','.','1','9','5','.','.','.'},
	//	{'.','9','8','.','.','.','.','6','.'},
	//	{'8','.','.','.','6','.','.','.','3'},
	//	{'4','.','.','8','.','3','.','.','1'},
	//	{'7','.','.','.','2','.','.','.','6'},
	//	{'.','6','.','.','.','.','2','8','.'},
	//	{'.','.','.','4','1','9','.','.','5'},
	//	{'.','.','.','.','8','.','.','7','9'},
	//}

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
	solveSudoku(input)
	printBoard(input)
}

func printBoard(output [][]byte) {
	for i := range output {
		fmt.Printf("%v\n", output[i])
	}
	fmt.Println()
}

func printAvailMap(output [][]uint16) {
	for i := range output {
		fmt.Printf("%v\n", output[i])
	}
	fmt.Println()
}
