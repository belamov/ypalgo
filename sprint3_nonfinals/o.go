package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23638/problems/O/
// https://leetcode.com/problems/find-k-th-smallest-pair-distance/description/

// Гоша долго путешествовал и измерил площадь каждого из n островов
//Алгосов, но ему этого мало! Теперь он захотел оценить, насколько
//разнообразными являются острова в составе архипелага.
//
//Для этого Гоша рассмотрел все пары островов (таких пар, напомним, n * (n-1) / 2) и посчитал
//попарно разницу площадей между всеми островами.
//Теперь он собирается упорядочить полученные разницы, чтобы взять k-ую по порядку из них.
//
//Помоги Гоше найти k-ю минимальную разницу между площадями эффективно.
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	l := 0
	r := nums[len(nums)-1] - nums[0]

	for l < r {
		guess := (r + l) / 2
		count := 0
		i := 0
		j := 1
		for i < len(nums) {
			for j < len(nums) && nums[j]-nums[i] <= guess {
				j++
			}
			count += j - i - 1
			i++
		}

		if count >= k {
			r = guess
		} else {
			l = guess + 1
		}
	}

	return l
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	k := readInt(scanner)
	fmt.Println(smallestDistancePair(arr, k))
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
