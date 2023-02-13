package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23991/problems/G/

// ители Алгосов любят устраивать турниры по спортивному программированию.
//Все участники разбиваются на пары и соревнуются друг с другом.
//А потом два самых сильных программиста встречаются в финальной
//схватке, которая состоит из нескольких раундов.
//Если в очередном раунде выигрывает первый участник, в таблицу с результатами
//записывается 0, если второй, то 1.
//Ничьей в раунде быть не может.
//
//Нужно определить наибольший по длине непрерывный отрезок раундов, по результатам
//которого суммарно получается ничья.
//Например, если дана последовательность 0 0 1 0 1 1 1 0 0 0, то раунды
//с 2-го по 9-й (нумерация начинается с единицы) дают ничью.
func getMaxSegment(rounds []int) int {
	advantage := 0
	maxSegmentLength := 0
	advantages := make(map[int]int)
	advantages[0] = 0
	for i := 0; i < len(rounds); i++ {
		if rounds[i] == 1 {
			advantage++
		} else {
			advantage--
		}

		advantagePosition, advantageWasInPast := advantages[advantage]
		if advantageWasInPast {
			if i-advantagePosition+1 > maxSegmentLength {
				maxSegmentLength = i - advantagePosition + 1
			}
		} else {
			advantages[advantage] = i + 1
		}
	}

	return maxSegmentLength
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	rounds := readArray(scanner)
	fmt.Println(getMaxSegment(rounds))
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
