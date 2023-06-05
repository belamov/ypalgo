package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://contest.yandex.ru/contest/26131/problems/A/
//В некоторых языках предложения пишутся и читаются не слева направо, а справа налево.
//
//Вам под руку попался странный текст –— в нём обычный (слева направо) порядок букв в словах.
//А вот сами слова идут в противоположном направлении.
//Вам надо преобразовать текст так, чтобы слова в нём были написаны слева направо.
func main() {
	scanner := makeScanner()
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	for i := len(listString) - 1; i >= 0; i-- {
		fmt.Print(listString[i], " ")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
