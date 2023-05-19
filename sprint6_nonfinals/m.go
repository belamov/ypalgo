package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/M/
//Определение: неориентированный граф называется двудольным (англ. bipartite), если его
//вершины можно разбить на два непересекающихся множества таким образом, что рёбра
//будут проведены только между вершинами из разных множеств.
//Эти два множества вершин ещё называют долями.

//Гоша узнал, что двудольными могут быть не только графы, но и растения (например, сирень).
//Теперь он в них путается и не может проверить граф на двудольность без мыслей о цветочках.
//Помогите Гоше: проверьте, является ли заданный неориентированный граф двудольным.
func main() {
	adjList := getGraph()

	if isBipartite(adjList) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func getGraph() [][]int {
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
	return adjList
}

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))

	// bfs search
	for v := range color {

		// vertex is already visited
		if color[v] != 0 {
			continue
		}

		// initial queue of vertex to visit
		planned := []int{v}

		// marking vertex as red (red = 1, blue = 2)
		color[v] = 1

		// while queue of planned for visit nodes is not empty
		for len(planned) > 0 {

			// pop from queue
			u := planned[0]
			planned = planned[1:]

			// iterate all adjacent nodes
			for _, v := range graph[u] {
				// if adjacent node has not been visited
				if color[v] == 0 {
					// marking it as opposite group
					color[v] = getAdjacentGroupColor(color[u])

					//adding it to queue, so we can search its adjacent vertex
					planned = append(planned, v)
				}

				// if current and adjacent node are in same group, then graph is not bipartite
				if color[v] == color[u] {
					return false
				}
			}
		}
	}

	return true
}

func getAdjacentGroupColor(currentGroupColor int) int {
	if currentGroupColor == 1 {
		return 2
	}
	return 1
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
