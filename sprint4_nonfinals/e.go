package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/23991/problems/E/
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

//На вход подается строка.
//Нужно определить длину наибольшей подстроки, которая не содержит повторяющиеся символы.
func main() {
	scanner := makeScanner()
	s := readString(scanner)
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	charHistory := make(map[uint8]interface{})
	longestLength := 0
	currentLength := 0
	beginPointer := 0
	endPointer := 0
	for beginPointer != len(s) {
		currentChar := s[beginPointer]
		_, charAlreadySeen := charHistory[currentChar]
		if !charAlreadySeen {
			beginPointer++
			currentLength++
			if currentLength > longestLength {
				longestLength = currentLength
			}
			charHistory[currentChar] = nil
			continue
		}

		for s[endPointer] != s[beginPointer] {
			delete(charHistory, s[endPointer])
			endPointer++
			currentLength--
		}
		delete(charHistory, s[endPointer])
		endPointer++
		currentLength--

	}
	return longestLength
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
