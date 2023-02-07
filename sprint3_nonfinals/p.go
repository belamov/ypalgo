package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/P/

// После того, как Гоша узнал про сортировку слиянием и быструю сортировку, он решил придумать свой
//метод сортировки, который предполагал бы разделение данных на части.
//Назвал он свою сортировку Частичной.
//Этим методом можно отсортировать n уникальных чисел a1, a2, … , an, находящихся в диапазоне от 0 до n - 1.
//Алгоритм сортировки состоит из трёх шагов:
//
//Разбить исходную последовательность на k блоков B1, …, Bk.
//Блоки могут иметь разные размеры.
//Если размер блока i равен si, то B1 ={ a1, …, as1 }, B2 = { as1 + 1, … , as1 + s2 } и так далее.

//Отсортировать каждый из блоков.

//Объединить блоки — записать сначала отсортированный блок B1, потом B2, … , Bk

//Частичная сортировка лучше обычной в том случае, если в первом пункте у нас получится
//разбить последовательность на большое число блоков.
//Однако далеко не каждое такое разбиение подходит.
//Определите максимальное число блоков, на которое можно разбить
//исходную последовательность, чтобы сортировка отработала корректно.
func getMaxSegmentsCount(arr []int) int {
	segmentsCount := 0
	elementsSum := 0
	indexSum := 0
	for i := 0; i < len(arr); i++ {
		elementsSum += arr[i]
		indexSum += i
		if elementsSum == indexSum {
			segmentsCount++
		}
	}
	return segmentsCount
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	fmt.Println(getMaxSegmentsCount(arr))
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
