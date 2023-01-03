package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/22779/problems/H/

// Вот какую задачу Тимофей предложил на собеседовании одному из кандидатов.
//Если вы с ней ещё не сталкивались, то наверняка столкнётесь –— она довольно популярная.
//
//Дана скобочная последовательность. Нужно определить, правильная ли она.
//
//Будем придерживаться такого определения:
//
//пустая строка —– правильная скобочная последовательность;
//правильная скобочная последовательность, взятая в скобки одного типа, –— правильная скобочная последовательность;
//правильная скобочная последовательность с приписанной слева или справа правильной скобочной последовательностью —– тоже правильная.
//На вход подаётся последовательность из скобок трёх видов: [], (), {}.
//
//Напишите функцию is_correct_bracket_seq, которая принимает на вход
//скобочную последовательность и возвращает True, если последовательность правильная, а иначе False.
func isCorrectBracketSeq(seq string) {
	var stack *ListNode = nil

	for i := 0; i < len(seq); i++ {
		bracket := string(seq[i])

		if bracket == "(" || bracket == "[" || bracket == "{" {
			newNode := &ListNode{
				data: bracket,
				prev: stack,
			}
			stack = newNode
			continue
		}

		if stack == nil {
			fmt.Println("False")
			return
		}

		if bracket == ")" && stack.data != "(" {
			fmt.Println("False")
			return
		}

		if bracket == "]" && stack.data != "[" {
			fmt.Println("False")
			return
		}

		if bracket == "}" && stack.data != "{" {
			fmt.Println("False")
			return
		}

		stack = stack.prev
	}

	if stack != nil {
		fmt.Println("False")
		return
	}

	fmt.Println("True")

}

type ListNode struct {
	data string
	prev *ListNode
}

func main() {
	scanner := makeScanner()
	brackets := readLine(scanner)
	isCorrectBracketSeq(brackets)
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
