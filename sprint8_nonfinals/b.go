package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/26131/problems/B/
//Представьте, что вы работаете пограничником и постоянно проверяете документы людей по записи из базы.
//При этом допустима ситуация, когда имя человека в базе отличается от имени в паспорте на одну
//замену, одно удаление или одну вставку символа.
//Если один вариант имени может быть получен из другого удалением одного
//символа, то человека пропустят через границу.
//А вот если есть какое-либо второе изменение, то человек грустно поедет домой или в посольство.
//
//Например, если первый вариант —– это «Лена», а второй — «Лера», то девушку пропустят.
//Также человека пропустят, если в базе записано «Коля», а в паспорте — «оля».
//
//Однако вариант, когда в базе числится «Иннокентий», а в паспорте
//написано «ннакентий», уже не сработает.
//Не пропустят также человека, у которого в паспорте
//записан «Иинннокентий», а вот «Инннокентий» спокойно пересечёт границу.
//
//Напишите программу, которая сравнивает имя в базе с именем в паспорте и решает, пропускать
//человека или нет. В случае равенства двух строк — путешественника, естественно, пропускают.
func main() {
	scanner := makeScanner()
	required := readString(scanner)
	actual := readString(scanner)

	if passing(required, actual) {
		fmt.Print("OK")
	} else {
		fmt.Print("FAIL")
	}
}

func passing(required string, actual string) bool {
	s := required
	t := actual
	if len(actual) > len(required) {
		s = actual
		t = required
	}

	if len(s)-len(t) > 1 {
		return false
	}

	if len(s) == 0 || len(t) == 0 {
		return true
	}

	sP := 0
	tP := 0
	mistakeHappen := false
	for sP < len(s) {
		if s[sP] == t[tP] {
			sP++
			tP++
			continue
		}

		if mistakeHappen {
			return false
		}

		mistakeHappen = true
		if len(s) > len(t) {
			sP++
			continue
		}

		sP++
		tP++
	}

	return true
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
