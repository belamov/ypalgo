package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/H/

// Тимофей записал два числа в двоичной системе счисления и попросил Гошу вывести их сумму, также в двоичной системе.
//Встроенную в язык программирования возможность сложения двоичных чисел применять нельзя.
//Помогите Гоше решить задачу.
//
//Решение должно работать за O(N), где N –— количество разрядов максимального числа на входе.
func getSum(firstNumber string, secondNumber string) string {
	firstLen := len(firstNumber) - 1
	secondLen := len(secondNumber) - 1

	i := 0

	result := ""
	carry := 0
	for i <= firstLen || i <= secondLen {
		firstDigit := 0
		secondDigit := 0

		firstDigitIndex := firstLen - i
		if firstDigitIndex >= 0 {
			firstDigit = int(firstNumber[firstDigitIndex] - 48)
		}

		secondDigitIndex := secondLen - i
		if secondDigitIndex >= 0 {
			secondDigit = int(secondNumber[secondDigitIndex] - 48)
		}

		sum := firstDigit + secondDigit + carry

		result = strconv.Itoa(sum%2) + result

		carry = sum / 2

		i++
	}

	if carry > 0 {
		result = "1" + result
	}

	return result
}

func main() {
	scanner := makeScanner()
	firstNumber := readLine(scanner)
	secondNumber := readLine(scanner)
	sum := getSum(firstNumber, secondNumber)
	fmt.Println(sum)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
