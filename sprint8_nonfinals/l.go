package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/26131/problems/L/
//В этой задаче вам необходимо посчитать префикс-функцию для заданной строки.
func main() {
	scanner := makeScanner()
	s := readString(scanner)

	p := prefixFunction(s)
	printArray(p)
}

func prefixFunction(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		k := p[i-1]
		for k > 0 && s[k] != s[i] {
			k = p[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		p[i] = k
	}
	return p
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
