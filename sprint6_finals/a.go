package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/25070/problems/A/

// https://contest.yandex.ru/contest/25070/run-report/87460184/ - ссылка на последнее ОК решение

// Тимофей решил соединить все компьютеры в своей компании в единую сеть.
//Для этого он придумал построить минимальное остовное дерево, чтобы эффективнее использовать ресурсы.
//
//Но от начальства пришла новость о том, что выделенный на сеть бюджет
//оказался очень большим и его срочно надо израсходовать.
//Поэтому Тимофея теперь интересуют не минимальные, а максимальные остовные деревья.
//
//Он поручил вам найти вес такого максимального остовного дерева в неориентированном графе, который задаёт схему офиса.

//Для построения максимального остовного дерева воспользуемся алгоритмом прима

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
//O(E*logЕ), где E - количество ребер
//в приоритетной очереди хранятся ребра, поэтому операции вставки/удаления в нее будут O(logE)
//в очередь в худшем случае мы добавим все ребра - на это уйдет O(E*logE)
//из очереди мы достанем максимум V ребер - на это уйдет O(V*logE)
//итого нам понадобится O(E*logE) + O(V*logE) = O((E+V)*logE) ~ O(E*logE)
//

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
//O(V+E), где V - количество вершин, E - количество ребер
// Для хранения списка смежности графа нам понадобится O(V+E)
//для хранения кеша еще не добавленных вершин нам потребуется O(V) памяти
//для хранения приоритетной очереди из ребер нам понадобится O(E) памяти
//получается всего мы потратим O(V+E) + O(V) + O(E) = O(2*V + 2*E) ~ O(V+E)
func main() {
	adjList := getAdjList()

	maximumSpanningTreeWeight, err := getMaximumSpanningTreeWeight(adjList)

	if err != nil {
		fmt.Print("Oops! I did it again")
	} else {
		fmt.Print(maximumSpanningTreeWeight)
	}
}

func getMaximumSpanningTreeWeight(adjList [][]Edge) (int, error) {
	notAddedVertices := make(map[int]bool) // Множество вершины, ещё не добавленных в остов.
	for v := range adjList {
		notAddedVertices[v] = true
	}

	edges := make(MaxHeap, 0) // Массив рёбер, исходящих из остовного дерева.
	totalWeight := 0

	//Берём первую попавшуюся вершину.
	heap.Push(&edges, Edge{0, 0, 0})

	for len(notAddedVertices) > 0 && edges.Len() > 0 {
		e := heap.Pop(&edges).(Edge) //извлекаем максимальное ребро
		_, endNotAdded := notAddedVertices[e.end]
		if endNotAdded {
			totalWeight += e.weight
			delete(notAddedVertices, e.end)

			// Добавляем все рёбра, которые инцидентны v, но их конец ещё не в остове.
			for _, edge := range adjList[e.end] {
				_, endOfEdgeNotAdded := notAddedVertices[edge.end]
				if endOfEdgeNotAdded {
					heap.Push(&edges, edge)
				}
			}
		}
	}

	if len(notAddedVertices) > 0 {
		return 0, errors.New("граф несвязный")
	}

	return totalWeight, nil
}

type Edge struct {
	start  int
	end    int
	weight int
}

type MaxHeap []Edge

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].weight > h[j].weight
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getAdjList() [][]Edge {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	adjList := make([][]Edge, n)

	for i := 0; i < m; i++ {
		vertices := readArray(scanner)

		adjList[vertices[0]-1] = append(adjList[vertices[0]-1], Edge{vertices[0] - 1, vertices[1] - 1, vertices[2]})
		adjList[vertices[1]-1] = append(adjList[vertices[1]-1], Edge{vertices[1] - 1, vertices[0] - 1, vertices[2]})
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
