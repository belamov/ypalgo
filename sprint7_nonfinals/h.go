package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/25596/problems/H/
//Черепаха Кондратина путешествует по клетчатому полю из n строк и m столбцов.
//В каждой клетке либо растёт цветочек, либо не растёт.
//Кондратине надо добраться из левого нижнего в правый верхний угол и собрать как можно больше цветочков.
//
//Помогите ей с этой сложной задачей и определите, какое наибольшее число
//цветочков она сможет собрать при условии, что Кондратина умеет передвигаться
//только на одну клетку вверх или на одну клетку вправо за ход.
func main() {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	flowers := readMatrix(scanner, n, m)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[n-1][0] = flowers[n-1][0]

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			dp[i][j] = max(getPrevValueAtIndex(dp, i+1, j), getPrevValueAtIndex(dp, i, j-1)) + flowers[i][j]
		}
	}

	fmt.Print(dp[0][m-1])
}

func getPrevValueAtIndex(dp [][]int, i int, j int) int {
	if i >= len(dp) || j < 0 {
		return 0
	}
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		row := readString(scanner)
		for j, number := range row {
			matrix[i][j] = int(number) - 48
		}
	}
	return matrix
}
