package main

type ListNode struct {
	data string
	next *ListNode
}

// https://contest.yandex.ru/contest/22779/problems/C/

//Вася размышляет, что ему можно не делать из того списка дел, который он составил.
//Но, кажется, все пункты очень важные! Вася решает загадать число и удалить дело, которое
//идёт под этим номером. Список дел представлен в виде односвязного списка.
//
//Напишите функцию solution, которая принимает на вход голову
//списка и номер удаляемого дела и возвращает голову обновлённого списка.
func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}

	n := 0
	nodeToDelete := head
	for nodeToDelete.next != nil {
		if n == idx-1 {
			nodeToDelete.next = nodeToDelete.next.next
			break
		}
		nodeToDelete = nodeToDelete.next
		n++
	}

	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
