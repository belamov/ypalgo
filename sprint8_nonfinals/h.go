package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/H/
//Напишите программу, которая будет заменять в тексте все вхождения строки s на строку t.
//Гарантируется, что никакие два вхождения шаблона s не пересекаются друг с другом.
func main() {
	scanner := makeScanner()
	text := readString(scanner)
	pattern := readString(scanner)
	replace := readString(scanner)

	foundPositions := find(pattern, text)

	if len(foundPositions) == 0 {
		fmt.Print(text)
		return
	}

	var result strings.Builder
	offset := 0
	foundIndex := 0
	for i := 0; i < len(text); i++ {
		if i == foundPositions[foundIndex] {
			result.WriteString(replace)
			offset += len(replace)
			if foundIndex != len(foundPositions)-1 {
				foundIndex++
			}
			i += len(pattern) - 1
		} else {
			result.WriteString(string(text[i]))
		}
	}

	fmt.Print(result.String())
}

func find(pattern, text string) []int {
	result := make([]int, 0)
	s := pattern + "#" + text
	p := make([]int, len(pattern))
	pprev := 0
	for i := 1; i < len(s); i++ {
		k := pprev
		for k > 0 && s[k] != s[i] {
			k = p[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		if i < len(p) {
			p[i] = k
		}
		pprev = k

		if k == len(pattern) {
			result = append(result, i-2*len(p))
		}
	}
	return result
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
