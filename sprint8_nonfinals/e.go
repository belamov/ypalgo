package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/E/
//У Риты была строка s, Гоша подарил ей на 8 марта ещё n других строк ti, 1≤ i≤ n.
//Теперь Рита думает, куда их лучше поставить.
//Один из вариантов —– расположить подаренные строки внутри имеющейся строки s, поставив
//строку ti сразу после символа строки s с номером ki (в частности,
//если ki=0, то строка вставляется в самое начало s).
//
//Помогите Рите и определите, какая строка получится после вставки в s всех подаренных Гошей строк.
func main() {
	scanner := makeScanner()
	s := readString(scanner)
	n := readInt(scanner)
	wordsToAdd := make(MinHeap, 0)
	newStrLength := len(s)
	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")
		t := listString[0]
		k, _ := strconv.Atoi(listString[1])
		wordsToAdd.Push(&word{pos: k, str: t})
		newStrLength += k
	}

	heap.Init(&wordsToAdd)

	accumulator := 0
	var result strings.Builder
	for i := 0; i < newStrLength; i++ {
		if wordsToAdd.Len() == 0 {
			result.WriteString(s[i-accumulator:])
			break
		}

		addingWord := heap.Pop(&wordsToAdd).(*word)

		for i != addingWord.pos+accumulator {
			result.WriteByte(s[i-accumulator])
			i++
		}
		result.WriteString(addingWord.str)
		accumulator += len(addingWord.str)
		i += len(addingWord.str) - 1
	}
	fmt.Print(result.String())
}

type MinHeap []*word

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].pos < h[j].pos
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*word))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type word struct {
	pos int
	str string
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
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
