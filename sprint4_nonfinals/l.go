package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/L/

//Дана длинная строка, состоящая из маленьких латинских букв.
//Нужно найти все её подстроки длины n, которые встречаются хотя бы k раз.
func main() {
	scanner := makeScanner()

	nk := readArray(scanner)
	s := readString(scanner)

	matchingSubstringsIndecies := getMatchingSubstrings(s, nk[0], nk[1])

	printArray(matchingSubstringsIndecies)
}

func getMatchingSubstrings(s string, n int, k int) []int {
	result := make([]int, 0)
	history := make(map[uint][][]int)
	m := uint(1e11 + 7)
	a := uint(31)
	pow := getPower(n, a, m)
	hash := polynomialHash(a, m, s[0:n])
	history[hash] = [][]int{{0, 1}}
	for i := 1; i+n-1 < len(s); i++ {
		hash = (hash + m - getCharCode(s[i-1])*pow%m) % m //отнимем хеш первого символа умноженный на базу в степени N
		hash = (hash*a%m + getCharCode(s[i+n-1])) % m     //прибавим хеш следующего символа (база будет в нулевой степени)

		if v, ok := history[hash]; ok {
			//обработаем коллизии
			for j, substringInfo := range v {
				if s[substringInfo[0]:substringInfo[0]+n] == s[i:i+n] {
					history[hash][j][1]++
					if history[hash][j][1] == k {
						result = append(result, history[hash][j][0])
					}
					break
				}
			}
		} else {
			history[hash] = [][]int{{i, 1}}
		}
	}

	return result
}

func getPower(n int, a, m uint) uint {
	power := uint(1)
	for i := 1; i < n; i++ {
		power = (power * a) % m
	}
	return power
}

func getCharCode(ch uint8) uint {
	return uint(ch) - uint('a') + 1
}

func polynomialHash(a uint, m uint, s string) uint {
	hash := uint(0)
	for i := 0; i < len(s); i++ {
		hash = (hash*a%m + getCharCode(s[i])) % m
	}
	return hash
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]))
		fmt.Print(" ")
	}
}
