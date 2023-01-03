package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/G/

// Вася реализовал функцию, которая переводит целое число из десятичной системы в двоичную.
//Но, кажется, она получилась не очень оптимальной.
//
//Попробуйте написать более эффективную программу.
//
//Не используйте встроенные средства языка по переводу чисел в бинарное представление.
func getBinaryNumber(n int) []int {
	if n == 0 {
		return []int{0}
	}

	result := make([]int, 0)

	for n > 0 {
		result = append(result, n%2)
		n = n / 2
	}

	reverse(result)

	return result
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	binaryNumber := getBinaryNumber(n)
	printArray(binaryNumber)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
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
	}
	writer.Flush()
}
