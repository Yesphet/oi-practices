package main

import "fmt"

type Node struct {
	left     *Node
	right    *Node
	key      int
	height   int
	childNum int
	num      int
}

func newNode(key int) *Node {
	return &Node{
		left:     nil,
		right:    nil,
		key:      key,
		height:   1,
		num:      1,
		childNum: 0,
	}
}

// insert and return the number of nodes whose key greater than the key
func (n *Node) insert(key int) (*Node, int) {
	if key == n.key {
		n.num++
		if n.left != nil {
			return n, n.left.num + n.left.childNum
		}
		return n, 0
	}

	var count int
	if key < n.key {
		if n.left == nil {
			n.left = newNode(key)
		} else {
			n.left, count = n.left.insert(key)
		}
	} else {
		if n.right == nil {
			n.right = newNode(key)
		} else {
			n.right, count = n.right.insert(key)
		}
		count += n.num
		if n.left != nil {
			count += n.left.childNum + n.left.num
		}
	}

	//n.childNum++
	//return n, count
	switch n.balanceFactor() {
	case 2:
		if n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotate()
		}
		return n.rightRotate(), count
	case -2:
		if n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotate()
		}
		return n.leftRotate(), count
	default:
		n.updateHeight()
		return n, count
	}
}

func (n *Node) rightRotate() *Node {
	b := n.left
	br := b.right

	b.right = n
	n.left = br

	n.updateHeight()
	b.updateHeight()
	return b
}

func (n *Node) leftRotate() *Node {
	b := n.right
	bl := b.left

	b.left = n
	n.right = bl

	n.updateHeight()
	b.updateHeight()
	return b
}

func (n *Node) updateHeight() {
	lh, rh, childNum := 0, 0, 0
	if n.left != nil {
		lh = n.left.height
		childNum += n.left.childNum + n.left.num
	}
	if n.right != nil {
		rh = n.right.height
		childNum += n.right.childNum + n.right.num
	}
	n.height = max(lh, rh) + 1
	n.childNum = childNum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func (n *Node) balanceFactor() int {
	lh, rh := 0, 0
	if n.left != nil {
		lh = n.left.height
	}
	if n.right != nil {
		rh = n.right.height
	}
	return lh - rh
}

func countSmaller(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	l := len(nums)
	root := newNode(nums[l-1])
	nums[l-1] = 0
	for i := l - 2; i >= 0; i-- {
		root, nums[i] = root.insert(nums[i])
	}

	return nums
}

func main() {
	//input := []int{5, 2, 6, 1}
	input := []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}
	out := countSmaller(input)
	fmt.Println(out)
}
