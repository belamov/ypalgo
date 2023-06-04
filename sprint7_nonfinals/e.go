package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/F/
//Алла хочет купить дом на Алгосах.
//Для этого ей надо много наличных, которые она собирается получить в банкомате.
//Банкомат приличный, поэтому в нём есть бесконечно много банкнот каждого номинала.
//Всего номиналов k штук.
//Дом мечты Аллы стоит x франков.
//
//Найдите минимальное количество банкнот, которые в сумме дадут x франков.
//Если в набор входит несколько банкнот одинакового номинала, то учитывать надо их все.
//
//Например, если необходимо набрать 15 франков, а в банкомате
//купюры по 5 франков, то минимальное число купюр - 3
func main() {
	scanner := makeScanner()
	x := readInt(scanner)

	readInt(scanner)
	sizes := readArray(scanner)
	curr := make([]int, x+1)
	prev := make([]int, x+1)
	for i, _ := range curr {
		curr[i] = math.MaxInt
		prev[i] = math.MaxInt
	}
	curr[0] = 0
	prev[0] = 0
	for i := 0; i < len(sizes); i++ {
		for j := 0; j <= x; j++ {
			withCurrent := math.MaxInt
			if j >= sizes[i] && curr[j-sizes[i]] != math.MaxInt {
				withCurrent = 1 + curr[j-sizes[i]]
			}
			curr[j] = min(prev[j], withCurrent)
		}
		curr, prev = prev, curr
	}

	fmt.Print(prev[x])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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
