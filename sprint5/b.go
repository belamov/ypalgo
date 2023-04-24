package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

//https://contest.yandex.ru/contest/24809/problems/B/
//https://leetcode.com/problems/balanced-binary-tree/description/
//Гоше очень понравилось слушать рассказ Тимофея про деревья.
//Особенно часть про сбалансированные деревья.
//Он решил написать функцию, которая определяет, сбалансировано ли дерево.
//Дерево считается сбалансированным, если левое и правое поддеревья каждой
//вершины отличаются по высоте не больше, чем на единицу.
func Solution(root *Node) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *Node) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.left)
	isRightBalanced, rightHeight := dfs(root.right)

	return isLeftBalanced && isRightBalanced && (abs(leftHeight-rightHeight) < 2), 1 + max(leftHeight, rightHeight)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{0, nil, &node2}
	fmt.Println(Solution(&node3))
}
