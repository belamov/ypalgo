package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/A/
//Алла пошла на стажировку в студию графического дизайна, где ей дали такое
//задание: для очень большого числа ориентированных графов преобразовать их список
//рёбер в список смежности.
//Чтобы побыстрее решить эту задачу, она решила автоматизировать процесс.
//
//Помогите Алле написать программу, которая по списку рёбер графа будет строить его список смежности.
func main() {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	edges := make([][]int, n)

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)

		if edges[vertices[0]-1] == nil {
			edges[vertices[0]-1] = make([]int, 0)
		}
		edges[vertices[0]-1] = append(edges[vertices[0]-1], vertices[1])
	}

	writer := bufio.NewWriter(os.Stdout)

	for _, vertex := range edges {
		if vertex == nil {
			writer.WriteString("0\n")
			continue
		}
		writer.WriteString(strconv.Itoa(len(vertex)) + " ")
		for i := 0; i < len(vertex); i++ {
			writer.WriteString(strconv.Itoa(vertex[i]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
