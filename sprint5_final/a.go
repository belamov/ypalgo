package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/24810/problems/A/

// https://contest.yandex.ru/contest/24810/run-report/86794938/ - ссылка на последнее ОК решение

// Тимофей решил организовать соревнование по спортивному программированию, чтобы найти талантливых стажёров.
//Задачи подобраны, участники зарегистрированы, тесты написаны.
//Осталось придумать, как в конце соревнования будет определяться победитель.
//
//Каждый участник имеет уникальный логин.
//Когда соревнование закончится, к нему будут привязаны два показателя: количество решённых
//задач Pi и размер штрафа Fi.
//Штраф начисляется за неудачные попытки и время, затраченное на задачу.
//
//Тимофей решил сортировать таблицу результатов следующим образом: при сравнении двух
//участников выше будет идти тот, у которого решено больше задач.
//При равенстве числа решённых задач первым идёт участник с меньшим штрафом.
//Если же и штрафы совпадают, то первым будет тот, у которого логин
//идёт раньше в алфавитном (лексикографическом) порядке.
//
//Тимофей заказал толстовки для победителей и накануне поехал за ними в магазин.
//В своё отсутствие он поручил вам реализовать алгоритм сортировки кучей (англ. Heapsort) для таблицы результатов.
//

//Общий алгоритм такой:
// 1. Создадим пустую бинарную неубывающую кучу (min-heap)
// 2. Вставим в неё по одному все элементы массива, сохраняя свойства кучи.
//    Так как нам нужна сортировка от меньшего к большему, на вершине пирамиды должен
//    оказаться самый маленький элемент.
//    Если бы мы захотели реализовать сортировку по убыванию — на вершине был бы самый большой элемент.
// 3. Будем извлекать из неё наиболее приоритетные элементы (с самым маленьким значением), удаляя их из кучи.

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// Сложность пирамидальной сортировки в худшем случае — O(n*log(n)), где n - количество сортируемых элементов
// Первый шаг — создание бинарной кучи. Сложность этой операции — O(1).
// Нам просто нужно выделить память под массив из n элементов.
//
// Далее вставим n элементов подряд в бинарную кучу.
// Сложность этого этапа:
// O(log(1)) + O(log(2)) + ... + O(log(n)) ~ O(log(n)) + O(log(n)) + ... + O(log(n)) ~ O(n*log(n))

// Последним шагом извлекаем n элементов. Сложность этой операции также не больше, чем O(n*log(n)):
// O(log(n)) + ... + O(log(2)) + O(log(1)) ~ O(log(n)) + ... + O(log(n)) + O(log(n)) ~ O(n*log(n))

// Получим:
// O(1) + O(n*log(n)) + O(n*log(n)) ~ O(n*log(n))

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// пространственная сложность - О(1)
// мы сразу складываем элементы в кучу, поэтому нам не требуется дополнительной памяти
func main() {
	scanner := makeScanner()
	n := readInt(scanner)

	//создаем пустую бинарную кучу
	participantsHeap := &Heap{participants: make([]*Participant, 0, n)}

	//Вставим в неё по одному всех участников, сохраняя свойства кучи
	for i := 0; i < n; i++ {
		participantsHeap.push(getParticipant(scanner))
	}

	for i := 0; i < n; i++ {
		fmt.Println(participantsHeap.pop().login)
	}
}

type Heap struct {
	participants []*Participant
}

func (heap *Heap) push(participant *Participant) {
	heap.participants = append(heap.participants, participant)
	heap.siftUp(len(heap.participants) - 1)
}

func (heap *Heap) siftUp(idx int) {
	if idx == 0 {
		return
	}

	parentIndex := (idx - 1) / 2

	if heap.participants[parentIndex].Less(heap.participants[idx]) {
		heap.participants[parentIndex], heap.participants[idx] = heap.participants[idx], heap.participants[parentIndex]
		heap.siftUp(parentIndex)
	}
}

func (heap *Heap) pop() *Participant {
	result := heap.participants[0]
	heap.participants[0] = heap.participants[len(heap.participants)-1]
	heap.participants = heap.participants[0 : len(heap.participants)-1]
	heap.siftDown(0)
	return result
}

func (heap *Heap) siftDown(idx int) {
	left := idx*2 + 1
	right := idx*2 + 2

	if left >= len(heap.participants) {
		return
	}

	idxLargest := left

	if right < len(heap.participants) && heap.participants[left].Less(heap.participants[right]) {
		idxLargest = right
	}

	if heap.participants[idx].Less(heap.participants[idxLargest]) {
		heap.participants[idx], heap.participants[idxLargest] = heap.participants[idxLargest], heap.participants[idx]
		heap.siftDown(idxLargest)
	}
}

type Participant struct {
	login   string
	points  int
	penalty int
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
