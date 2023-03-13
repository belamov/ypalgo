package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// https://contest.yandex.ru/contest/23991/problems/B/

//Гоша написал программу, которая сравнивает строки исключительно по их хешам.
//Если хеш равен, то и строки равны.
//Тимофей увидел это безобразие и поручил вам сломать программу Гоши, чтобы остальным неповадно было.
//
//В этой задаче вам надо будет лишь найти две различные строки,
//которые для заданной хеш-функции будут давать одинаковое значение.
func main() {
	rand.Seed(time.Now().UnixNano())
	h := make(map[uint]string)
	i := 1
	for {
		fmt.Println(i)
		s := RandStringRunes(10)
		hash := polynomialHash(1000, 123987123, s)
		if v, ok := h[hash]; ok {
			fmt.Println(v)
			fmt.Println(s)
			return
		}
		h[hash] = s
		i++
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
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
