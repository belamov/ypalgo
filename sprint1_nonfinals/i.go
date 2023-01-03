package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/I/

// Напишите программу, которая определяет, будет ли положительное целое число степенью четвёрки.
//
//Подсказка: степенью четвёрки будут все числа вида 4^n, где n – целое неотрицательное число.
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	isOdd := (n & (n - 1)) == 0
	is1BitInEvenPosition := (n & 0xAAAAAAAA) == 0

	return isOdd && is1BitInEvenPosition
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	if isPowerOfFour(number) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
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
