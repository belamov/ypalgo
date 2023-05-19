package main

type Node struct {
	value int
	left  *Node
	right *Node
}

var maxPath int

func Solution(root *Node) int {
	maxPath = root.value
	getMaxPath(root)
	return maxPath
}

func getMaxPath(root *Node) int {
	if root == nil {
		return 0
	}

	leftMaxPathWithoutSplit := getMaxPath(root.left)
	rightMaxPathWithoutSplit := getMaxPath(root.right)

	leftMaxPathWithoutSplit = max(leftMaxPathWithoutSplit, 0)
	rightMaxPathWithoutSplit = max(rightMaxPathWithoutSplit, 0)

	if root.value+leftMaxPathWithoutSplit+rightMaxPathWithoutSplit > maxPath {
		maxPath = root.value + leftMaxPathWithoutSplit + rightMaxPathWithoutSplit
	}

	return root.value + max(leftMaxPathWithoutSplit, rightMaxPathWithoutSplit)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test() {
	node1 := Node{5, nil, nil}
	node2 := Node{1, nil, nil}
	node3 := Node{-3, &node2, &node1}
	node4 := Node{2, nil, nil}
	node5 := Node{2, &node4, &node3}
	if Solution(&node5) != 6 {
		panic("WA")
	}
}
