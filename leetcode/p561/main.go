package main

import "fmt"

func arrayPairSum(nums []int) int {
	base := make([]int, 20001, 20001)
	for _, v := range nums {
		base[v+10000]++
	}

	sum := 0
	isNewPair := true
	for i := range base {
		if base[i] == 0 {
			continue
		}
		if isNewPair {
			if base[i]%2 == 0 {
				sum += (i - 10000) * (base[i] / 2)
			} else {
				sum += (i - 10000) * (base[i]/2 + 1)
				isNewPair = false
			}
		} else {
			if base[i]%2 == 0 {
				sum += (i - 10000) * (base[i] / 2)
			} else {
				sum += (i - 10000) * (base[i] / 2)
				isNewPair = true
			}
		}
	}

	return sum
}

func main() {
	input := []int{1, 2, 3, 2}
	fmt.Println(arrayPairSum(input))
}
