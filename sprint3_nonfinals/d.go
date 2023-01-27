package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23638/problems/D/

// К Васе в гости пришли одноклассники. Его мама решила угостить ребят печеньем.
//
//Но не всё так просто. Печенья могут быть разного размера.
//А у каждого ребёнка есть фактор жадности —– минимальный размер печенья, которое он возьмёт.
//Нужно выяснить, сколько ребят останутся довольными в лучшем случае, когда они действуют оптимально.
//
//Каждый ребёнок может взять не больше одного печенья.
func getSatisfiedCount(greed []int, cookies []int) int {
	sort.Ints(greed)
	sort.Ints(cookies)
	satisfied := 0
	cookiesIndex := 0
	greedIndex := 0
	for greedIndex < len(greed) && cookiesIndex < len(cookies) {
		if greed[greedIndex] > cookies[cookiesIndex] {
			cookiesIndex++
			continue
		}
		cookiesIndex++
		greedIndex++
		satisfied++
	}
	return satisfied
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	greed := readArray(scanner)
	readInt(scanner)
	cookies := readArray(scanner)
	fmt.Println(getSatisfiedCount(greed, cookies))
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
