package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23638/problems/L/

// Алла захотела, чтобы у неё под окном были узкие клумбы с тюльпанам.
//На схеме земельного участка клумбы обозначаются просто горизонтальными отрезками,
//лежащими на одной прямой.
//Для ландшафтных работ было нанято n садовников.
//Каждый из них обрабатывал какой-то отрезок на схеме.
//Процесс был организован не очень хорошо, иногда один и тот же отрезок
//или его часть могли быть обработаны сразу несколькими садовниками.
//Таким образом, отрезки, обрабатываемые двумя разными садовниками, сливаются в один.
//Непрерывный обработанный отрезок затем станет клумбой. Нужно определить границы будущих клумб.
func getSegments(lines [][]int) [][]int {
	if len(lines) == 1 {
		return lines
	}

	left := getSegments(lines[0 : len(lines)/2])
	right := getSegments(lines[len(lines)/2:])

	result := make([][]int, 0, len(lines))

	l, r := 0, 0
	var merging []int
	for l < len(left) || r < len(right) {
		if l >= len(left) {
			merging = right[r]
			r++
		} else if r >= len(right) {
			merging = left[l]
			l++
		} else if left[l][0] <= right[r][0] {
			merging = left[l]
			l++
		} else {
			merging = right[r]
			r++
		}

		lastIndex := len(result) - 1
		if lastIndex < 0 {
			result = append(result, merging)
			continue
		}
		if merging[0] > result[lastIndex][1] {
			result = append(result, merging)
			continue
		}
		if result[lastIndex][0] >= merging[0] {
			result[lastIndex][0] = merging[0]
		}
		if result[lastIndex][1] <= merging[1] {
			result[lastIndex][1] = merging[1]
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	lines := make([][]int, n)
	for i := 0; i < n; i++ {
		lines[i] = readArray(scanner)
	}
	segments := getSegments(lines)
	for _, segment := range segments {
		printArray(segment)
	}
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.WriteString("\n")
	writer.Flush()
}
