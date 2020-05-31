package main

import "fmt"

// Time Complexity O(N)
// Space Complexity O(N)
func longestValidParenthesesStackMethod(s string) int {
	stack := make([]int, len(s), len(s))
	p := -1
	endCharLength := make([]int, len(s), len(s))
	max := 0
	for index, c := range s {
		if c == '(' {
			p++
			stack[p] = index
		} else { // c==')'
			if p < 0 {
				continue
			}
			lastLeftIndex := stack[p]
			endCharLength[index] = index - stack[p] + 1
			if lastLeftIndex > 0 {
				endCharLength[index] += endCharLength[lastLeftIndex-1]
			}
			if endCharLength[index] > max {
				max = endCharLength[index]
			}
			p--
		}
	}
	return max
}

func longestValidParenthesesO1(s string) int {
	max := 0
	left, right := 0, 0
	for _, c := range s {
		if c == '(' {
			left++
		} else { // c == ')'
			right++
			if right > left {
				left, right = 0, 0
			} else if right == left {
				if right+left > max {
					max = right + left
				}
			}
		}
	}

	left, right = 0, 0
	for index := len(s) - 1; index >= 0; index-- {
		if s[index] == ')' {
			right++
		} else {
			left++
			if left > right {
				left, right = 0, 0
			} else if left == right {
				if left+right > max {
					max = left + right
				}
			}
		}
	}
	return max
}

func longestValidParentheses(s string) int {
	//return longestValidParenthesesStackMethod(s)
	return longestValidParenthesesO1(s)
}

func main() {
	s := "(())()"
	fmt.Println(longestValidParentheses(s))
}
