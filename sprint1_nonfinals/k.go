package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22449/problems/K/

//Вася просил Аллу помочь решить задачу. На этот раз по информатике.
//
//Для неотрицательного целого числа X списочная форма –— это массив его цифр слева направо.
//К примеру, для 1231 списочная форма будет [1,2,3,1]. На вход подается количество цифр числа Х, списочная форма
//неотрицательного числа Х и неотрицательное число K. Число К не превосходят 10000. Длина числа Х не превосходит 1000.
//
//Нужно вернуть списочную форму числа X + K.
func getSum(bigNumber []int, smallNumber int) []int {
	result := make([]int, 0)

	carry := 0
	firstDigit := 0
	secondDigit := 0

	secondDigits := make([]int, 0)
	for smallNumber > 0 {
		digit := smallNumber % 10
		secondDigits = append([]int{digit}, secondDigits...)
		smallNumber = smallNumber / 10
	}
	for i := 0; i < len(bigNumber) || i < len(secondDigits); i++ {
		if i < len(bigNumber) {
			firstDigit = bigNumber[len(bigNumber)-1-i]
		} else {
			firstDigit = 0
		}

		if i < len(secondDigits) {
			secondDigit = secondDigits[len(secondDigits)-1-i]
		} else {
			secondDigit = 0
		}

		sum := firstDigit + secondDigit + carry
		carry = sum / 10
		result = append([]int{sum % 10}, result...)
	}

	if carry > 0 {
		result = append([]int{carry}, result...)
	}
	return result
}

func main() {
	scanner := makeScanner()
	readLine(scanner)
	bigNumer := readArray(scanner)
	smallNumber := readInt(scanner)
	printArray(getSum(bigNumer, smallNumber))
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
