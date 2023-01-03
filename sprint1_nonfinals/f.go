package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://contest.yandex.ru/contest/22449/problems/F/

// Помогите Васе понять, будет ли фраза палиндромом.
//Учитываются только буквы и цифры, заглавные и строчные буквы считаются одинаковыми.
//
//Решение должно работать за O(N), где N — длина строки на входе.
func isPalindrome(line string) bool {
	start := 0
	end := len(line) - 1

	for start < end {
		if !unicode.IsLetter(rune(line[start])) {
			start++
			continue
		}
		if !unicode.IsLetter(rune(line[end])) {
			end--
			continue
		}

		if !strings.EqualFold(string(line[start]), string(line[end])) {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	scanner := makeScanner()
	line := readLine(scanner)
	if isPalindrome(line) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
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
