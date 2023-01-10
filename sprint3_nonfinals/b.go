package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/23638/problems/B/

// На клавиатуре старых мобильных телефонов каждой цифре соответствовало несколько букв.
// Примерно так:
//
// 2:'abc',
// 3:'def',
// 4:'ghi',
// 5:'jkl',
// 6:'mno',
// 7:'pqrs',
// 8:'tuv',
// 9:'wxyz'
//
// Вам известно в каком порядке были нажаты кнопки телефона, без учета повторов.
// Напечатайте все комбинации букв, которые можно набрать такой последовательностью нажатий.
var abc = map[uint8][]string{
	50: {"a", "b", "c"},
	51: {"d", "e", "f"},
	52: {"g", "h", "i"},
	53: {"j", "k", "l"},
	54: {"m", "n", "o"},
	55: {"p", "q", "r", "s"},
	56: {"t", "u", "v"},
	57: {"w", "x", "y", "z"},
}

func generateCombinations(n int, numbers string, prefix string) {
	if n == len(numbers) {
		fmt.Print(prefix + " ")
		return
	}

	for _, ch := range abc[numbers[n]] {
		generateCombinations(n+1, numbers, prefix+ch)
	}
}

func main() {
	scanner := makeScanner()
	numbers := readLine(scanner)
	generateCombinations(0, numbers, "")
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
