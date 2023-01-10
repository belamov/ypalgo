package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/23638/problems/A/

// Рита по поручению Тимофея наводит порядок в правильных скобочных
// последовательностях (ПСП), состоящих только из круглых скобок ().
// Для этого ей надо сгенерировать все ПСП длины 2n в алфавитном
// порядке —– алфавит состоит из ( и ) и открывающая скобка идёт раньше закрывающей.
//
// Помогите Рите —– напишите программу, которая по заданному n выведет все ПСП в нужном порядке.
func generateBraces(open, closed, n int, prefix string) {
	if open == closed && closed == n {
		fmt.Println(prefix)
		return
	}
	if open < n {
		generateBraces(open+1, closed, n, prefix+"(")
	}

	if closed < open {
		generateBraces(open, closed+1, n, prefix+")")
	}
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	generateBraces(1, 0, n, "(")
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
