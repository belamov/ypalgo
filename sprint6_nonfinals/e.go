package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/E/
// Вам дан неориентированный граф. Найдите его компоненты связности.
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

		if adjList[vertices[1]-1] == nil {
			adjList[vertices[1]-1] = make([]int, 0)
		}
		adjList[vertices[1]-1] = append(adjList[vertices[1]-1], vertices[0]-1)
	}

	color := make([]int, len(adjList))

	groups := make([][]int, 0)
	currentGroupIndex := 0

	var stack []int

	for i := 0; i < len(color); i++ {
		if color[i] != 0 {
			continue
		}

		groups = append(groups, make([]int, 0))

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
			}

			if color[v] == 1 {
				// Серую вершину мы могли получить из стека только на обратном пути.
				// Следовательно, её следует перекрасить в чёрный.
				color[v] = 2
				groups[currentGroupIndex] = append(groups[currentGroupIndex], v)
			}
		}
		currentGroupIndex++
	}

	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString(strconv.Itoa(len(groups)))
	writer.WriteString("\n")

	for _, group := range groups {
		sort.Ints(group)
		for _, v := range group {
			writer.WriteString(strconv.Itoa(v + 1))
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
