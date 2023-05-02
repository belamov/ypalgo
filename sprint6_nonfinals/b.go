package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/B/
//Алла успешно справилась с предыдущим заданием, и теперь ей дали новое.
//На этот раз список рёбер ориентированного графа надо переводить в матрицу смежности. К
//онечно же, Алла попросила вас помочь написать программу для этого.
func main() {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)
		matrix[vertices[0]-1][vertices[1]-1] = 1
	}

	printMatrix(matrix)
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

func printMatrix(matrix [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			writer.WriteString(strconv.Itoa(matrix[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}
