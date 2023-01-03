package main

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

//https://contest.yandex.ru/contest/22779/problems/E/

// Вася решил запутать маму —– делать дела в обратном порядке.
//Список его дел теперь хранится в двусвязном списке.
//Напишите функцию, которая вернёт список в обратном порядке.
func Solution(head *ListNode) *ListNode {
	currentNode := head
	for currentNode.next != nil {
		nextNode := currentNode.next
		currentNode.next = currentNode.prev
		currentNode.prev = nextNode
		currentNode = nextNode
	}

	currentNode.next = currentNode.prev
	currentNode.prev = nil

	return currentNode
}

func test() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	/*newHead :=*/ Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.next = node3
	// node1.next == node0
	// node1.next == node2
	// node0.next == node1
}
