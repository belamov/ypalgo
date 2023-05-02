package main

import "fmt"

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

//https://contest.yandex.ru/contest/24809/problems/K/
//https://leetcode.com/problems/trim-a-binary-search-tree/description/

//Напишите функцию, которая будет выводить по неубыванию все ключи от L до R включительно в заданном бинарном дереве поиска.
//Ключи в дереве могут повторяться. Решение должно иметь сложность O(h+k), где
//h — глубина дерева,
//k — число элементов в ответе.
//В данной задаче если в узле содержится ключ x, то другие ключи, равные x, могут быть
//как в правом, так и в левом поддереве данного узла. (Дерево строил стажёр, так что ничего страшного).
func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}

	if root.value < left {
		printRange(root.right, left, right)
		return
	}

	if root.value > right {
		printRange(root.left, left, right)
		return
	}

	printRange(root.left, left, right)
	fmt.Println(root.value)
	printRange(root.right, left, right)

}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}
