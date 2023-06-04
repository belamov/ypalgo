package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/25596/problems/D/
func main() {
	scanner := makeScanner()
	n := readInt(scanner)

	p, pp, v := 1, 1, 0

	for i := 2; i <= n; i++ {
		v = (p + pp) % 1000000007
		p, pp = pp, v
	}

	fmt.Print(v)
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
