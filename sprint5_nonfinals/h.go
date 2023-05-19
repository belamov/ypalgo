package main

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func Solution(root *Node) int {
	return getSum(root, 0)
}

func getSum(root *Node, num int) int {
	if root == nil {
		return 0
	}

	num = num*10 + root.value
	if root.left == nil && root.right == nil {
		return num
	}

	return getSum(root.left, num) + getSum(root.right, num)
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, nil, nil}
	node5 := Node{1, &node4, &node3}
	if Solution(&node5) != 275 {
		panic("WA")
	}
}
