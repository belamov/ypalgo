package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/24414/problems/A/
// https://contest.yandex.ru/contest/24414/run-report/84286536/ - последнее ОК решение

//Тимофей пишет свою поисковую систему.
//Имеется n документов, каждый из которых представляет собой текст из слов.
//По этим документам требуется построить поисковый индекс.
//На вход системе будут подаваться запросы. Запрос —– некоторый набор слов.
//По запросу надо вывести 5 самых релевантных документов.
//Релевантность документа оценивается следующим образом: для каждого уникального
//слова из запроса берётся число его вхождений в документ, полученные числа для всех слов из запроса суммируются.
//Итоговая сумма и является релевантностью документа.
//Чем больше сумма, тем больше документ подходит под запрос.
//Сортировка документов на выдаче производится по убыванию релевантности.
//Если релевантности документов совпадают —– то по возрастанию их
//порядковых номеров в базе (то есть во входных данных).

//Поисковый индекс будет иметь следующую структуру:
// слово -> отображение(номер документа -> количество нахождений слова в документе)
// это позволит быстро считать релевантность документов даже с большим повторением слов в документе

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// итоговая сложность O(N+M), где N - количество слов во всех документах, M - максимальное количество уникальных слов среди поисковых запросов
// из этого времени:
//  - индекс строится за О(N), где N - количество слов во всех документах
//  - каждый поисковый запрос выполняется за О(M), где M - количество уникальных слов в запросе
//
// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// программа будет потреблять O(K+L+P) памяти, где
//  K - количество уникальных слов во всех документах (для хранения поискового индекса)
//  L - количество документов (для хранения массива релевантности документов при каждом поисковом запросе)
//  P - максимальное количество уникальных слов среди поисковых запросов (для хранения истории слов, которые мы искали в индексе)
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
		queryResult := searchDocuments(searchIndex, query, n)
		printArray(writer, queryResult)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func searchDocuments(searchIndex map[string]map[int]int, query string, documentsCount int) []int {
	// заведем кеш для отброса повторений слов в документе
	wordsHistory := make(map[string]bool)

	// заведем мапу для построения релевантности документов. тут номер документа -> релевантность документа
	relevance := make([]int, documentsCount)

	// разделяем запрос на отдельные слова
	words := strings.Split(query, " ")

	// для каждого слова в запросе
	for _, word := range words {
		// если слово уже встречалось - пропускаем его
		if _, alreadySearched := wordsHistory[word]; alreadySearched {
			continue
		}

		// берем из поискового индекса все документы, в которых есть данное слово
		documentsContainingWord, wordExistsInDocuments := searchIndex[word]
		if !wordExistsInDocuments {
			continue
		}

		// для каждого найденного документа прибавляем к релевантности
		// документа количество повторений данного слова в этом документе
		for documentIndex, wordsCount := range documentsContainingWord {
			relevance[documentIndex] += wordsCount
		}

		//запоминаем слово из запроса в кеш, чтобы не учитывать повторяющие слова
		wordsHistory[word] = true
	}

	// нам нужны топ 5 релевантностей документов - для этого подойдет
	// приоритетная очередь на куче, у нее есть свойство, что она всегда хранит элементы упорядоченно.
	// если мы правильно определим метод PriorityQueue.Less, то сможем хранить элементы в убывающем порядке
	// причем нам нужны только 5 первых элементов в этой очереди, поэтому память для этой очереди будет константна
	pq := make(PriorityQueue, 0, 5)
	var minRelevance *DocumentRelevance
	for document, documentRelevance := range relevance {
		if documentRelevance == 0 {
			continue
		}
		// если мы уже выкидывали минимальный элемент из кучи,
		// то можем проверить заранее, что новый добавляемый элемент тоже будет выкинут - сэкономим немного
		// времени на перестроении дерева
		if minRelevance != nil && (minRelevance.relevance == documentRelevance && minRelevance.document < document || minRelevance.relevance > documentRelevance) {
			continue
		}
		// добавляем элемент в очередь. тут дерево перестроится, чтобы сохранить порядок элементов
		// сложность операции - log 5
		heap.Push(&pq, &DocumentRelevance{
			document:  document,
			relevance: documentRelevance,
		})

		// если мы добавили 6 элемент, то можем удалить наименьший из 6, чтобы у нас всегда было топ 5 релевантностей
		// в этой очереди. запомним этот удаленный элемент, чтобы не добавлять в очередь элементы меньше
		if pq.Len() > 5 {
			minRelevance = heap.Pop(&pq).(*DocumentRelevance)
		}
	}

	// тут в приоритетной очереди у нас остались топ 5 релевантностей документов, преобразуем их в массив с
	// номерами документов
	relevantDocuments := make([]int, 0, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*DocumentRelevance)
		relevantDocuments = append(relevantDocuments, item.document+1)
	}

	// поскольку при изъятии из очереди мы берем минимальный элемент, то в массиве они будут
	// располагаться от меньшего к большему. а нам нужны документы по уменьшению релевантности
	// поэтому нам нужен этот массив в обратном порядке
	for i, j := 0, len(relevantDocuments)-1; i < j; i, j = i+1, j-1 {
		relevantDocuments[i], relevantDocuments[j] = relevantDocuments[j], relevantDocuments[i]
	}

	return relevantDocuments
}

func makeSearchIndex(documents []string) map[string]map[int]int {
	//структура индекса = слово -> отображение(номер документа -> количество нахождений слова в документе)
	index := make(map[string]map[int]int)

	//переберем все документы
	for documentIndex, document := range documents {
		// в каждом документе возьмем все слова
		words := strings.Split(document, " ")

		// для каждого слова в документе
		for _, word := range words {
			// если в индексе уже есть это слово
			if wordDocumentsMap, ok := index[word]; ok {
				// если в индексе в этом документе уже есть слово
				if _, ok := wordDocumentsMap[documentIndex]; ok {
					// увеличим счетчик вхождений этого слова в документе
					index[word][documentIndex]++
				} else {
					// иначе инициируем счетчик этого слова в документе
					index[word][documentIndex] = 1
				}
			} else {
				// в индексе нет этого слова, инициируем запись
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

type DocumentRelevance struct {
	document  int // Индекс документа
	relevance int // Релевантность документа
	index     int // служебное поле для позиционирования элемента в дереве
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*DocumentRelevance

func (pq PriorityQueue) Len() int { return len(pq) }

//Less определяет релевантность документа согласно условию задачи
// если релевантность равна, то меньшим является тот документ, у которого больше номер
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
	item := x.(*DocumentRelevance)
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
