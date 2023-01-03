package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/I/

// Астрологи объявили день очередей ограниченного размера.
//Тимофею нужно написать класс LinkedListQueue, который принимает
//параметр max_size, означающий максимально допустимое количество элементов в очереди.
//
//Помогите ему —– реализуйте программу, которая будет
//эмулировать работу такой очереди. Функции, которые надо поддержать, описаны в формате ввода.
type MyQueueSized struct {
	queue       []int
	maxSize     int
	head        int
	tail        int
	currentSize int
}

func newMyQueueSized(maxSize int) *MyQueueSized {
	return &MyQueueSized{
		queue:   make([]int, maxSize),
		maxSize: maxSize,
		head:    0,
		tail:    0,
	}
}

func (q *MyQueueSized) push(x int) {
	if q.currentSize == q.maxSize {
		fmt.Println("error")
		return
	}
	q.queue[q.head] = x
	q.head = (q.head + 1) % q.maxSize
	q.currentSize++
}

func (q *MyQueueSized) pop() {
	q.peek()
	if q.currentSize == 0 {
		return
	}
	q.tail = (q.tail + 1) % q.maxSize
	q.currentSize--
}

func (q *MyQueueSized) peek() {
	if q.currentSize == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(q.queue[q.tail])
}

func (q *MyQueueSized) size() {
	fmt.Println(q.currentSize)
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))
	queueSize, _ := strconv.Atoi(readLine(scanner))

	queue := newMyQueueSized(queueSize)

	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, queue)
	}

}

func runCommand(scanner *bufio.Scanner, queue *MyQueueSized) {
	command := readLine(scanner)

	if strings.Contains(command, "push") {
		arg, _ := strconv.Atoi(command[5:])
		queue.push(arg)
		return
	}

	if command == "pop" {
		queue.pop()
		return
	}

	if command == "peek" {
		queue.peek()
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
