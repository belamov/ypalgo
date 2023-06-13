package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"unicode"
)

//https://contest.yandex.ru/contest/26131/problems/J/
//В некоторых IDE поддерживается навигация по файлам через их сокращённые названия.
//Если в языке принято называть классы CamelCase'ом (как в Java, например), то по заглавным
//буквам названия можно быстро найти нужный класс.
//Например, если название класса «MyFavouriteConfigurableScannerFactory», то его можно
//найти по строке «MFCSF». Но если в проекте есть класс «theMultiFunctionalCommaSeparatedFile», то он
//тоже будет подходить под этот паттерн, и при поиске надо будет выбрать между этими двумя вариантами.
//
//Вам дан набор строк в CamelCase.
//Далее будут поступать запросы в виде строк-паттернов из прописных букв английского алфавита.
//Вам надо находить такие строки среди исходных, которые удовлетворяют
//заданному шаблону, и выводить их в лексикографическом порядке.
//
//Также в паттерне может быть только несколько первых
//заглавных букв. Например, если бы в указанном выше примере был бы паттерн «MFCS», то существующие
//две строки походили бы под него, а также подходил бы, например, «MamaFicusCodingSouthWestNorth».
//А вот «MamaCodingSouthWestNorth» –— уже нет.
func main() {
	scanner := makeScanner()

	n := readInt(scanner)
	classes := make([]string, n)
	for i := 0; i < n; i++ {
		classes[i] = readString(scanner)
	}

	trie := makeTrie(classes)

	m := readInt(scanner)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < m; i++ {
		matchingClasses := find(trie, readString(scanner))
		sort.Strings(matchingClasses)
		printArray(matchingClasses, writer)
	}
	writer.Flush()
}

type TrieNode struct {
	isTerminal     bool
	value          rune
	terminalValues []string
	adjNodes       []*TrieNode
}

func makeTrie(classes []string) *TrieNode {
	root := &TrieNode{
		isTerminal:     false,
		value:          0,
		terminalValues: nil,
		adjNodes:       nil,
	}
	for _, class := range classes {
		addString(root, class)
	}
	return root
}

func addString(root *TrieNode, str string) {
	cur := root
	for i := 0; i < len(str); i++ {
		if !unicode.IsUpper(rune(str[i])) {
			continue
		}

		var nextNode *TrieNode
		for _, adjNode := range cur.adjNodes {
			if adjNode.value == rune(str[i]) {
				nextNode = adjNode
			}
		}

		if nextNode == nil {
			nextNode = &TrieNode{
				isTerminal:     false,
				value:          rune(str[i]),
				terminalValues: make([]string, 0),
				adjNodes:       nil,
			}
			cur.adjNodes = append(cur.adjNodes, nextNode)
		}

		cur = nextNode
	}
	cur.isTerminal = true
	cur.terminalValues = append(cur.terminalValues, str)
}

func find(root *TrieNode, pattern string) []string {
	results := make([]string, 0)

	if len(pattern) == 0 {
		for _, value := range root.terminalValues {
			results = append(results, value)
		}
		return getTerminalValues(root, results)
	}

	charFound := true
	var cur *TrieNode
	offset := 0
	cur = root
	for charFound && offset < len(pattern) {
		charFound = false
		char := rune(pattern[offset])
		for _, node := range cur.adjNodes {
			if node.value == char {
				cur = node
				charFound = true
				break
			}
		}
		offset++
	}

	if charFound || len(pattern) == 0 {
		if cur.isTerminal {
			for _, value := range cur.terminalValues {
				results = append(results, value)
			}
		}
		results = getTerminalValues(cur, results)
	}

	return results
}

func getTerminalValues(cur *TrieNode, results []string) []string {
	var stack []*TrieNode
	for _, node := range cur.adjNodes {
		stack = append(stack, node)
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]     // Получаем из стека очередную вершину.
		stack = stack[:len(stack)-1] // Удаляем вершину из стека.

		if v.isTerminal {
			for _, value := range v.terminalValues {
				results = append(results, value)
			}
		}
		for _, node := range v.adjNodes {
			stack = append(stack, node)

		}
	}
	return results
}

func printArray(arr []string, writer *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		writer.WriteString(arr[i])
		writer.WriteString("\n")
	}
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
