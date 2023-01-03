package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/G/

// Реализуйте класс StackMaxEffective, поддерживающий операцию определения
//максимума среди элементов в стеке. Сложность операции должна быть O(1).
//Для пустого стека операция должна возвращать None.
//При этом push(x) и pop() также должны выполняться за константное время.
type StackMaxEffective struct {
	lastElement *ListNode
}

func newStackMax() *StackMaxEffective {
	return &StackMaxEffective{lastElement: nil}
}

func (s *StackMaxEffective) push(x int) {
	element := &ListNode{
		data: x,
		max:  x,
		prev: s.lastElement,
	}

	if s.lastElement != nil && s.lastElement.max > element.max {
		element.max = s.lastElement.max
	}

	s.lastElement = element

}

func (s *StackMaxEffective) pop() {
	if s.lastElement == nil {
		fmt.Println("error")
		return
	}

	s.lastElement = s.lastElement.prev
}

func (s *StackMaxEffective) getMax() {
	if s.lastElement == nil {
		fmt.Println("None")
		return
	}
	fmt.Println(s.lastElement.max)
}

type ListNode struct {
	data int
	prev *ListNode
	max  int
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))

	stack := newStackMax()

	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, stack)
	}

}

func runCommand(scanner *bufio.Scanner, stack *StackMaxEffective) {
	command := readLine(scanner)

	if strings.Contains(command, "push") {
		arg, _ := strconv.Atoi(command[5:])
		stack.push(arg)
		return
	}

	if command == "pop" {
		stack.pop()
		return
	}

	if command == "get_max" {
		stack.getMax()
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
