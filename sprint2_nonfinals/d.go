package main

type ListNode struct {
	data string
	next *ListNode
}

// https://contest.yandex.ru/contest/22779/problems/D/

// Мама Васи хочет знать, что сын планирует делать и когда.
//Помогите ей: напишите функцию solution, определяющую индекс
//первого вхождения передаваемого ей на вход значения в связном списке, если значение присутствует.
func Solution(head *ListNode, elem string) int {
	idx := 0
	currentNode := head

	for currentNode != nil {
		if currentNode.data == elem {
			return idx
		}
		currentNode = currentNode.next
		idx++
	}

	return -1
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
