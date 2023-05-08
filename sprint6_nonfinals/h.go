package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/H/
// Вам дан ориентированный граф.
//Известно, что все его вершины достижимы из вершины s=1.
//Найдите время входа и выхода при обходе в глубину, производя первый запуск из вершины s.
//Считайте, что время входа в стартовую вершину равно 0.
//Соседей каждой вершины обходите в порядке увеличения номеров.
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
	time := 0
	entry := make([]int, len(adjList))
	leave := make([]int, len(adjList))

	stack := []int{0}

	for len(stack) > 0 {
		v := stack[len(stack)-1]     // Получаем из стека очередную вершину.
		stack = stack[:len(stack)-1] // Удаляем вершину из стека.

		if color[v] == 0 {
			// Красим вершину в серый. И сразу кладём её обратно в стек:
			// это позволит алгоритму позднее вспомнить обратный путь по графу.
			color[v] = 1
			entry[v] = time
			time++

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
			leave[v] = time
			time++
		}
	}

	for i := 0; i < len(entry); i++ {
		fmt.Println(fmt.Sprintf("%d %d", entry[i], leave[i]))
	}
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
