package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/C/
// Задан неориентированный граф.
//Обойдите с помощью DFS все вершины, достижимые из заданной вершины s, и выведите
//их в порядке обхода, если начинать обход из s.
//Выведите вершины в порядке обхода, считая что при запуске от каждой
//конкретной вершины её соседи будут рассматриваться в порядке
//возрастания (то есть если вершина 2 соединена с 1 и 3, то сначала обход пойдёт в 1, а уже потом в 3).
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

	for _, ints := range adjList {
		sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	}

	startVertex := readInt(scanner)

	color := make([]int, len(adjList))

	DFS(startVertex, adjList, color)
}

func DFS(startVertex int, adjList [][]int, color []int) {
	stack := []int{startVertex - 1}

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
			fmt.Print(v+1, " ")
		}
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
