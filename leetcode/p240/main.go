package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	mostBottom := len(matrix) - 1
	k, phead, ptail := 0, 0, 0
	for k < len(matrix[0]) {
		phead, ptail = 0, mostBottom
		for phead <= ptail {
			middle := (phead + ptail) / 2
			if matrix[middle][k] > target {
				ptail = middle - 1
			} else if matrix[middle][k] < target {
				phead = middle + 1
			} else if matrix[middle][k] == target {
				return true
			}
		}
		if ptail < 0 {
			return false
		}
		mostBottom = ptail
		k++
	}
	return false
}

func main() {
	input := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(input, 20))
}
