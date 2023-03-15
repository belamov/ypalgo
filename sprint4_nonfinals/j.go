package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/J/
// https://leetcode.com/problems/4sum/description/

//У Гоши есть любимое число S.
//Помогите ему найти все уникальные четвёрки чисел в
//массиве, которые в сумме дают заданное число S.
func main() {
	scanner := makeScanner()
	readString(scanner)
	s := readInt(scanner)
	arr := readArray(scanner)

	quadruplets := fourSum(arr, s)
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(quadruplets)))
	writer.WriteString("\n")
	for _, quadruplet := range quadruplets {
		printArray(writer, quadruplet)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func fourSum(nums []int, target int) [][]int {
	history := make(map[int][][]int)
	quadruplets := make(map[string][]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			diff := target - sum
			if marchingDoublesIndicies, ok := history[diff]; ok {
				for _, doubleIndex := range marchingDoublesIndicies {
					if doubleIndex[0] == i || doubleIndex[1] == i || doubleIndex[0] == j || doubleIndex[1] == j {
						continue
					}
					quadruplet := []int{nums[doubleIndex[0]], nums[doubleIndex[1]], nums[i], nums[j]}
					sort.Ints(quadruplet)
					quadrupletKey := fmt.Sprintf("%d,%d,%d,%d", quadruplet[0], quadruplet[1], quadruplet[2], quadruplet[3])
					quadruplets[quadrupletKey] = quadruplet
				}
			}
			if doubles, ok := history[sum]; ok {
				history[sum] = append(doubles, []int{i, j})
			} else {
				history[sum] = make([][]int, 0)
				history[sum] = append(history[sum], []int{i, j})
			}
		}
	}

	result := make([][]int, 0, len(quadruplets))
	for _, quadruplet := range quadruplets {
		result = append(result, quadruplet)
	}

	sort.Slice(result[:], func(i, j int) bool {
		for x := range result[i] {
			if result[i][x] == result[j][x] {
				continue
			}
			return result[i][x] < result[j][x]
		}
		return false
	})

	return result
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

func printArray(writer *bufio.Writer, arr []int) {
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
}
