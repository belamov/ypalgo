package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23991/problems/I/
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/

//Гоша увлёкся хоккеем и часто смотрит трансляции матчей.
//Чтобы более-менее разумно оценивать силы команд, он сравнивает очки, набранные во всех матчах каждой командой.
//
//Гоша попросил вас написать программу, которая по результатам игр двух выбранных
//команд найдёт наибольший по длине отрезок матчей, когда эти команды зарабатывали одинаковые очки.
func main() {
	scanner := makeScanner()
	readString(scanner)
	nums1 := readArray(scanner)
	readString(scanner)
	nums2 := readArray(scanner)

	fmt.Println(findLength(nums1, nums2))
}
func findLength(nums1 []int, nums2 []int) int {
	a := 331
	m := 201893
	rollingHash1 := rollingHash(a, m, nums1)
	rollingHash2 := rollingHash(a, m, nums2)
	powers := getPowers(nums1, nums2, a, m)

	maxAvailableLength := len(nums1)
	if len(nums2) < maxAvailableLength {
		maxAvailableLength = len(nums2)
	}

	left := 0
	right := maxAvailableLength
	maxLength := 0
	for left <= right {
		mid := (left + right) / 2
		if checkIfArraysHasCommonSubarrays(mid, nums1, nums2, rollingHash1, rollingHash2, powers, m) {
			maxLength = mid
			left = mid + 1
			continue
		}
		right = mid - 1
	}

	return maxLength
}

func rollingHash(a int, m int, arr []int) []int {
	hash := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		hash[i+1] = (hash[i]*a + arr[i]) % m
	}

	return hash
}

func getPowers(arr1, arr2 []int, a, m int) []int {
	n := len(arr1)
	if len(arr2) > n {
		n = len(arr2)
	}
	powers := make([]int, n+2)
	powers[0] = 1

	for i := 0; i < n; i++ {
		powers[i+1] = (powers[i] * a) % m
	}

	return powers
}

func checkIfArraysHasCommonSubarrays(subarrayLength int, nums1 []int, nums2 []int, hash1 []int, hash2 []int, powers []int, m int) bool {
	if subarrayLength == 0 {
		return true
	}
	hashes := make(map[int][]int)
	for i := 0; i < len(nums1)-subarrayLength+1; i++ {
		hashKey := hashSubstring(m, i, i+subarrayLength-1, hash1, powers)
		hashes[hashKey] = append(hashes[hashKey], i)
	}
	for i := 0; i < len(nums2)-subarrayLength+1; i++ {
		hashKey := hashSubstring(m, i, i+subarrayLength-1, hash2, powers)
		if foundSubarrays, ok := hashes[hashKey]; ok {
			for _, subarrayBeginIndex := range foundSubarrays {
				if isArraysEqual(nums1[subarrayBeginIndex:subarrayBeginIndex+subarrayLength-1], nums2[i:i+subarrayLength-1]) {
					return true
				}
			}
		}
	}
	return false

}

func isArraysEqual(ints []int, ints2 []int) bool {
	for i := 0; i < len(ints); i++ {
		if ints[i] != ints2[i] {
			return false
		}
	}
	return true
}

func hashSubstring(m int, left int, right int, hashes []int, powers []int) int {
	hash := hashes[right+1] - hashes[left]*powers[right-left+1]
	hash = mathematicalModulus(hash, m)
	return hash
}

func mathematicalModulus(d, m int) int {
	return (d%m + m) % m
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
