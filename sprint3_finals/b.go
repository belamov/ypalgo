package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/23815/problems/B/

// https://contest.yandex.ru/contest/23815/run-report/81902439/ - ссылка на последнее ОК решение

// Тимофей решил организовать соревнование по спортивному программированию, чтобы найти талантливых стажёров.
//Задачи подобраны, участники зарегистрированы, тесты написаны.
//Осталось придумать, как в конце соревнования будет определяться победитель.
//
//Каждый участник имеет уникальный логин.
//Когда соревнование закончится, к нему будут привязаны два показателя: количество решённых
//задач Pi и размер штрафа Fi.
//Штраф начисляется за неудачные попытки и время, затраченное на задачу.
//
//Тимофей решил сортировать таблицу результатов следующим образом: при сравнении
//двух участников выше будет идти тот, у которого решено больше задач.
//При равенстве числа решённых задач первым идёт участник с меньшим штрафом.
//Если же и штрафы совпадают, то первым будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке.
//
//Тимофей заказал толстовки для победителей и накануне поехал за ними в магазин.
//В своё отсутствие он поручил вам реализовать алгоритм быстрой сортировки (англ. quick sort) для
//таблицы результатов. Так как Тимофей любит спортивное программирование и не любит
//зря расходовать оперативную память, то ваша реализация сортировки
//не может потреблять O(n) дополнительной памяти
//для промежуточных данных (такая модификация быстрой сортировки называется "in-place").
//
//Как работает in-place quick sort
//
//Как и в случае обычной быстрой сортировки, которая использует дополнительную память,
//необходимо выбрать опорный элемент (англ. pivot), а затем переупорядочить массив.
//Сделаем так, чтобы сначала шли элементы, не превосходящие опорного, а затем —– большие опорного.
//
//Затем сортировка вызывается рекурсивно для двух полученных частей. Именно на
//этапе разделения элементов на группы в обычном алгоритме используется
//дополнительная память. Теперь разберёмся, как реализовать этот шаг in-place.
//
//Пусть мы как-то выбрали опорный элемент.
//Заведём два указателя left и right, которые изначально будут указывать на левый и правый концы отрезка соответственно.
//Затем будем двигать левый указатель вправо до тех пор, пока он указывает на элемент, меньший (больший, если соотируем
//по убыванию) опорного.
//Аналогично двигаем правый указатель влево, пока он стоит на элементе, превосходящим (не превосходящим, если
//сортируем по убыванию) опорный.
//В итоге окажется, что что левее от left все элементы точно принадлежат первой группе, а правее от right — второй.
//Элементы, на которых стоят указатели, нарушают порядок.
//Поменяем их местами и продвинем указатели на следующие элементы.
//Будем повторять это действие до тех пор, пока left и right не столкнутся.
//
// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// как и быстрая сортировка, алгоритм работает за O(n^2) в худшем случае
// но поскольку мы выбираем опорный элемент каждый раз случайно, то в среднем алгоритм будет работать за O(n*log(n)),
// где n - длина массива
//
// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// мы заводим только указатели, поэтому алгоритм потребляет O(1) памяти и не зависит от длины массива
func QuickSortInPlace(arr []*Participant, left int, right int) {
	if right <= left {
		return
	}

	// выбираем случайный опорный элемент
	pivot := arr[rand.Intn(right-left)+left]

	// заводим указатели
	l, r := left, right

	// пока указатели не столкнутся
	for l < r {
		// двигаем левый указатель вправо до тех пор, пока он
		// указывает на элемент, больший опорного
		for arr[l].Bigger(pivot) {
			l++
		}

		// двигаем правый указатель влево, пока он стоит на элементе, не превосходящим опорный
		for arr[r].Less(pivot) {
			r--
		}

		// В итоге окажется, что что левее от left все элементы точно
		//принадлежат первой группе, а правее от right — второй.
		// если все элементы уже стоят как нужно, можем прерваться
		if l > r {
			break
		}

		// Элементы, на которых стоят указатели, нарушают порядок.
		// Поменяем их местами и продвинем указатели на следующие элементы.
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	QuickSortInPlace(arr, left, r)
	QuickSortInPlace(arr, l, right)
}

type Participant struct {
	login   string
	points  int
	penalty int
}

func (p *Participant) Bigger(participant *Participant) bool {
	if p.points != participant.points {
		return p.points > participant.points
	}
	if p.penalty != participant.penalty {
		return p.penalty < participant.penalty
	}
	return p.login < participant.login
}

func (p *Participant) Less(participant *Participant) bool {
	if p.points != participant.points {
		return p.points < participant.points
	}
	if p.penalty != participant.penalty {
		return p.penalty > participant.penalty
	}
	return p.login > participant.login
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	participants := make([]*Participant, n)
	for i := 0; i < n; i++ {
		participants[i] = getParticipant(scanner)
	}
	QuickSortInPlace(participants, 0, n-1)
	for _, participant := range participants {
		fmt.Println(participant.login)
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 7 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func getParticipant(scanner *bufio.Scanner) *Participant {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	points, _ := strconv.Atoi(listString[1])
	penalty, _ := strconv.Atoi(listString[2])
	return &Participant{
		login:   listString[0],
		points:  points,
		penalty: penalty,
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}