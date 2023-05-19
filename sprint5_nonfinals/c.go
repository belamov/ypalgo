package main

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}

	return isMirrored(root.left, root.right)
}

func isMirrored(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.value != root2.value {
		return false
	}

	return isMirrored(root1.left, root2.right) && isMirrored(root1.right, root2.left)
}

func test() {
	node1 := Node{3, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{4, nil, nil}
	node4 := Node{3, nil, nil}
	node5 := Node{2, &node1, &node2}
	node6 := Node{2, &node3, &node4}
	node7 := Node{1, &node5, &node6}

	if !Solution(&node7) {
		panic("WA")
	}
}
