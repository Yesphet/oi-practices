package main

import "fmt"

func moveZeroes(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums) && nums[j] != 0; j++ {
				tmp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = tmp
			}
		}
	}
}

func main() {
	input := []int{0, 1, 0, 3, 12}
	moveZeroes(input)
	fmt.Printf("%v", input)
}
