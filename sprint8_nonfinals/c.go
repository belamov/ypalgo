package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/C/
// Палиндром —– это строка, которая одинаково читается как слева направо, так и справа налево.
//
// Из данной строки s путём удаления и перестановки букв надо получить палиндром максимальной длины.
// Среди всех таких палиндромов надо получить лексикографически минимальный.
// Количество удалений и перестановок символов может быть любым.
func main() {
	scanner := makeScanner()
	s := readString(scanner)
	fmt.Print(getPalindrome(s))
}

func getPalindrome(s string) string {
	hash := [26]int{}
	for _, char := range s {
		hash[char-97]++
	}

	var result strings.Builder

	for char, count := range hash {
		for i := 0; i < count/2; i++ {
			result.WriteByte(byte(char + 97))
			hash[char] -= 2
		}
	}

	if result.Len() == 0 {
		for char, count := range hash {
			if count > 0 {
				return string(byte(char + 97))
			}
		}
	}

	halvesIntersect := false
	if len(s)%2 == 0 {
		for char, count := range hash {
			if count > 0 {
				result.WriteByte(byte(char) + 97)
				halvesIntersect = true
				break
			}
		}
	}

	firstHalf := result.String()
	var secondHalf strings.Builder

	lastIdx := len(firstHalf) - 1
	if halvesIntersect {
		lastIdx--
	}
	for i := lastIdx; i >= 0; i-- {
		secondHalf.WriteByte(firstHalf[i])
	}
	return firstHalf + secondHalf.String()
}
func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
