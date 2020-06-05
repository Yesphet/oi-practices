package main

import "fmt"

const (
	formFile = iota
	formDir
)

type node struct {
	length int
	level  int
	form   int
}

func lengthLongestPath(input string) int {
	stack := make([]*node, 0)
	p := -1
	inStackLength := 0
	ret := 0
	i := -1
	for ; ; {
		for ; i+1 < len(input) && input[i+1] == '\n'; {
			i += 1
		}
		level := 0
		for ; i+1 < len(input) && input[i+1] == '\t'; {
			i = i + 1
			level += 1
		}
		if i >= len(input)-1 {
			break
		}
		node := &node{
			length: 0,
			level:  level,
			form:   formDir,
		}
		for ; i < len(input)-1 && input[i+1] != '\n'; {
			i++
			node.length++
			if input[i] == '.' {
				node.form = formFile
			}
		}

		for ; p >= 0 && stack[p].level >= node.level; {
			inStackLength -= stack[p].length
			p--
		}
		p++
		if p >= len(stack) {
			stack = append(stack, node)
		} else {
			stack[p] = node
		}
		inStackLength += node.length
		if node.form == formFile && inStackLength+p > ret {
			ret = inStackLength+p
		}
	}
	return ret
}

func main() {
	//input := `dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext`
	input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println(lengthLongestPath(input))
}
