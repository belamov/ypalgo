package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	documents := make([]string, n)
	for i := 0; i < n; i++ {
		documents[i] = readString(scanner)
	}

	searchIndex := makeSearchIndex(documents)

	m := readInt(scanner)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < m; i++ {
		query := readString(scanner)
		queryResult := searchDocuments(searchIndex, query)
		printArray(writer, queryResult)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func searchDocuments(searchIndex map[string]map[int]int, query string) []int {
	wordsHistory := make(map[string]bool)
	relevance := make(map[int]int)
	words := strings.Split(query, " ")
	for _, word := range words {
		if _, alreadySearched := wordsHistory[word]; alreadySearched {
			continue
		}

		documentsContainingWord, wordExistsInDocuments := searchIndex[word]
		if !wordExistsInDocuments {
			continue
		}

		for documentIndex, wordsCount := range documentsContainingWord {
			relevance[documentIndex] += wordsCount
		}

		wordsHistory[word] = true
	}

	pq := make(PriorityQueue, 0, 5)
	var minElement *Item
	for document, documentRelevance := range relevance {
		if minElement != nil && (minElement.relevance == documentRelevance && minElement.document < document || minElement.relevance > documentRelevance) {
			continue
		}
		heap.Push(&pq, &Item{
			document:  document,
			relevance: documentRelevance,
		})
		if pq.Len() > 5 {
			minElement = heap.Pop(&pq).(*Item)
		}
	}

	relevantDocuments := make([]int, 0, 5)
	for i := 0; pq.Len() > 0; i++ {
		item := heap.Pop(&pq).(*Item)
		relevantDocuments = append(relevantDocuments, item.document+1)
	}

	for i, j := 0, len(relevantDocuments)-1; i < j; i, j = i+1, j-1 {
		relevantDocuments[i], relevantDocuments[j] = relevantDocuments[j], relevantDocuments[i]
	}

	return relevantDocuments
}

func makeSearchIndex(documents []string) map[string]map[int]int {
	index := make(map[string]map[int]int)
	for documentIndex, document := range documents {
		words := strings.Split(document, " ")
		for _, word := range words {
			if wordDocumentsMap, ok := index[word]; ok {
				if _, ok := wordDocumentsMap[documentIndex]; ok {
					index[word][documentIndex]++
				} else {
					index[word][documentIndex] = 1
				}
			} else {
				index[word] = map[int]int{documentIndex: 1}
			}
		}
	}
	return index
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
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

func printArray(writer *bufio.Writer, arr []int) {
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
}

// An Item is something we manage in a priority queue.
type Item struct {
	document  int // The value of the item; arbitrary.
	relevance int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].relevance == pq[j].relevance {
		return pq[i].document > pq[j].document
	}
	return pq[i].relevance < pq[j].relevance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
