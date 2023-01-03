package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/J/

//Основная теорема арифметики говорит: любое число раскладывается на произведение простых множителей
//единственным образом, с точностью до их перестановки. Например:
//
//Число 8 можно представить как 2 × 2 × 2.
//Число 50 –— как 2 × 5 × 5 (или 5 × 5 × 2, или 5 × 2 × 5). Три варианта отличаются лишь порядком следования множителей.
//Разложение числа на простые множители называется факторизацией числа.
//
//Напишите программу, которая производит факторизацию переданного числа.
func factorize(number int) []int {
	factors := make([]int, 0)
	originalNumber := number

	for i := 2; i*i <= originalNumber; i++ {
		for number%i == 0 {
			factors = append(factors, i)
			number /= i
		}
	}

	if number != 1 {
		factors = append(factors, number)
	}

	return factors
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	factorization := factorize(number)
	printArray(factorization)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
