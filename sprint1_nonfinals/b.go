package main

import (
	"fmt"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/B/

// Представьте себе онлайн-игру для поездки в метро: игрок нажимает на кнопку, и на экране
// появляются три случайных числа. Если все три числа оказываются одной чётности, игрок выигрывает.
//
// Напишите программу, которая по трём числам определяет, выиграл игрок или нет.
func checkParity(a int, b int, c int) bool {
	if a%2 == 0 && b%2 == 0 && c%2 == 0 {
		return true
	}
	if a%2 != 0 && b%2 != 0 && c%2 != 0 {
		return true
	}
	return false
}

func main() {
	a := readInt()
	b := readInt()
	c := readInt()
	if checkParity(a, b, c) {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
