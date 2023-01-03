package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/problems/A/

//###############################################################################################
//
// айди успешной посылки 80077757: https://contest.yandex.ru/contest/22450/run-report/80077757/
//
//###############################################################################################

// Тимофей ищет место, чтобы построить себе дом.
//Улица, на которой он хочет жить, имеет длину n, то есть состоит из n одинаковых идущих подряд участков.
//Каждый участок либо пустой, либо на нём уже построен дом.
//
//Общительный Тимофей не хочет жить далеко от других людей на этой улице.
//Поэтому ему важно для каждого участка знать расстояние до ближайшего пустого участка.
//Если участок пустой, эта величина будет равна нулю — расстояние до самого себя.
//
//Помогите Тимофею посчитать искомые расстояния. Для этого у вас есть карта улицы.
//Дома в городе Тимофея нумеровались в том порядке, в котором
//строились, поэтому их номера на карте никак не упорядочены.
//Пустые участки обозначены нулями.
func getDistances(emptySectorsPositions []int, streetLength int) []int {
	distances := make([]int, streetLength)

	leftSideLimit := 0
	rightSideLimit := 0

	//пройдем по позициям пустых участков
	for i := 0; i < len(emptySectorsPositions); i++ {
		// заполним расстояния до середины между текущим пустым участком и предыдущим пустым участком.
		// если это первый пустой участок, то до начала улицы
		leftSideLimit = -1
		if i != 0 {
			leftSideLimit = (emptySectorsPositions[i] + emptySectorsPositions[i-1]) / 2
		}
		for currentSectorIndex := emptySectorsPositions[i] - 1; currentSectorIndex > leftSideLimit; currentSectorIndex-- {
			distances[currentSectorIndex] = emptySectorsPositions[i] - currentSectorIndex
		}

		// заполним расстояния до середины между текущим пустым участком и следующим пустым участком.
		// если это последний пустой участок, то до конца улицы
		rightSideLimit = streetLength - 1
		if i+1 != len(emptySectorsPositions) {
			rightSideLimit = (emptySectorsPositions[i] + emptySectorsPositions[i+1]) / 2
		}
		for currentSectorIndex := emptySectorsPositions[i] + 1; currentSectorIndex <= rightSideLimit; currentSectorIndex++ {
			distances[currentSectorIndex] = currentSectorIndex - emptySectorsPositions[i]
		}

		//массив расстояний уже проинициализирован нулями,
		//поэтому на пустых участках уже будут нули
	}

	return distances
}

func main() {
	scanner := makeScanner()
	streetLength := readInt(scanner)
	emptySectorsPositions := readStreet(scanner)
	printArray(getDistances(emptySectorsPositions, streetLength))
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readStreet(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, 0, len(listString))
	for i := 0; i < len(listString); i++ {
		//нам не важны номера домов, важна только занятость/незанятость участка и длина улицы
		//запомним позиции пустых участков
		if listString[i] == "0" {
			arr = append(arr, i)
		}
	}
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 100 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
