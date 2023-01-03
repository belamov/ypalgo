package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/A/

// Алла получила задание, связанное с мониторингом работы различных серверов.
//Требуется понять, сколько времени обрабатываются определённые запросы на конкретных серверах.
//Эту информацию нужно хранить в матрице, где номер столбца соответствуют идентификатору
//запроса, а номер строки — идентификатору сервера. Алла перепутала строки и столбцы местами.
//С каждым бывает. Помогите ей исправить баг.
//
//Есть матрица размера m × n. Нужно написать функцию, которая её транспонирует.
//
//Транспонированная матрица получается из исходной заменой строк на столбцы.
func transpone(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := range result {
		result[i] = make([]int, len(matrix))
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	cols := readInt(scanner)

	if rows == 0 && cols == 0 {
		return
	}

	matrix := readMatrix(scanner, rows, cols)
	printMatrix(transpone(matrix))
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

func printMatrix(matrix [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			writer.WriteString(strconv.Itoa(matrix[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
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
