package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//https://contest.yandex.ru/contest/26133/problems/A/

// https://contest.yandex.ru/contest/26133/run-report/88187940/ - ссылка на последнее ОК решение

//Вам даны строки в запакованном виде.
//Определим запакованную строку (ЗС) рекурсивно.
//Строка, состоящая только из строчных букв английского алфавита является ЗС.
//Если A и B —– корректные ЗС, то и AB является ЗС.
//Если A —– ЗС, а n — однозначное натуральное число, то n[A] тоже ЗС.
//При этом запись n[A] означает, что при распаковке строка A записывается подряд n раз.
//Найдите наибольший общий префикс распакованных строк и выведите его (в распакованном виде).
//
//Иными словами, пусть сложение —– это конкатенация двух строк, а умножение строки
//на число — повтор строки соответствующее число раз.
//Пусть функция f умеет принимать ЗС и распаковывать её.
//Если ЗС D имеет вид D=AB, где A и B тоже ЗС, то f(D) = f(A) + f(B).
//Если D=n[A], то f(D) = f(A) × n.

//Сначала распакуем все строки. Каждая строка распаковывается за один проход:
// если мы встречаем цифру, то записываем ее в стек мультипликаторов
// если встречаем открывающую скобку, то создаем билдер строк, который будем пополнять последующими буквами
// если встречаем букву, то записываем ее в билдер умножаемой строки. если билдера нет, то просто приписываем ее к результату
// если встречаем закрывающую скобку, то берем из стека множитель и умножаемую строку. если у нас еще остались
//  умножаемые строки в стеке, то приписываем умножаемую строку к последнему биллдеру в стеке
//  если умножаемых строк нет, то приписываем умножаемую строку к результату
//Далее пойдем по всем распакованным строкам одновременно с нуля:
//  пока все символы совпадают, добавляем их в префикс
//  как только найдем первый несовпадающий символ или какая-то строка закончится - возвращаем сформированный префикс

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// O(n*|s|), где n - количество строк, |s| - длина наибольшей строки

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// O(|s|), где |s| - длина наибольшей строки. эту память мы будем потреблять для хранения стеков при распаковке
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	unpackedStrings := make([]string, n)
	for i := 0; i < n; i++ {
		unpackedStrings[i] = unpackString(readString(scanner))
	}
	fmt.Println(longestCommonPrefix(unpackedStrings))
}

func unpackString(s string) string {
	var result strings.Builder
	var multipliersStack []int
	var toMultiplyStack []*strings.Builder
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			multiplier, _ := strconv.Atoi(string(s[i]))
			multipliersStack = append(multipliersStack, multiplier)
			continue
		}
		if unicode.IsLetter(rune(s[i])) {
			if len(toMultiplyStack) != 0 {
				toMultiplyStack[len(toMultiplyStack)-1].WriteByte(s[i])
			} else {
				result.WriteByte(s[i])
			}
			continue
		}
		if rune(s[i]) == '[' {
			toMultiplyStack = append(toMultiplyStack, &strings.Builder{})
			continue
		}
		if rune(s[i]) == ']' {
			multiplier := multipliersStack[len(multipliersStack)-1]
			multipliersStack = multipliersStack[:len(multipliersStack)-1]

			toMultiply := toMultiplyStack[len(toMultiplyStack)-1].String()
			toMultiplyStack = toMultiplyStack[:len(toMultiplyStack)-1]

			for j := 0; j < multiplier; j++ {
				if len(toMultiplyStack) == 0 {

					result.WriteString(toMultiply)
				} else {
					toMultiplyStack[len(toMultiplyStack)-1].WriteString(toMultiply)
				}
			}
			continue
		}
	}
	return result.String()
}

func longestCommonPrefix(strs []string) string {
	var commonPrefix strings.Builder
	for i := range strs[0] {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return commonPrefix.String()
			}
		}
		commonPrefix.WriteByte(strs[0][i])
	}

	return commonPrefix.String()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
