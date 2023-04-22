package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := makeScanner()

	n := readInt(scanner)
	fmt.Println(getNumberOfDifferentBST(n))
}

//https://contest.yandex.ru/contest/24809/problems/I/
//https://leetcode.com/problems/unique-binary-search-trees/description/
func getNumberOfDifferentBST(n int) int {
	catalanNumbers := make([]int, n+1)
	catalanNumbers[0] = 1
	catalanNumbers[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			catalanNumbers[i] += catalanNumbers[j] * catalanNumbers[i-j-1]
		}
	}
	return catalanNumbers[n]
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
