package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/E/
//Гоша измерял температуру воздуха n дней подряд.
//В результате у него получился некоторый временной ряд.
//Теперь он хочет посмотреть, как часто встречается некоторый шаблон в получившейся последовательности.
//Однако температура — вещь относительная, поэтому Гоша решил, что при поиске
//шаблона длины m (a1, a2, ..., am) стоит также рассматривать сдвинутые на константу вхождения.
//Это значит, что если для некоторого числа c в исходной последовательности
//нашёлся участок вида (a1 + c, a2 + c, ... , am + c), то он тоже считается вхождением шаблона (a1, a2, ..., am).
//
//По заданной последовательности измерений X и шаблону A=(a1, a2, ..., am) определите
//все вхождения A в X, допускающие сдвиг на константу.
func main() {
	scanner := makeScanner()
	readInt(scanner)
	x := readArray(scanner)
	readInt(scanner)
	a := readArray(scanner)

	start := 0
	for {
		pos := find(x, a, start)
		if pos == -1 {
			break
		}
		fmt.Print(pos+1, " ")
		start = pos + 1
	}
}

func find(arr []int, pattern []int, start int) int {
	if len(arr) < len(pattern) {
		return -1
	}

	for pos := start; pos < len(arr)-len(pattern)+1; pos++ {
		match := true
		diff := pattern[0] - arr[pos]
		for offset := 0; offset < len(pattern); offset++ {
			if pattern[offset]-arr[pos+offset] != diff {
				match = false
				break
			}
		}
		if match {
			return pos
		}
	}
	return -1
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
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
