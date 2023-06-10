package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/26131/problems/K/
//Алла придумала новый способ сравнивать две строки: чтобы сравнить строки a и b, в них надо
//оставить только те буквы, которые в английском алфавите стоят на четных позициях.
//Затем полученные строки сравниваются по обычным правилам.
//Помогите Алле реализовать новое сравнение строк.
func main() {
	scanner := makeScanner()
	s := readString(scanner)
	t := readString(scanner)

	fmt.Print(compare(s, t))

}

func compare(s string, t string) int {
	sp, tp := 0, 0

	for sp < len(s) && tp < len(t) {
		for sp < len(s) && s[sp]%2 != 0 {
			sp++
		}
		for tp < len(t) && t[tp]%2 != 0 {
			tp++
		}
		if tp == len(t) && sp == len(s) {
			return 0
		}
		if tp == len(t) && sp < len(s) {
			return 1
		}
		if tp < len(t) && sp == len(s) {
			return -1
		}
		if s[sp] > t[tp] {
			return 1
		}
		if s[sp] < t[tp] {
			return -1
		}
		sp++
		tp++
	}
	return 0
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
