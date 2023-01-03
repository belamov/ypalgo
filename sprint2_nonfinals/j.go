package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/J/

// Любимый вариант очереди Тимофея — очередь, написанная с использованием связного списка.
//Помогите ему с реализацией. Очередь должна поддерживать выполнение трёх команд:
//
//get() — вывести элемент, находящийся в голове очереди, и удалить его. Если очередь пуста, то вывести «error».
//put(x) — добавить число x в очередь
//size() — вывести текущий размер очереди
type LinkedListQueue struct {
	head        *ListNode
	tail        *ListNode
	currentSize int
}

func newLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) put(x int) {
	element := &ListNode{
		data: x,
		prev: q.head,
		next: nil,
	}

	if q.head != nil {
		q.head.next = element
	}

	q.head = element

	if q.tail == nil {
		q.tail = q.head
	}

	q.currentSize++
}

func (q *LinkedListQueue) get() {
	if q.tail == nil {
		fmt.Println("error")
		return
	}

	fmt.Println(q.tail.data)

	q.tail = q.tail.next
	q.currentSize--
}

func (q *LinkedListQueue) size() {
	fmt.Println(q.currentSize)
}

type ListNode struct {
	data int
	prev *ListNode
	next *ListNode
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))

	queue := newLinkedListQueue()

	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, queue)
	}

}

func runCommand(scanner *bufio.Scanner, queue *LinkedListQueue) {
	command := readLine(scanner)

	if strings.Contains(command, "put") {
		arg, _ := strconv.Atoi(command[4:])
		queue.put(arg)
		return
	}

	if command == "get" {
		queue.get()
		return
	}

	if command == "size" {
		queue.size()
		return
	}

}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
