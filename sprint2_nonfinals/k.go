package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22779/problems/K/

// У Тимофея было n стажёров. Каждый стажёр хотел быть лучше своих предшественников, поэтому i-й стажёр
//делал столько коммитов, сколько делали два предыдущих стажёра в сумме.
//Два первых стажёра были менее инициативными —– они сделали по одному коммиту.
//Пусть Fi —– число коммитов, сделанных i-м стажёром (стажёры нумеруются с нуля).
//Тогда выполняется следующее: F0=F1=1.
//Для всех i≥2 выполнено Fi=Fi−1+Fi−2.
//Определите, сколько кода напишет следующий стажёр –— найдите Fn.
//Решение должно быть реализовано рекурсивно.
func commitsCount(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return commitsCount(n-1) + commitsCount(n-2)
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	fmt.Println(commitsCount(n))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
