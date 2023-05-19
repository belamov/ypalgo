package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/25069/problems/L/
//Неориентированный граф называется полным, если в нём каждая пара вершин соединена ребром.
//Вам дан неориентированный граф из n вершин и m рёбер. Выясните, является ли этот граф полным.
func main() {
	graph := getGraph()

	for v := range graph {
		if len(graph[v]) != len(graph)-1 {
			fmt.Print("NO")
			return
		}
	}

	fmt.Print("YES")
}

func getGraph() []map[int]bool {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]
	graph := make([]map[int]bool, n)

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)

		if vertices[0] == vertices[1] {
			continue
		}

		v := vertices[0] - 1
		u := vertices[1] - 1

		if graph[v] == nil {
			graph[v] = make(map[int]bool)
		}
		graph[v][u] = true

		if graph[u] == nil {
			graph[u] = make(map[int]bool)
		}
		graph[u][v] = true
	}

	return graph
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
