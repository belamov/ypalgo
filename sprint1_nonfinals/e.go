package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22449/problems/E/

// Чтобы подготовиться к семинару, Гоше надо прочитать статью по эффективному менеджменту.
//Так как Гоша хочет спланировать день заранее, ему необходимо оценить сложность статьи.
//
//Он придумал такой метод оценки: берётся случайное предложение из текста и в нём ищется
//самое длинное слово. Его длина и будет условной сложностью статьи.
//
//Помогите Гоше справиться с этой задачей.
func getLongestWord(line string) string {
	longestWordLength := 0
	longestWordBeginIndex := 0

	currentWordBeginIndex := 0
	currentWordLength := 0
	for i, symbol := range line {
		if i == len(line)-1 {
			if currentWordLength+1 > longestWordLength {
				longestWordLength = currentWordLength + 1
				longestWordBeginIndex = currentWordBeginIndex
			}
			continue
		}
		if string(symbol) == " " {
			if currentWordLength > longestWordLength {
				longestWordLength = currentWordLength
				longestWordBeginIndex = currentWordBeginIndex
			}
			currentWordLength = 0
			currentWordBeginIndex = i + 1
		} else {
			currentWordLength++
		}

	}

	return line[longestWordBeginIndex : longestWordBeginIndex+longestWordLength]
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	line := readLine(scanner)
	longestWord := getLongestWord(line)
	fmt.Println(longestWord)
	fmt.Println(len(longestWord))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
