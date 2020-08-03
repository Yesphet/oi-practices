package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	sum := make([]int, n)
	sum[0] = stones[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + stones[i]
	}

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	for l := k - 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			f[i][i+l] = math.MaxInt32
			for mid := i; mid < i+l; mid += k - 1 {
				f[i][i+l] = min(f[i][i+l], f[i][mid]+f[mid+1][i+l])
			}
			if l%(k-1) == 0 {
				f[i][i+l] += sum[i+l] - sum[i] + stones[i]
			}
		}
	}
	//fmt.Printf("%v\n", f)
	if (n-1)%(k-1) == 0 {
		return f[0][n-1]
	} else {
		return -1
	}
}

func main() {
	input := []int{3, 5, 1, 2, 6}
	k := 3
	//input := []int{3, 2, 4, 1}
	//k := 2
	fmt.Println(mergeStones(input, k))
}
