package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

// https://contest.yandex.ru/contest/22779/problems/B/

//Вася размышляет, что ему можно не делать из того списка дел, который он составил.
//Но, кажется, все пункты очень важные! Вася решает загадать число и удалить дело, которое
//идёт под этим номером. Список дел представлен в виде односвязного списка.
//Напишите функцию solution, которая принимает на вход голову списка и номер
//удаляемого дела и возвращает голову обновлённого списка.
func Solution(head *ListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	Solution(&node0)
	/*
	   Output is:
	   node0
	   node1
	   node2
	   node3
	*/
}
