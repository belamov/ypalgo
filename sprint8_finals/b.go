package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://contest.yandex.ru/contest/26133/problems/B/

// https://contest.yandex.ru/contest/26133/run-report/88188394/ - ссылка на последнее ОК решение

//Вася готовится к экзамену по алгоритмам и на всякий случай пишет шпаргалки.
//
//Чтобы уместить на них как можно больше информации, он не разделяет слова пробелами.
//В итоге получается одна очень длинная строка.
//Чтобы на самом экзамене из-за нервов не запутаться в прочитанном, он просит вас
//написать программу, которая по этой длинной строке и набору допустимых слов
//определит, можно ли разбить текст на отдельные слова из набора.
//
//Более формально: дан текст T и набор строк s1, ... ,sn.
//Надо определить, представим ли T как sk1sk2...skr, где где ki — индексы строк.
//Индексы могут повторяться.
//Строка si может встречаться в разбиении текста T произвольное число раз.
//Можно использовать не все строки для разбиения.
//Строки могут идти в любом порядке.

//Данную задачу можно решить при помощи динамического программирования:

// - Что будет храниться в dp?
//   dp[i] - можем ли мы разбить подстроку от 0 до i символа на слова из словаря
//
// - Каким будет базовый случай для задачи?
//   dp[0] = true, поскольку пустая строка может быть представлена любым словарем (не используя никакие строки из словаря)
//
// - Каким будет переход динамики?
//   мы пойдем по возможным разделителям строки, то есть по таким значениям dp[i], которые равны true
//   для каждой такой позиции мы будем искать возможные слова в оставшейся строке от i до конца строки
//   для каждого найденного слова мы будем помечать dp[i+len(word)] = true
//   поиск подходящих слов будем производить с помощью бора, составленного из словаря
//
// - Каким будет порядок вычисления данных в массиве dp?
//   от dp[0] до dp[len(s)]. причем мы можем пропускать значения dp[i] = false
//
// - Где будет располагаться ответ на исходный вопрос?
//   в dp[len(s)]

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// O(L+|s|^2), где L - суммарная длина слов в словаре, |s| - длина строки
// бор строится за O(L), где L - суммарная длина слов в словаре
// каждый суффикс ищется за O(|suf|), где |suf| - длина искомого суффикса
// на каждой итерации вычисления dp мы в худшем случае будем искать все суффиксы строки
// каждый суффикс будет как минимум меньше предыдущего на 1 символ
// значит для поиска всех суффиксов в боре нам понадобится O(|s| + |s|-1 + |s|-2 + ... + 1) = O(|s|^2), где |s| - длина строки
// итого итоговая сложность получается O(L+|s|^2)

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// O(|s|), где |s| - длина строки
// мы будем хранить бор
// в нем будет |s| узлов, в каждом узле может быть не более 26 смежных узлов (буквы только маленькие латинские)
// итого будет O(|s|*26)~O(|s|)
func main() {
	scanner := makeScanner()
	s := readString(scanner)

	n := readInt(scanner)
	trieRoot := &TrieNode{adjNodes: make(map[uint8]*TrieNode)}
	for i := 0; i < n; i++ {
		addString(trieRoot, readString(scanner))
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(dp); i++ {
		if dp[i] == false {
			continue
		}

		availableBreaks := findMatchingWords(trieRoot, s[i:])
		for _, availableBreak := range availableBreaks {
			dp[i+availableBreak] = true
			if i+availableBreak == len(s) {
				fmt.Print("YES")
				return
			}
		}
	}

	if dp[len(s)] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

type TrieNode struct {
	isTerminal bool
	adjNodes   map[uint8]*TrieNode
	strLen     int
}

func addString(root *TrieNode, str string) {
	cur := root
	for i := 0; i < len(str); i++ {
		nextNode, ok := cur.adjNodes[str[i]]

		if !ok {
			nextNode = &TrieNode{adjNodes: make(map[uint8]*TrieNode)}
			cur.adjNodes[str[i]] = nextNode
		}

		cur = nextNode
	}
	cur.isTerminal = true
	cur.strLen = len(str)
}

func findMatchingWords(root *TrieNode, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	results := make([]int, 0)

	charFound := true
	var cur *TrieNode
	offset := 0
	cur = root
	for charFound && offset < len(pattern) {
		charFound = false
		char := pattern[offset]
		next, ok := cur.adjNodes[char]
		if ok {
			cur = next
			charFound = true
			if next.isTerminal {
				results = append(results, next.strLen)
			}
		}
		offset++
	}

	return results
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
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
