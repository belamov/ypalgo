package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22449/problems/D/

// Метеорологическая служба вашего города решила исследовать погоду новым способом.
//
//Под температурой воздуха в конкретный день будем понимать максимальную температуру в этот день.
//Под хаотичностью погоды за n дней служба понимает количество дней, в которые температура
//строго больше, чем в день до (если такой существует) и в день после текущего (если такой существует).
//
//Например, если за 5 дней максимальная температура воздуха составляла [1, 2, 5, 4, 8] градусов, то
//хаотичность за этот период равна 2: в 3-й и 5-й дни выполнялись описанные условия.

//Определите по ежедневным показаниям температуры хаотичность погоды за этот период.
//
//Заметим, что если число показаний n=1, то единственный день будет хаотичным.
func getWeatherRandomness(temperatures []int) int {
	if len(temperatures) == 1 {
		return 1
	}

	chaos := 0

	if temperatures[0] > temperatures[1] {
		chaos++
	}

	if temperatures[len(temperatures)-1] > temperatures[len(temperatures)-2] {
		chaos++
	}

	for i := 1; i < len(temperatures)-1; i++ {
		if temperatures[i] > temperatures[i-1] && temperatures[i] > temperatures[i+1] {
			chaos++
		}
	}

	return chaos
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	temperatures := readArray(scanner)
	fmt.Println(getWeatherRandomness(temperatures))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
