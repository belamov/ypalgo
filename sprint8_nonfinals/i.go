package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/26131/problems/L/
// Будем говорить, что строка s является повтором длины k, если существует такая строка t, что s = t * k, где
// под умножением подразумевается конкатенация k экземпляров строки t один за другим.
//
// Например, строка abababab является повтором строки abab длины 2, а также повторением строки ab длины 4.
// Тогда имеет смысл говорить о наибольшем повторе.
// Строка является наибольшим повтором длины k, если она является повтором некоторой строки
// длины k и если не существует такой строки t, что s —– повтор t длины m > k.
// Например, строка aaaa является наибольшим повтором длины 4.
//
// Вам дана строка, которая является наибольшим повтором длины x. Найдите x.
//
// Заметим, что ответ всегда равен хотя бы единице, так как строка является повтором самой себя.
func main() {
	scanner := makeScanner()
	s := readString(scanner)

	p := prefixFunction(s)
	if p[len(s)-1] >= len(s)/2 {
		fmt.Print(len(s) / (len(s) - p[len(s)-1]))
	} else {
		fmt.Print(1)
	}
}

func prefixFunction(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		k := p[i-1]
		for k > 0 && s[k] != s[i] {
			k = p[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		p[i] = k
	}
	return p
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
