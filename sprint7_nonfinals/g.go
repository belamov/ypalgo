package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/G/
// Тимофей пошёл снять деньги в банкомат.
// Ему нужно m франков.
// В банкомате в бесконечном количестве имеются купюры различных достоинств.
// Всего различных достоинств n. Купюр каждого достоинства можно взять бесконечно много.
// Нужно определить число способов, которыми Тимофей сможет набрать нужную сумму.
func main() {
	scanner := makeScanner()
	m := readInt(scanner)
	readInt(scanner)
	banknotes := readArray(scanner)

	dp := make([][]int, len(banknotes))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < len(banknotes); i++ {
		if banknotes[i] <= m {
			dp[i][banknotes[i]] = 1
		}
		for j := 1; j <= m; j++ {
			a := 0
			if i > 0 {
				a = dp[i-1][j]
			}
			b := 0
			if j-banknotes[i] >= 0 {
				b = dp[i][j-banknotes[i]]
			}
			dp[i][j] += a + b
		}
	}
	fmt.Print(dp[len(banknotes)-1][m])
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
