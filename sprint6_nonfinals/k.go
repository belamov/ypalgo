package main

import (
	"bufio"
	"container/heap"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25069/problems/K/
// Вы приехали на архипелаг Алгосы (наконец-то!).
//Здесь есть n достопримечательностей.
//Ваша лодка может высадить вас у одной из них, забрать у какой-то другой, возможно, той же
//самой достопримечательности и увезти на материк.
//
//Чтобы более тщательно спланировать свой маршрут, вы хотите
//узнать расстояния между каждой парой достопримечательностей Алгосов.
//Некоторые из них соединены мостами, по которым вы можете передвигаться в любую сторону.
//Всего мостов m.
//
//Есть вероятность, что карта архипелага устроена так, что нельзя
//добраться от какой-то одной достопримечательности до другой без использования лодки.
//
//Найдите кратчайшие расстояния между всеми парами достопримечательностей.
func main() {
	adjList := getAdjList()

	distances := make([][]int, len(adjList))
	for i, _ := range distances {
		distances[i] = getMinDistancesFromVertex(i, adjList)
	}

	printMatrix(distances)
}

func getMinDistancesFromVertex(vertex int, adjList [][][]int) []int {
	distances := make([]int, len(adjList))
	for i, _ := range distances {
		distances[i] = math.MaxInt
	}
	visited := make([]bool, len(adjList))

	distances[vertex] = 0

	pq := make(MinHeap, 0)

	heap.Init(&pq)
	heap.Push(&pq, []int{vertex, 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).([]int)
		visited[u[0]] = true

		for _, v := range adjList[u[0]] {
			if !visited[v[0]] && distances[v[0]] > distances[u[0]]+v[1] {
				distances[v[0]] = distances[u[0]] + v[1]
				heap.Push(&pq, []int{v[0], distances[v[0]]})
			}
		}
	}

	for i, isVisited := range visited {
		if !isVisited {
			distances[i] = -1
		}
	}

	return distances
}

type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getAdjList() [][][]int {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	adjList := make([][][]int, n)

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)

		if adjList[vertices[0]-1] == nil {
			adjList[vertices[0]-1] = make([][]int, 0)
		}
		adjList[vertices[0]-1] = append(adjList[vertices[0]-1], []int{vertices[1] - 1, vertices[2]})

		if adjList[vertices[1]-1] == nil {
			adjList[vertices[1]-1] = make([][]int, 0)
		}
		adjList[vertices[1]-1] = append(adjList[vertices[1]-1], []int{vertices[0] - 1, vertices[2]})
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
