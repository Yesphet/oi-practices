package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	sort.Ints(nums)
	var l int
	ret := make([][]int, 0)
	for ; ; {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		ret = append(ret, newNums)

		l = -1
		for i := len(nums) - 2; i >= 0; i-- {
			if nums[i] < nums[i+1] {
				l = i
				break
			}
		}
		if l == -1 {
			return ret
		}
		for i := len(nums) - 1; i > l ; i-- {
			if nums[i] > nums[l] {
				tmp := nums[l]
				nums[l] = nums[i]
				nums[i] = tmp
				break
			}
		}
		i := l + 1
		j := len(nums) - 1
		for ; i < j; {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
			j--
		}
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Printf("%v", permute(input))
}
