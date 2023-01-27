package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/G/

// Рита решила оставить у себя одежду только трёх цветов: розового, жёлтого и малинового.
//После того как вещи других расцветок были убраны, Рита захотела отсортировать
//свой новый гардероб по цветам.
//Сначала должны идти вещи розового цвета, потом —– жёлтого, и в конце —– малинового.
//Помогите Рите справиться с этой задачей.
func countingSort(arr []int) []int {
	colors := []int{0, 0, 0}
	for _, color := range arr {
		colors[color]++
	}

	result := make([]int, len(arr))
	index := 0
	for i, colorCount := range colors {
		for j := 0; j < colorCount; j++ {
			result[index] = i
			index++
		}
	}

	return result
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	if n == 0 {
		return
	}
	arr := readArray(scanner)
	printArray(countingSort(arr))
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
	writer.Flush()
}
