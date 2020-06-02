package main

import "fmt"

func trap(height []int) int {
	waterHeight := make([]int, len(height), len(height))
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		waterHeight[i] = max
	}

	ret := 0
	max = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		if max < waterHeight[i] {
			waterHeight[i] = max
		}
		ret += waterHeight[i] - height[i]
	}
	return ret
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("%v\n", trap(input))
}
