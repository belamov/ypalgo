package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23638/problems/L/

// Вася решил накопить денег на два одинаковых велосипеда — себе и сестре.
// У Васи есть копилка, в которую каждый день он может добавлять
// деньги (если, конечно, у него есть такая финансовая возможность).
// В процессе накопления Вася не вынимает деньги из копилки.
//
// У вас есть информация о росте Васиных накоплений — сколько у Васи
// в копилке было денег в каждый из дней.
//
// Ваша задача — по заданной стоимости велосипеда определить
//
// - первый день, в которой Вася смог бы купить один велосипед,
// - и первый день, в который Вася смог бы купить два велосипеда.
//
//Подсказка: решение должно работать за O(log n).
func findDayWithEnoughMoney(amounts []int, amount int, left int, right int) int {
	if right <= left {
		return -1
	}
	mid := (left + right) / 2
	if amounts[mid-1] < amount && amounts[mid] >= amount {
		return mid
	}

	if amounts[mid] >= amount {
		return findDayWithEnoughMoney(amounts, amount, left, mid)
	}

	return findDayWithEnoughMoney(amounts, amount, mid+1, right)
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	amounts := readArray(scanner)
	amount := readInt(scanner)
	result := make([]int, 2)
	result[0] = findDayWithEnoughMoney(amounts, amount, 0, n+1)
	if result[0] >= 0 {
		result[1] = findDayWithEnoughMoney(amounts, amount*2, result[0]-1, n+1)
	} else {
		result[1] = -1
	}
	printArray(result)
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
	arr := make([]int, len(listString)+2)
	arr[0] = 0
	for i := 0; i < len(listString); i++ {
		arr[i+1], _ = strconv.Atoi(listString[i])
	}
	arr[len(arr)-1] = arr[len(arr)-2]
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
