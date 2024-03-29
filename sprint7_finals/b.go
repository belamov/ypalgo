package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/25597/problems/B/

// https://contest.yandex.ru/contest/25597/run-report/87843902/ - ссылка на последнее ОК решение

//На Алгосах устроили турнир по настольному теннису.
//Гоша выиграл n партий, получив при этом некоторое количество очков за каждую из них.
//Гоше стало интересно, можно ли разбить все заработанные
//им во время турнира очки на две части так, чтобы сумма в них была одинаковой.

//Нас спрашивают, можем ли мы разделить массивы на две части,
//у которых равны суммы элементов. Это значит, что каждая такая часть будет иметь половину
//от суммы всего массива. То есть нас спрашивают, можем ли мы составить такой подмассив, у которого
//сумма элементов равна половине суммы элементов всего массива
//Данную задачу можно решить при помощи двумерного динамического программирования:

// - Что будет храниться в dp?
//   dp[i][j] - можем ли мы составить подмассив с суммой элементов = j из подмассива в котором лежат
//              элементы от 0 до i индекса. максимальное значение j = половина суммы элементов массива
//
// - Каким будет базовый случай для задачи?
//   из любого подмассива (i=0..n) мы сможем составить сумму равную 0 - просто не взяв ни один элемент
//   то есть dp[i][0] = true
//   также если наш подмассив состоит только из первого элемента (i=0), то мы сможем составить подмассив
//   с заданной суммой элементов j, если в нем будет элемент, равный j
//   то есть dp[0][arr[0]] = true
//
// - Каким будет переход динамики?
//	 мы сможем построить подмассив из элементов 0..i с суммой элементов j если:
//      - если не возьмем текущий i-й элемент и мы все равно сможем построить подмассив суммой j
//        эта возможность будет храниться в dp[i-1][j]
//      - если мы возьмем текущий i-й элемент - тогда нам нужно составить сумму j - arr[i] из
//        элементов 0..i-1 (поскольку i-й элемент мы уже взяли).
//        эта возможность будет храниться в dp[i-1][j-arr[i]]
//   итак, переход динамики будет:
//   dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i]]
//   также заметим, что для расчета текущей строки нам нужна только предыдущая строка, а значит мы можем
//   не хранить весь массив dp, а обойтись только текущей и предыдущей строкой. но поскольку во внутреннем
//   цикле мы можем идти в обратном направлении, то можем использовать только один массив, так как мы не будем
//   перезаписывать значения
//
// - Каким будет порядок вычисления данных в массиве dp?
//   сначала заполняем возможности составления всех сумм от 1 до j для подмассива, состоящего из перового
//   элемента, потом из двух первых элементов, потом из трех и так далее. то есть сначала заполняем первую
//   строку dp, потом вторую и так далее
//
// - Где будет располагаться ответ на исходный вопрос?
//   нас интересуют значения в последнем столбце dp
//   если в каком-нибудь моменте в нем будет true, то мы можем сказать, что нашли способ
//   составить подмассив с нужной суммой элементов

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// O(N*S), где N - количество элементов в массиве, S - полусумма всех элементов в массиве
// у нас два вложенных цикла - внешний идет от 1 до N, внутренний - от S до 1. действия внутри цикла константны

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// O(S), где S - полусумма всех элементов в массиве
// мы храним только один массив размером S+1
func main() {
	scanner := makeScanner()
	readInt(scanner)
	ratings := readArray(scanner)

	if canPartition(ratings) {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}

func canPartition(nums []int) bool {
	fullSum := 0
	for _, rating := range nums {
		fullSum += rating
	}

	if fullSum%2 != 0 {
		return false
	}

	halfSum := fullSum / 2

	curr := make([]bool, halfSum+1)

	// базовые случаи
	curr[0] = true
	if nums[0] <= halfSum {
		curr[nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := halfSum; j >= 1; j-- {
			prevWithI := false
			if j-nums[i] >= 0 {
				prevWithI = curr[j-nums[i]]
			}
			curr[j] = curr[j] || prevWithI
			if curr[j] && j == halfSum {
				return true
			}
		}
	}

	return false
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
