package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/K/
//В мире последовательностей нет гороскопов.
//Поэтому когда две последовательности хотят понять, могут ли они счастливо жить
//вместе, они оценивают свою совместимость как длину их наибольшей общей подпоследовательности.
//
//Подпоследовательность получается из последовательности удалением некоторого (возможно, нулевого) числа элементов.
//То есть элементы сохраняют свой относительный порядок, но не обязаны изначально идти подряд.
//
//Найдите наибольшую общую подпоследовательность двух одиноких последовательностей и выведите её!
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	a := readArray(scanner)
	m := readInt(scanner)
	b := readArray(scanner)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i] == b[j] {
				dp[i][j] = getPrev(dp, i-1, j-1) + 1
			} else {
				dp[i][j] = max(getPrev(dp, i, j-1), getPrev(dp, i-1, j))
			}
		}
	}

	fmt.Println(dp[n-1][m-1])

	aIndicies := make([]int, 0, n)
	bIndicies := make([]int, 0, m)
	i := n - 1
	j := m - 1
	for getPrev(dp, i, j) != 0 {
		if a[i] == b[j] {
			aIndicies = append(aIndicies, i+1)
			bIndicies = append(bIndicies, j+1)
			i--
			j--
			continue
		}

		if getPrev(dp, i-1, j) == dp[i][j] {
			i--
			continue
		}

		j--
	}

	reverse(aIndicies)
	reverse(bIndicies)

	printArray(aIndicies)
	fmt.Println()
	printArray(bIndicies)

}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func getPrev(dp [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	writer.Flush()
}
