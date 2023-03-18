package main

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/F/
// https://leetcode.com/problems/group-anagrams/

//Вася решил избавиться от проблем с произношением и стать певцом.
//Он обратился за помощью к логопеду.
//Тот посоветовал Васе выполнять упражнение, которое называется анаграммная группировка.
//В качестве подготовительного этапа нужно выбрать из множества строк анаграммы.
//
//Анаграммы –— это строки, которые получаются друг из друга перестановкой символов.
//Например, строки «SILENT» и «LISTEN» являются анаграммами.
//
//Помогите Васе найти анаграммы.
func main() {
	scanner := makeScanner()
	readString(scanner)
	arr := readArray(scanner)

	anagrams := groupAnagrams(arr)
	writer := bufio.NewWriter(os.Stdout)
	for _, anagramGroup := range anagrams {
		printArray(writer, anagramGroup)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func groupAnagrams(strs []string) [][]int {
	stringsMap := make(map[string][]int)
	for strPos, str := range strs {
		stringRunes := []rune(str)
		stringMap := make(map[rune]int)
		for _, char := range stringRunes {
			stringMap[char]++
		}
		stringKey, _ := json.Marshal(stringMap)
		stringsMap[string(stringKey)] = append(stringsMap[string(stringKey)], strPos)
	}

	result := make([][]int, 0, len(stringsMap))
	for _, stringPositions := range stringsMap {
		result = append(result, stringPositions)
	}

	sort.Slice(result[:], func(i, j int) bool {
		for x := range result[i] {
			if result[i][x] == result[j][x] {
				continue
			}
			return result[i][x] < result[j][x]
		}
		return false
	})

	return result
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func printArray(writer *bufio.Writer, arr []int) {
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
}
