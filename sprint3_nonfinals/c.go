package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/23638/problems/C/

// Гоша любит играть в игру «Подпоследовательность»: даны 2 строки, и нужно понять, является ли
//первая из них подпоследовательностью второй.
//Когда строки достаточно длинные, очень трудно получить ответ на этот вопрос, просто посмотрев на них.
//Помогите Гоше написать функцию, которая решает эту задачу.
func isSubsequence(s, t string) bool {
	sPos, tPos := 0, 0
	for sPos != len(s) && tPos != len(t) {
		if s[sPos] == t[tPos] {
			sPos++
			tPos++
			continue
		}
		tPos++
	}
	return sPos == len(s)
}

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	isSubseq := isSubsequence(s, t)
	if isSubseq {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
