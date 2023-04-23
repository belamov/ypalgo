package main

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

//https://contest.yandex.ru/contest/24809/problems/J/
// Дано BST. Надо вставить узел с заданным ключом. Ключи в дереве могут повторяться.
//На вход функции подаётся корень корректного бинарного дерева поиска и ключ, который надо вставить в дерево.
//Осуществите вставку этого ключа. Если ключ уже есть в дереве, то его дубликаты уходят в правого сына.
//Таким образом вид дерева после вставки определяется однозначно.
//Функция должна вернуть корень дерева после вставки вершины.
//Ваше решение должно работать за O(h), где h — высота дерева.
func insert(root *Node, key int) *Node {
	if key < root.value {
		if root.left == nil {
			root.left = &Node{value: key}
		} else {
			insert(root.left, key)
		}
	} else {
		if root.right == nil {
			root.right = &Node{value: key}
		} else {
			insert(root.right, key)
		}
	}
	return root
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		panic("WA")
	}
	if newHead.left.value != 6 {
		panic("WA")
	}
}
