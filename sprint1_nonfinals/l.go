package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22449/problems/L/

//Васе очень нравятся задачи про строки, поэтому он придумал свою.
//Есть 2 строки s и t, состоящие только из строчных букв.
//Строка t получена перемешиванием букв строки s и добавлением 1 буквы в случайную позицию.
//Нужно найти добавленную букву.
func getExcessiveLetter(s1 string, t1 string) string {
	for i := 0; i < len(s1); i++ {
		t1 = strings.Replace(t1, string(s1[i]), "", 1)
	}
	return t1
}

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	fmt.Printf(getExcessiveLetter(s, t))
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
