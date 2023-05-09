package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/D/
// Задан неориентированный граф.
//Обойдите поиском в ширину все вершины, достижимые из заданной
//вершины s, и выведите их в порядке обхода, если начинать обход из s.
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

	s := readInt(scanner)

	color := make([]int, len(adjList))

	// Создадим очередь вершин и положим туда стартовую вершину.
	planned := []int{s - 1}
	color[s-1] = 1
	fmt.Print(s, " ")

	for len(planned) > 0 {
		u := planned[0] // Возьмём вершину из очереди.
		planned = planned[1:]
		sort.Sort(sort.IntSlice(adjList[u]))
		for _, v := range adjList[u] { // adjList - список смежности графа.
			if color[v] == 0 { // Серые и чёрные вершины уже
				// либо в очереди, либо обработаны.
				color[v] = 1
				planned = append(planned, v) // Запланируем посещение вершины.
				fmt.Print(v+1, " ")
			}
		}
		color[u] = 2 // Теперь вершина считается обработанной.
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
