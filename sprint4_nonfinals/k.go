package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/K/

//Гоша едет в гости к друзьям.
//Ему придётся сначала ехать на метро, а потом пересаживаться на автобус.
//Гоша не любит долго ждать, поэтому хочет выбрать такую станцию метро, рядом с которой
//расположено как можно больше остановок автобуса.
//Гоша считает, что остановка находится рядом с метро, если расстояние между ними не превосходит 20 метров.
//
//Гоше известны все координаты автобусных остановок и координаты выходов из метро.
//Помогите ему найти выход из метро, рядом с которым расположено больше всего остановок.
func main() {
	scanner := makeScanner()

	n := readInt(scanner)
	metros := make([][]int, n)
	for i := 0; i < n; i++ {
		metros[i] = readArray(scanner)
	}

	m := readInt(scanner)
	busStops := make([][]int, m)
	for i := 0; i < m; i++ {
		busStops[i] = readArray(scanner)
	}

	d := 20
	fmt.Println(getMetroWithMostCloseBusStops(metros, busStops, d))
}

func getMetroWithMostCloseBusStops(metros [][]int, busStops [][]int, d int) int {
	busStopSearchIndex := make(map[int][]int)
	for _, busStop := range busStops {
		if busStopSearchIndex[busStop[0]] == nil {
			busStopSearchIndex[busStop[0]] = make([]int, 0)
		}
		busStopSearchIndex[busStop[0]] = append(busStopSearchIndex[busStop[0]], busStop[1])
	}

	maximumCloseBusStopCount := 0
	metroWithMostCloseBusStops := -1
	closeBusStopsCount := make([]int, len(metros))

	for metroIndex, metro := range metros {
		for i := metro[0] - d; i <= metro[0]+d; i++ {
			if busStopsWithCloseX, okX := busStopSearchIndex[i]; okX {
				for _, y := range busStopsWithCloseX {
					if isBusStopCloseToMetro(metro, i, y, d) {
						closeBusStopsCount[metroIndex]++
						if closeBusStopsCount[metroIndex] > maximumCloseBusStopCount {
							maximumCloseBusStopCount = closeBusStopsCount[metroIndex]
							metroWithMostCloseBusStops = metroIndex
						}
					}
				}
			}
		}
	}

	return metroWithMostCloseBusStops + 1
}

func isBusStopCloseToMetro(metro []int, x int, y int, d int) bool {
	return d*d >= (x-metro[0])*(x-metro[0])+(y-metro[1])*(y-metro[1])
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
