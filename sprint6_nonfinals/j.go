package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/J/
// Дан ациклический ориентированный граф (так называемый DAG, directed acyclic graph).
//Найдите его топологическую сортировку, то есть выведите его вершины в таком
//порядке, что все рёбра графа идут слева направо.
//У графа может быть несколько подходящих перестановок вершин.
//Вам надо найти любую топологическую сортировку.
func main() {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	adjList := make([][]int, n)

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)

		if adjList[vertices[0]-1] == nil {
			adjList[vertices[0]-1] = make([]int, 0)
		}
		adjList[vertices[0]-1] = append(adjList[vertices[0]-1], vertices[1]-1)
	}

	for _, ints := range adjList {
		sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	}

	color := make([]int, len(adjList))
	sorted := make([]int, 0, len(adjList))
	var stack []int

	for i := 0; i < len(color); i++ {
		if color[i] != 0 {
			continue
		}

		stack = []int{i}

		for len(stack) > 0 {
			v := stack[len(stack)-1]     // Получаем из стека очередную вершину.
			stack = stack[:len(stack)-1] // Удаляем вершину из стека.

			if color[v] == 0 {
				// Красим вершину в серый. И сразу кладём её обратно в стек:
				// это позволит алгоритму позднее вспомнить обратный путь по графу.
				color[v] = 1

				stack = append(stack, v)
				// Теперь добавляем в стек все непосещённые соседние вершины,
				// вместо вызова рекурсии
				for _, w := range adjList[v] { // Перебираем смежные вершины.
					if color[w] == 0 { // Если вершина не посещена, то
						stack = append(stack, w)
					}
				}
				continue
			}

			if color[v] == 1 {
				// Серую вершину мы могли получить из стека только на обратном пути.
				// Следовательно, её следует перекрасить в чёрный.
				color[v] = 2
				sorted = append(sorted, v)
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := len(sorted) - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(sorted[i] + 1))
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
