package main

import (
	"fmt"
	"strconv"
)

// https://contest.yandex.ru/contest/22449/problems/A/

// Вася делает тест по математике: вычисляет значение функций в различных точках.
// Стоит отличная погода, и друзья зовут Васю гулять. Но мальчик решил сначала закончить тест и только после этого
// идти к друзьям. К сожалению, Вася пока не умеет программировать. Зато вы умеете. Помогите Васе написать
// код функции, вычисляющей y = ax2 + bx + c. Напишите программу, которая будет
// по коэффициентам a, b, c и числу x выводить значение функции в точке x.
func evaluateFunction(a int, b int, c int, x int) int {
	return a*x*x + b*x + c
}

func main() {
	a := readInt()
	x := readInt()
	b := readInt()
	c := readInt()
	fmt.Println(evaluateFunction(a, b, c, x))
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
