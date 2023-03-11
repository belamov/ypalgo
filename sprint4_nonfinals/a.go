package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/23991/problems/A/

//Алле очень понравился алгоритм вычисления полиномиального хеша.
//Помогите ей написать функцию, вычисляющую хеш строки s.
//В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.
func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	m := readInt(scanner)
	s := readString(scanner)

	fmt.Println(polynomialHash(a, m, s))
}

func polynomialHash(a uint, m uint, s string) uint {
	if len(s) == 0 {
		return 0
	}

	sum := uint(s[0])
	for i := 1; i < len(s); i++ {
		sum = (sum*a + uint(s[i])) % m
	}
	return sum % m
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) uint {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return uint(res)
}
