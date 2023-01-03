package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	queue       []int
	maxSize     int
	head        int
	tail        int
	currentSize int
}

func newMyQueueSized(maxSize int) *Deque {
	return &Deque{
		queue:   make([]int, maxSize),
		maxSize: maxSize,
		head:    0,
		tail:    0,
	}
}

func (d *Deque) pushFront(x int) {
	if d.currentSize == d.maxSize {
		fmt.Println("error")
		return
	}
	d.queue[d.head] = x
	d.head = (d.head + 1) % d.maxSize
	d.currentSize++
}

func (d *Deque) pushBack(x int) {
	if d.currentSize == d.maxSize {
		fmt.Println("error")
		return
	}
	d.tail = (d.tail - 1 + d.maxSize) % d.maxSize
	d.queue[d.tail] = x

	d.currentSize++
}

func (d *Deque) popFront() {
	if d.currentSize == 0 {
		fmt.Println("error")
		return
	}
	d.head = (d.head - 1 + d.maxSize) % d.maxSize
	fmt.Println(d.queue[d.head])
	d.currentSize--
}

func (d *Deque) popBack() {
	if d.currentSize == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(d.queue[d.tail])
	d.tail = (d.tail + 1) % d.maxSize
	d.currentSize--
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

func runCommand(scanner *bufio.Scanner, queue *Deque) {
	command := readLine(scanner)

	if strings.Contains(command, "push_front") {
		arg, _ := strconv.Atoi(command[11:])
		queue.pushFront(arg)
		return
	}
	if strings.Contains(command, "push_back") {
		arg, _ := strconv.Atoi(command[10:])
		queue.pushBack(arg)
		return
	}

	if command == "pop_front" {
		queue.popFront()
		return
	}

	if command == "pop_back" {
		queue.popBack()
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
