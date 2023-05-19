package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/F/
//Найдите кратчайшее расстояние между парой вершин в неориентированном графе. Граф может быть несвязным.
func main() {
	scanner := makeScanner()
	graph := getGraph(scanner)
	startEnd := readArray(scanner)

	start := startEnd[0] - 1
	end := startEnd[1] - 1

	fmt.Print(getMinDistance(graph, start, end))
}

func getMinDistance(graph [][]int, start int, end int) int {
	color := make([]int, len(graph))
	distance := make([]int, len(graph))

	// initial queue of vertex to visit
	planned := []int{start}

	color[start] = 1
	distance[start] = 0

	// bfs
	// while queue of planned for visit nodes is not empty
	for len(planned) > 0 {

		// pop from queue
		u := planned[0]
		planned = planned[1:]

		// iterate all adjacent nodes
		for _, v := range graph[u] {
			// if adjacent node has not been visited
			if color[v] == 0 {
				color[v] = 1
				distance[v] = distance[u] + 1
				//adding it to queue, so we can search its adjacent vertex
				planned = append(planned, v)
			}
		}
	}

	if color[end] == 0 {
		return -1
	}

	return distance[end]
}

func getGraph(scanner *bufio.Scanner) [][]int {
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
	return adjList
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
