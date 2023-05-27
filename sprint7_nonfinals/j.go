package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/J/
//Гоша решил отправиться в турне по островам Алгосского архипелага.
//Туристическая программа состоит из последовательного посещения n достопримечательностей.
//У i-й достопримечательности есть свой рейтинг ri.
//
//Впечатление от i-й достопримечательности равно её рейтингу ri.
//Гоша хочет, чтобы его впечатление от каждой новой посещённой достопримечательности
//было сильнее, чем от предыдущей.
//Ради этого он даже готов пропустить некоторые места в
//маршруте –— в случае, если они нарушают этот порядок плавного возрастания.
//
//Помогите Гоше и найдите наибольшую возрастающую подпоследовательность в массиве рейтингов ri.
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(scanner)

	dp := make([]int, n)
	dp[n-1] = 1
	maxLenth := 1
	for i := n - 2; i >= 0; i-- {
		nextElem := 0
		for j := i; j < n; j++ {
			if arr[i] < arr[j] && dp[j] > nextElem {
				nextElem = dp[j]
			}
		}
		dp[i] = 1 + nextElem
		if dp[i] > maxLenth {
			maxLenth = dp[i]
		}
	}

	fmt.Println(maxLenth)

	for i := 0; i < len(dp); i++ {
		if dp[i] == maxLenth {
			fmt.Print(i+1, " ")
			maxLenth--
		}
	}
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
