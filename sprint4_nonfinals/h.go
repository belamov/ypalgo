package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://contest.yandex.ru/contest/23991/problems/H/
//https://leetcode.com/problems/isomorphic-strings/

// Жители Алгосского архипелага придумали новый способ сравнения строк.
//Две строки считаются равными, если символы одной из них можно заменить на символы
//другой так, что первая строка станет точной копией второй строки.
//При этом необходимо соблюдение двух условий:
//
// - Порядок вхождения символов должен быть сохранён.
// - Одинаковым символам первой строки должны соответствовать одинаковые
//   символы второй строки. Разным символам —– разные.

//Например, если строка s = «abacaba», то ей будет равна
//строка t = «xhxixhx», так как все вхождения «a» заменены на «x», «b» –— на «h», а «c» –— на «i».
//Если же первая строка s=«abc», а вторая t=«aaa», то строки уже не будут
//равны, так как разные буквы первой строки соответствуют одинаковым буквам второй.
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := make(map[byte]byte)
	hashRev := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		existingMorph, ok := hash[s[i]]
		if ok && existingMorph != t[i] {
			return false
		}

		revMorph, ok := hashRev[t[i]]
		if ok && revMorph != s[i] {
			return false
		}

		hash[s[i]] = t[i]
		hashRev[t[i]] = s[i]
	}
	return true
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	t := readString(scanner)
	if isIsomorphic(s, t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
