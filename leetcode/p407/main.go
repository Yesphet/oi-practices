package main

import (
	"fmt"
)

func trapRainWater(heightMap [][]int) int {
	return trapRainWaterCustomHeap(heightMap)
}

func main() {
	//input := [][]int{
	//	{1, 4, 3, 1, 3, 2},
	//	{3, 2, 1, 3, 2, 4},
	//	{2, 3, 3, 2, 3, 1},
	//}
	input := [][]int{
		{5, 5, 5, 1},
		{5, 1, 1, 5},
		{5, 1, 5, 5},
		{5, 2, 5, 8},
	}
	fmt.Println(trapRainWater(input))
}
