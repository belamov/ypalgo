package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23638/problems/J/

// Чтобы выбрать самый лучший алгоритм для решения задачи, Гоша продолжил изучать разные сортировки.
// На очереди сортировка пузырьком — https://ru.wikipedia.org/wiki/Сортировка_пузырьком
//
// Её алгоритм следующий (сортируем по неубыванию):
//
// На каждой итерации проходим по массиву, поочередно сравнивая пары соседних элементов.
// Если элемент на позиции i больше элемента на позиции i + 1, меняем их местами.
// После первой итерации самый большой элемент всплывёт в конце массива.
// Проходим по массиву, выполняя указанные действия до тех пор, пока на очередной итерации не окажется,
// что обмены больше не нужны, то есть массив уже отсортирован.
// После не более чем n – 1 итераций выполнение алгоритма заканчивается, так как
// на каждой итерации хотя бы один элемент оказывается на правильной позиции.
//
// Помогите Гоше написать код алгоритма.
func bubbleSort(arr []int) {
	n := len(arr) - 1
	for j := 0; j < n; j++ {
		isSorted := true
		for i := 0; i < n-j; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		printArray(arr)
	}
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	bubbleSort(arr)
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
	isSorted := true
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
		if isSorted && i > 0 {
			if arr[i] < arr[i-1] {
				isSorted = false
			}
		}
	}
	if isSorted {
		printArray(arr)
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
