package main

import "fmt"

func sortIntervalStartPoint(intervals [][]int, indices []int, start, end int) {
	if start >= end {
		return
	}

	i, j := start, end
	pivot := intervals[indices[(start+end)/2]][0]
	for i != j {
		for i < j && intervals[indices[i]][0] < pivot {
			i++
		}
		for i < j && intervals[indices[j]][0] > pivot {
			j--
		}
		if i < j {
			tmp := indices[i]
			indices[i] = indices[j]
			indices[j] = tmp
		}
	}

	sortIntervalStartPoint(intervals, indices, start, i-1)
	sortIntervalStartPoint(intervals, indices, j+1, end)
}

func binaryChop(intervals [][]int, indices []int, v int) int {
	start, end := 0, len(intervals)-1
	ret := -1
	for start <= end {
		mid := (start + end) / 2
		if intervals[indices[mid]][0] == v {
			return indices[mid]
		}
		if intervals[indices[mid]][0] < v {
			start = mid + 1
		} else {
			end = mid - 1
			ret = indices[mid]
		}
	}
	return ret
}

func findRightInterval(intervals [][]int) []int {
	ret := make([]int, len(intervals))
	indices := make([]int, len(intervals))
	for i := 0; i < len(indices); i++ {
		indices[i] = i
	}
	sortIntervalStartPoint(intervals, indices, 0, len(intervals)-1)
	for i := 0; i < len(intervals); i++ {
		ret[i] = binaryChop(intervals, indices, intervals[i][1])
	}
	return ret
}

func main() {
	//input := [][]int{{1, 2}, {3, 4}, {5, 6}, {2, 7}}
	input := [][]int{{1,4}, {2,3},{3,4}}
	fmt.Printf("%v", findRightInterval(input))
}
