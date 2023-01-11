package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23638/problems/H/

// Вечером ребята решили поиграть в игру «Большое число».
// Даны числа. Нужно определить, какое самое большое число можно из них составить.
func getBiggestNumber(numbers []string) string {
	sort.Slice(numbers, func(i, j int) bool {
		ij, _ := strconv.Atoi(numbers[i] + numbers[j])
		ji, _ := strconv.Atoi(numbers[j] + numbers[i])
		return ij < ji
	})
	result := ""
	for i := len(numbers) - 1; i >= 0; i-- {
		result += numbers[i]
	}
	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	fmt.Println(getBiggestNumber(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
