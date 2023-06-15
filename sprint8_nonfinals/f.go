package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/26131/problems/F/
//Дан набор слов, слова могут повторяться.
//Среди них надо найти самое частое слово.
//Если таких слов несколько, то выведите лексикографически наименьшее.
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	mostFrequentCount := -1
	mostFrequentWord := ""
	wordsHash := make(map[string]int)
	for i := 0; i < n; i++ {
		word := readString(scanner)
		wordsHash[word]++
		if i == 0 {
			mostFrequentCount = 1
			mostFrequentWord = word
			continue
		}
		wordOccurrencesCount := wordsHash[word]
		if wordOccurrencesCount > mostFrequentCount || (wordOccurrencesCount == mostFrequentCount && word < mostFrequentWord) {
			mostFrequentCount = wordOccurrencesCount
			mostFrequentWord = word
		}
	}
	fmt.Print(mostFrequentWord)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
