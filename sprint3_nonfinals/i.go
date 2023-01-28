package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/I/

// На IT-конференции присутствовали студенты из разных вузов со всей страны.
//Для каждого студента известен ID университета, в котором он учится.
//
//Тимофей предложил Рите выяснить, из каких k вузов на конференцию пришло больше всего учащихся.
func getRating(students []int, k int) []int {
	uniRatings := make(map[int]int)
	for _, university := range students {
		uniRatings[university]++
	}

	ratings := make([][]int, 0, 10000)
	for uni, students := range uniRatings {
		ratings = append(ratings, []int{uni, students})
	}

	sort.Slice(ratings, func(i, j int) bool {
		if ratings[i][1] == ratings[j][1] {
			return ratings[i][0] < ratings[j][0]
		}
		return ratings[i][1] > ratings[j][1]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = ratings[i][0]
	}

	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	students := readArray(scanner)
	k := readInt(scanner)
	printArray(getRating(students, k))
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
