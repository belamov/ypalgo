package main

/**
Comment it before submitting
type Node struct {
	value  int
	left   *Node
	right  *Node
}
**/

func Solution(root *Node) int {
	return maxInSubtree(root)
}

func maxInSubtree(root *Node) int {
	if root.left == nil && root.right == nil {
		return root.value
	}

	max := root.value
	maxLeft := max
	maxRight := max
	if root.left != nil {
		maxLeft = maxInSubtree(root.left)
	}
	if root.right != nil {
		maxRight = maxInSubtree(root.right)
	}
	if maxLeft > max {
		max = maxLeft
	}
	if maxRight > max {
		max = maxRight
	}
	return max
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if Solution(&node4) != 3 {
		panic("WA")
	}
}
