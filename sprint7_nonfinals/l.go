package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/L/
//Лепреконы в данной задаче появились по соображениям общей морали, так как грабить банки — нехорошо.
//
//Вам удалось заключить неплохую сделку с лепреконами, поэтому они пустили вас в своё хранилище золотых слитков.
//Все слитки имеют единую пробу, то есть стоимость 1 грамма золота в двух разных слитках одинакова.
//В хранилище есть n слитков, вес i-го слитка равен wi кг.
//У вас есть рюкзак, вместимость которого M килограмм.
//
//Выясните максимальную суммарную массу золотых слитков, которую вы сможете унести.
func main() {
	scanner := makeScanner()
	nm := readArray(scanner)
	n := nm[0]
	m := nm[1]

	goldBars := readArray(scanner)

	newCalculation := make([]int, m+1)
	previousCalculation := make([]int, m+1)

	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			bestValueOfBagSizeJWithoutIBar := 0
			if i > 0 {
				bestValueOfBagSizeJWithoutIBar = previousCalculation[j]
			}

			bestValueOfBagSizeWithJBar := 0
			if j-goldBars[i] >= 0 {
				bestValueOfBagSizeWithJBar = goldBars[i]
				if i > 0 {
					bestValueOfBagSizeWithJBar = goldBars[i] + previousCalculation[j-goldBars[i]]
				}
			}

			newCalculation[j] = max(bestValueOfBagSizeJWithoutIBar, bestValueOfBagSizeWithJBar)

		}
		newCalculation, previousCalculation = previousCalculation, newCalculation

	}

	fmt.Print(previousCalculation[m])
}

func max(a, b int) int {
	if a > b {
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
