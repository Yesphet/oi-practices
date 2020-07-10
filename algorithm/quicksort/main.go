package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort(list []int, index1, index2 int) {
	if index1 >= index2 {
		return
	}
	var tmp int
	i, j := index1, index2

	pivot := list[(index1+index2)/2]
	for i != j {
		for i < j && list[i] < pivot {
			i++
		}
		for i < j && list[j] > pivot {
			j--
		}
		if i < j {
			tmp = list[i]
			list[i] = list[j]
			list[j] = tmp
		}
	}
	sort(list, index1, i-1)
	sort(list, j+1, index2)
}

func sortWithDuplicateElement(list []int, index1, index2 int) {
	if index1 >= index2 {
		return
	}
	var tmp int
	i, j, k := index1, index2, index2

	pivot := list[(index1+index2)/2]
split:
	for i != j {
		for i < j && list[i] < pivot {
			i++
		}
		for i < j && list[j] >= pivot {
			if list[j] == pivot {
				for k = j; i < k && list[k] == pivot; {
					k--
				}
				tmp = list[k]
				list[k] = list[j]
				list[j] = tmp
				if list[j] > pivot {
					j--
				}
				if k == i {
					break split
				}
			} else {
				j--
			}
		}
		if i < j {
			tmp = list[i]
			list[i] = list[j]
			list[j] = tmp
		}
	}
	sortWithDuplicateElement(list, index1, i-1)
	sortWithDuplicateElement(list, j+1, index2)
}

func validate(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] < list[i-1] {
			return false
		}
	}
	return true
}

func main() {
	list1 := rand.Perm(5000000)
	//for i := 0; i < len(list); i++ {
	//	list[i] = list[i] / 5
	//}
	list2 := make([]int, len(list1))
	copy(list2, list1)

	start := time.Now()
	sortWithDuplicateElement(list1, 0, len(list1)-1)
	fmt.Printf("sortWithDuplicateElement:\t %dns valid %t\n", time.Now().Sub(start).Nanoseconds(), validate(list1))

	start = time.Now()
	sort(list2, 0, len(list2)-1)
	fmt.Printf("sort:\t\t\t\t %dns valid %t\n", time.Now().Sub(start).Nanoseconds(), validate(list2))

}
