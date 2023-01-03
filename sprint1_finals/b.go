package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22450/problems/B/

//###############################################################################################
//
// айди успешной посылки 80077842: https://contest.yandex.ru/contest/22450/run-report/80077842/
//
//###############################################################################################

// Игра «Тренажёр для скоростной печати» представляет собой поле из клавиш 4x4.
//В нём на каждом раунде появляется конфигурация цифр и точек.
//На клавише написана либо точка, либо цифра от 1 до 9.
//
//В момент времени t игрок должен одновременно нажать на все клавиши, на которых написана цифра t.
//Гоша и Тимофей могут нажать в один момент времени на k клавиш каждый.
//Если в момент времени t нажаты все нужные клавиши, то игроки получают 1 балл.
//
//Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём.
func getMaximumPoints(game []int, availableEach int) int {
	maximumPoints := 0
	const playersCount = 2

	// составим мапу, формата раунд - количество клавиш, необходимое нажать в раунде
	rounds := make([]int, 10)

	for _, t := range game {
		rounds[t]++
	}

	for _, required := range rounds {
		// если количество клавиш не больше, чем количество клавиш, которое могут нажать игроки
		// и если эти раунды были
		// то за этот раунд они смогут получить балл
		if required <= availableEach*playersCount && required > 0 {
			maximumPoints++
		}
	}

	return maximumPoints
}

func main() {
	scanner := makeScanner()
	k := readInt(scanner)
	game := readGame(scanner)
	fmt.Println(getMaximumPoints(game, k))
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readGame(scanner *bufio.Scanner) []int {
	// нам ни к чему хранить матрицу, нам нужны только значения t
	// поэтому можем сложить их в обычный массив
	game := make([]int, 0, 16)
	for i := 0; i < 4; i++ {
		line := readLine(scanner)
		for j := 0; j < 4; j++ {
			if string(line[j]) != "." {
				t, _ := strconv.Atoi(string(line[j]))
				game = append(game, t)
			}
		}
	}

	return game
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 100 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
