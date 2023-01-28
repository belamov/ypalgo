package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/F/

// Перед сном Рита решила поиграть в игру на телефоне.
//Дан массив целых чисел, в котором каждый элемент обозначает длину стороны треугольника.
//Нужно определить максимально возможный периметр треугольника, составленного из
//сторон с длинами из заданного массива. Помогите Рите скорее закончить игру и пойти спать.
//
//Напомним, что из трёх отрезков с длинами a ≤ b ≤ c можно составить треугольник, если
//выполнено неравенство треугольника: c < a + b
func getMaxPerimeter(sides []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(sides)))
	for i := 0; i < len(sides)-2; i++ {
		if sides[i] < sides[i+1]+sides[i+2] {
			return sides[i] + sides[i+1] + sides[i+2]
		}
	}
	return 0
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	sides := readArray(scanner)
	fmt.Println(getMaxPerimeter(sides))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
