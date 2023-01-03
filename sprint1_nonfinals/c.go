package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22449/problems/C/

// Дана матрица. Нужно написать функцию, которая для элемента возвращает всех его соседей.
// Соседним считается элемент, находящийся от текущего на одну ячейку влево, вправо, вверх или вниз.
// Диагональные элементы соседними не считаются.
//
// Например, в матрице A соседними элементами для (0, 0) будут 2 и 0. А для (2, 1) –— 1, 2, 7, 7.
func getNeighbours(matrix [][]int, row int, col int) []int {
	result := make([]int, 0)
	if row-1 >= 0 {
		result = append(result, matrix[row-1][col])
	}
	if row+1 < len(matrix) {
		result = append(result, matrix[row+1][col])
	}
	if col-1 >= 0 {
		result = append(result, matrix[row][col-1])
	}
	if col+1 < len(matrix[row]) {
		result = append(result, matrix[row][col+1])
	}
	sort.Ints(result)
	return result
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	cols := readInt(scanner)
	matrix := readMatrix(scanner, rows, cols)
	rowId := readInt(scanner)
	colId := readInt(scanner)
	neighbours := getNeighbours(matrix, rowId, colId)
	for _, elem := range neighbours {
		fmt.Print(elem)
		fmt.Print(" ")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readArray(scanner)
	}
	return matrix
}
