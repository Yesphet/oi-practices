package main

import (
	"fmt"
	"sort"
)

type Courses [][]int

func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	if c[i][1] < c[j][1] || c[i][1] == c[j][1] && c[i][0] < c[j][0] {
		return true
	}
	return false
}

func (c Courses) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func scheduleCourse(courses [][]int) int {
	c := Courses(courses)
	sort.Sort(c)

	k := c[len(c)-1][1]
	f := make([]int, k+1)

	for i := 0; i < len(c); i++ {
		for j := c[i][1]; j >= c[i][0]; j-- {
			f[j] = max(f[j-c[i][0]]+1, f[j])
		}
		if i != len(c)-1 {
			for j := c[i][1] + 1; j <= c[i+1][1]; j++ {
				f[j] = max(f[j-1], f[j])
			}
		}
	}

	return f[k]
}

func main() {
	//input := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	input := [][]int{{7, 16}, {2, 3}, {3, 12}, {3, 14}, {10, 19}, {10, 16}, {6, 8}, {6, 11}, {3, 13}, {6, 16}}
	fmt.Println(scheduleCourse(input))
}
