package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/F/

// Нужно реализовать класс StackMaxEffective, который поддерживает операцию
//определения максимума среди всех элементов в стеке.
//Класс должен поддерживать операции push(x), где x – целое число, pop() и get_max().
type StackMax struct {
	stack []int
	max   int
}

func newStackMax() *StackMax {
	return &StackMax{stack: []int{}, max: math.MinInt}
}

func (s *StackMax) push(x int) {
	s.stack = append(s.stack, x)
	if x > s.max {
		s.max = x
	}
}

func (s *StackMax) pop() {
	if len(s.stack) == 0 {
		fmt.Println("error")
		return
	}

	popped := s.stack[len(s.stack)-1]

	s.stack = s.stack[:len(s.stack)-1]

	if len(s.stack) == 0 {
		s.max = math.MinInt
		return
	}

	if s.max == popped {
		s.refreshMax()
	}
}

func (s *StackMax) getMax() {
	if len(s.stack) == 0 {
		fmt.Println("None")
		return
	}
	fmt.Println(s.max)
}

func (s *StackMax) refreshMax() {
	s.max = math.MinInt
	for _, x := range s.stack {
		if x > s.max {
			s.max = x
		}
	}
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))

	stack := newStackMax()

	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, stack)
	}

}

func runCommand(scanner *bufio.Scanner, stack *StackMax) {
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
