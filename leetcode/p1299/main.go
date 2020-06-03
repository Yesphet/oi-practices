package main

import "fmt"

func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		k := arr[i]
		arr[i] = max
		if k > max {
			max = k
		}
	}
	return arr
}

func main() {
	input := []int{17, 18, 5, 4, 6, 1}
	fmt.Printf("%v", replaceElements(input))
}
