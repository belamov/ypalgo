package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/A/
//Рита хочет попробовать поиграть на бирже.
//Но для начала она решила потренироваться на исторических данных.
//
//Даны стоимости акций в каждый из n дней.
//В течение дня цена акции не меняется.
//Акции можно покупать и продавать, но только по одной штуке в день.
//В один день нельзя совершать более одной операции (покупки или продажи).
//Также на руках не может быть более одной акции в каждый момент времени.
//
//Помогите Рите выяснить, какую максимальную прибыль она могла бы получить.
func main() {
	scanner := makeScanner()
	readInt(scanner)
	prices := readArray(scanner)

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	fmt.Print(profit)
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
