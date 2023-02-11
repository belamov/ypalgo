package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/23991/problems/D/

//В компании, где работает Тимофей, заботятся о досуге сотрудников и устраивают
//различные кружки по интересам.
//Когда кто-то записывается на занятие, в лог вносится название кружка.
//
//По записям в логе составьте список всех кружков, в которые ходит хотя бы один человек.
func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	names := make(map[string]interface{})
	for i := 0; i < n; i++ {
		scanner.Scan()
		name := scanner.Text()
		_, ok := names[name]
		if ok {
			continue
		}
		names[name] = nil
		fmt.Println(name)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
