package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/L/

// У Тимофея было n стажёров. Каждый стажёр хотел быть лучше своих предшественников, поэтому i-й стажёр
//делал столько коммитов, сколько делали два предыдущих стажёра в сумме.
//Два первых стажёра были менее инициативными —– они сделали по одному коммиту.
//Пусть Fi —– число коммитов, сделанных i-м стажёром (стажёры нумеруются с нуля).
//Тогда выполняется следующее: F0=F1=1.
//Для всех i≥2 выполнено Fi=Fi−1+Fi−2.
//Определите, сколько кода напишет следующий стажёр –— найдите последние k цифр числа Fn.

//Как найти k последних цифр
//
//Чтобы вычислить k последних цифр некоторого числа x, достаточно взять остаток от его деления на число 10k.
//Эта операция обозначается как x mod 10k. Узнайте, как записывается операция взятия
//остатка по модулю в вашем языке программирования.
//
//Также обратите внимание на возможное переполнение целочисленных типов, если в вашем языке такое случается.
func commitsCount(n int, k int) int {
	f := make([]int, n+2)
	f[0] = 1
	f[1] = 1
	modulus := int(math.Pow10(k))
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2]) % modulus
	}
	return f[n]
}

func main() {
	scanner := makeScanner()
	input := readArray(scanner)
	n := input[0]
	k := input[1]
	fmt.Println(commitsCount(n, k))
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
