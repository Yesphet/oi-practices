package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	phead, ptail := 0, 0
	count := 0
	current := 1

	for ; ptail < len(nums); {
		current *= nums[ptail]
		for ; current >= k && phead <= ptail; {
			current = current / nums[phead]
			phead++
		}
		if ptail >= phead {
			count += ptail - phead + 1
		}
		ptail++
	}
	return count
}

func main() {
	input := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(input, k))
}
