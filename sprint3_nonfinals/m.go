package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/M/
//https://leetcode.com/problems/median-of-two-sorted-arrays/

// На каждом острове в архипелаге Алгосы живёт какое-то количество
//людей или же остров необитаем (тогда на острове живёт 0 людей).
//Пусть на i-м острове численность населения составляет ai.
//Тимофей захотел найти медиану среди всех значений численности населения.
//
//Определение: Медиана массива чисел a_i —– это такое число, что половина
//чисел из массива не больше него, а другая половина не меньше.
//В общем случае медиану массива можно найти, отсортировав числа и взяв среднее из них.
//Если количество чисел чётно, то возьмём в качестве медианы полусумму соседних средних чисел, (a[n/2] + a[n/2 + 1])/2.
//
//У Тимофея уже есть отдельно данные по северной части архипелага и по южной, причём
//значения численности населения в каждой группе отсортированы по неубыванию.
//
//Определите медианную численность населения по всем островам Алгосов.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var a []int
	var b []int
	if len(nums1) > len(nums2) {
		a = nums2
		b = nums1
	} else {
		a = nums1
		b = nums2
	}

	n := len(nums1) + len(nums2)

	half := (n + 1) / 2
	l := 0
	r := len(a) - 1

	var al, ar, bl, br int
	for {
		ma := int(math.Floor(float64(l+r) / 2))
		mb := half - ma - 2

		if ma >= 0 {
			al = a[ma]
		} else {
			al = -math.MaxInt
		}

		if ma+1 < len(a) {
			ar = a[ma+1]
		} else {
			ar = math.MaxInt
		}

		if mb >= 0 {
			bl = b[mb]
		} else {
			bl = -math.MaxInt
		}

		if mb+1 < len(b) {
			br = b[mb+1]
		} else {
			br = math.MaxInt
		}

		if al > br {
			r = ma - 1
			continue
		}

		if bl > ar {
			l = ma + 1
			continue
		}

		if n%2 == 0 {
			return float64(max(al, bl)+min(ar, br)) / 2
		}
		return float64(max(al, bl))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	readInt(scanner)
	nums1 := readArray(scanner)
	nums2 := readArray(scanner)
	fmt.Println(findMedianSortedArrays(nums1, nums2))
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.WriteString("\n")
	writer.Flush()
}
