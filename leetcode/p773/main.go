package main

import "fmt"

func senaryToDecimal(num [6]int) int {
	decimal := 0
	for _, i := range num {
		decimal = decimal*6 + i
	}
	return decimal
}

func slidingPuzzle(board [][]int) int {
	steps := make(map[int]int, 0)
	org := senaryToDecimal([6]int{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]})
	target := senaryToDecimal([6]int{1, 2, 3, 4, 5, 0})
	if org == target {
		return 0
	}
	steps[org] = 1

	fifo := make([]int, 721)

	phead, ptail := -1, 0
	fifo[0] = org

	for ; phead < ptail; {
		phead++
		tos := availableState(fifo[phead])
		for _, v := range tos {
			if v == target {
				return steps[fifo[phead]]
			}
			if _, exist := steps[v]; !exist {
				steps[v] = steps[fifo[phead]] + 1
				ptail++
				fifo[ptail] = v
			}
		}
	}
	return -1
}

var toSwap = [][]int{
	{1, 3},
	{0, 2, 4},
	{1, 5},
	{0, 4},
	{1, 3, 5},
	{2, 4},
}

var n6power = []int{1, 6, 36, 216, 1296, 7776, 46656, 279936}

func availableState(org int) []int {
	state := make([]int, 0, 2)
	find0 := org
	for i := 5; i >= 0; i-- {
		if find0%6 == 0 {
			for _, j := range toSwap[i] {
				state = append(state, swap(org, i, j))
			}
			break
		} else {
			find0 = find0 / 6
		}
	}
	return state
}

// value in position a is always 0
func swap(from int, a, b int) int {
	valueInB := from % n6power[6-b] / n6power[6-b-1]
	return from - valueInB*n6power[6-b-1] + valueInB*n6power[6-a-1]
}

func printSenary(a int) {
	for i := 5; i >= 0; i-- {
		fmt.Print(a / n6power[i])
		a = a % n6power[i]
	}
	fmt.Println()
}

func main() {
	input := [][]int{{4, 1, 2}, {5, 0, 3}}
	//input := [][]int{{1, 2, 3}, {4, 0, 5}}
	fmt.Println(slidingPuzzle(input))
}
