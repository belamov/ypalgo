package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/B/
//Дано количество учебных занятий, проходящих в одной аудитории.
//Для каждого из них указано время начала и конца.
//Нужно составить расписание, в соответствии с которым в классе можно будет провести как можно больше занятий.
//
//Если возможно несколько оптимальных вариантов, то выведите любой.
//Возможно одновременное проведение более чем одного занятия нулевой длительности.
func main() {
	scanner := makeScanner()
	n := readInt(scanner)

	lectures := make([][]float64, n)
	for i := 0; i < n; i++ {
		lectures[i] = readArray(scanner)
	}

	timetable := getTimetable(lectures)

	fmt.Println(len(timetable))
	printMatrix(timetable)
}

func getTimetable(lectures [][]float64) [][]float64 {
	currentTime := -1.0
	timetable := make([][]float64, 0)
	var minLecture []float64
	for len(lectures) > 0 {
		lectures = filterNotStartedLectures(lectures, currentTime)
		if len(lectures) == 0 {
			break
		}
		minLecture, lectures = getMinLecture(lectures)

		timetable = append(timetable, minLecture)
		currentTime = minLecture[1]
	}

	return timetable
}

func filterNotStartedLectures(lectures [][]float64, startTime float64) [][]float64 {
	filtered := make([][]float64, 0)
	for _, lecture := range lectures {
		if lecture[0] >= startTime {
			filtered = append(filtered, lecture)
		}
	}
	return filtered
}

func getMinLecture(lectures [][]float64) ([]float64, [][]float64) {
	minLecture := []float64{100, 99}
	minLectureIndex := -1

	for i, lecture := range lectures {
		if lecture[1] < minLecture[1] || (minLecture[1] == lecture[1] && lecture[0] < minLecture[0]) {
			minLecture = lecture
			minLectureIndex = i
		}
	}

	lectures[minLectureIndex], lectures[len(lectures)-1] = lectures[len(lectures)-1], lectures[minLectureIndex]
	lectures = lectures[:len(lectures)-1]

	return minLecture, lectures
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readArray(scanner *bufio.Scanner) []float64 {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]float64, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.ParseFloat(listString[i], 64)
	}
	return arr
}

func printMatrix(matrix [][]float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			writer.WriteString(strconv.FormatFloat(matrix[i][j], 'f', -1, 64))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}
