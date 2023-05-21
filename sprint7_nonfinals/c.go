package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25596/problems/C/
//Гуляя по одному из островов Алгосского архипелага, Гоша набрёл на пещеру, в которой лежат кучи золотого песка.
//К счастью, у Гоши есть с собой рюкзак грузоподъёмностью до M килограмм, поэтому он
//может унести с собой какое-то ограниченное количество золота.
//
//Всего золотых куч n штук, и все они разные.
//В куче под номером i содержится mi килограммов золотого песка, а стоимость одного килограмма — ci алгосских франков.
//
//Помогите Гоше наполнить рюкзак так, чтобы общая стоимость золотого песка в пересчёте на алгосские франки была максимальной.
func main() {
	scanner := makeScanner()
	M := readInt(scanner)
	n := readInt(scanner)

	piles := make([]*Pile, n)
	for i := 0; i < n; i++ {
		pileData := readArray(scanner)
		piles[i] = &Pile{cost: pileData[0], amount: pileData[1]}
	}

	total := 0

	for M > 0 {
		mostExpensivePileIndex := getMostExpensivePile(piles)

		if mostExpensivePileIndex == -1 {
			break
		}

		var availableAmount int

		if M >= piles[mostExpensivePileIndex].amount {
			availableAmount = piles[mostExpensivePileIndex].amount
		} else {
			availableAmount = M
		}

		M -= availableAmount
		total += availableAmount * piles[mostExpensivePileIndex].cost
		piles[mostExpensivePileIndex].amount -= availableAmount
	}

	fmt.Print(total)
}

func getMostExpensivePile(piles []*Pile) int {
	maxPile := &Pile{cost: -1, amount: 0}
	maxIndex := -1
	for i, pile := range piles {
		if pile.cost > maxPile.cost && pile.amount > 0 {
			maxIndex = i
			maxPile = pile
		}
	}
	return maxIndex
}

type Pile struct {
	cost   int
	amount int
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
