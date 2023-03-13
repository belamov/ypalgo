package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/C/

//Алла не остановилась на достигнутом –— теперь она хочет научиться быстро вычислять
//хеши произвольных подстрок данной строки. Помогите ей!
//
//На вход поступают запросы на подсчёт хешей разных подстрок.
//Ответ на каждый запрос должен выполняться за O(1).
//Допустимо в начале работы программы сделать предподсчёт для дальнейшей работы со строкой.
func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	m := readInt(scanner)
	s := readString(scanner)

	prefixesHashes := polynomialHashAllPrefixes(a, m, s)
	powers := getPowers(a, m, len(s))

	n := readInt(scanner)

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		startEnd := readArray(scanner)
		writer.WriteString(hashSubstring(m, startEnd[0]-1, startEnd[1]-1, prefixesHashes, powers))
		writer.WriteString("\n")
	}
	writer.Flush()
}

func getPowers(a int, m int, n int) []int {
	powers := make([]int, n)
	powers[0] = 1
	for i := 1; i < n; i++ {
		powers[i] = (powers[i-1] * a) % m
	}
	return powers
}

func hashSubstring(m int, start int, end int, hashes []int, powers []int) string {
	if start == 0 {
		return strconv.Itoa(hashes[end])
	}

	hash := hashes[end] - hashes[start-1]*powers[end-start+1]
	hash = modLikePython(hash, m)

	return strconv.Itoa(hash)
}

func modLikePython(d, m int) int {
	var res = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func polynomialHashAllPrefixes(a int, m int, s string) []int {
	hash := make([]int, len(s))

	sum := int(s[0])
	hash[0] = sum % m

	for i := 1; i < len(s); i++ {
		hash[i] = (hash[i-1]*a + int(s[i])) % m
	}

	return hash
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
