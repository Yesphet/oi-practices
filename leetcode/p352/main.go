package main

import "fmt"

type Node struct {
	start, end          int
	left, right, parent *Node
}

func NewNode(v int, n *Node) *Node {
	return &Node{
		start:  v,
		end:    v,
		left:   nil,
		right:  nil,
		parent: n,
	}
}

func (n *Node) leftBiggest() *Node {
	if n.left == nil {
		return nil
	}
	k := n.left
	for k.right != nil {
		k = k.right
	}
	return k
}

func (n *Node) rightSmallest() *Node {
	if n.right == nil {
		return nil
	}
	k := n.right
	for k.left != nil {
		k = k.left
	}
	return k
}

func (n *Node) selfRemove() {
	if n.parent == nil {
		return
	}

	child := n.left
	if child == nil {
		child = n.right
	}
	if n.parent.left == n {
		n.parent.left = child
	} else {
		n.parent.right = child
	}
	if child != nil {
		child.parent = n.parent
	}
}

// v must greater or equal than n.start and less or equal than n.end
func (n *Node) insert(v int) {
	if v >= n.start && v <= n.end {
		return
	}

	if v == n.start-1 {
		n.start = v
		k := n.leftBiggest()
		if k != nil {
			if k.end == n.start-1 {
				n.start = k.start
				k.selfRemove()
			}
		}
		return
	}
	if v == n.end+1 {
		n.end = v
		k := n.rightSmallest()
		if k != nil {
			if k.start == n.end+1 {
				n.end = k.end
				k.selfRemove()
			}
		}
		return
	}

	if v < n.start {
		if n.left != nil {
			n.left.insert(v)
		} else {
			n.left = NewNode(v, n)
		}
	} else {
		if n.right != nil {
			n.right.insert(v)
		} else {
			n.right = NewNode(v, n)
		}
	}
}

type SummaryRanges struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{
		nil,
	}
}

func (this *SummaryRanges) AddNum(val int) {
	if this.root == nil {
		this.root = NewNode(val, nil)
		return
	}

	this.root.insert(val)
}

func (this *SummaryRanges) GetIntervals() [][]int {
	if this.root == nil {
		return nil
	}
	stack := []*Node{this.root}
	p := 0
	recall := false
	ret := make([][]int, 0)
	for p >= 0 {
		if !recall && stack[p].left != nil {
			if p == len(stack)-1 {
				stack = append(stack, stack[p].left)
			} else {
				stack[p+1] = stack[p].left
			}
			p++
			continue
		}
		ret = append(ret, []int{stack[p].start, stack[p].end})
		if stack[p].right != nil {
			recall = false
			stack[p] = stack[p].right
		} else {
			p--
			recall = true
		}
	}
	return ret
}

func main() {
	obj := Constructor()
	//input := []int{1, 3, 7, 2, 6}
	input := []int{1, 3, 7, 2, 6, 9, 4, 10, 5}
	for _, v := range input {
		obj.AddNum(v)
		fmt.Printf("%v\n", obj.GetIntervals())
	}
}
