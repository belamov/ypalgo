package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/F/
//Алла хочет доказать, что она умеет прыгать вверх по лестнице быстрее всех.
//На этот раз соревнования будут проходить на специальной прыгательной лестнице.
//С каждой её ступеньки можно прыгнуть вверх на любое расстояние от 1 до k.
//Число k придумывает Алла.
//
//Гоша не хочет проиграть, поэтому просит вас посчитать количество способов допрыгать от первой ступеньки до n-й.
//Изначально все стоят на первой ступеньке.
func main() {
	scanner := makeScanner()
	nk := readArray(scanner)
	n := nk[0]
	k := nk[1]

	dp := make([]int, n)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := i; j > i-k && j > 0; j-- {
			dp[i] += dp[j-1]
			dp[i] = dp[i] % 1000000007
		}
	}

	fmt.Print(dp[n-1])
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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
