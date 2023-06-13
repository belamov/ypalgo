package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/D/
//Найдите наибольший по длине общий префикс нескольких строк.
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = readString(scanner)
	}

	fmt.Print(len(longestCommonPrefix(strs)))
}

func longestCommonPrefix(strs []string) string {
	var commonPrefix strings.Builder
	for i := range strs[0] {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return commonPrefix.String()
			}
		}
		commonPrefix.WriteByte(strs[0][i])
	}

	return commonPrefix.String()
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
