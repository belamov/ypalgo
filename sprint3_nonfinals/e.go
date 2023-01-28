package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/E/

// Тимофей решил купить несколько домов на знаменитом среди разработчиков Алгосском архипелаге.
//Он нашёл n объявлений о продаже, где указана стоимость каждого дома в алгосских франках.
//А у Тимофея есть k франков.
//Помогите ему определить, какое наибольшее количество домов на Алгосах он сможет приобрести за эти деньги.
func getAvailableCount(prices []int, money int) int {
	sort.Ints(prices)
	totalSpend := 0
	count := 0
	for i := 0; i < len(prices) && totalSpend+prices[i] <= money; i++ {
		count++
		totalSpend += prices[i]
	}
	return count
}

func main() {
	scanner := makeScanner()
	nk := readArray(scanner)
	k := nk[1]
	prices := readArray(scanner)
	fmt.Println(getAvailableCount(prices, k))
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
