package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22781/run-report/80434035/ - ссылка на последнее ОК решение

// https://contest.yandex.ru/contest/22781/problems/A/

// Гоша реализовал структуру данных Дек, максимальный размер которого определяется заданным числом.
// Методы push_back(x), push_front(x), pop_back(), pop_front() работали корректно.
// Но, если в деке было много элементов, программа работала очень долго.
// Дело в том, что не все операции выполнялись за O(1).
// Помогите Гоше! Напишите эффективную реализацию.
//
// Внимание: при реализации используйте кольцевой буфер.

// Deque - реализация дека на основе кольцевого буфера
// по сути реализация представляет собой два стека:
// один отвечает за начало очереди, добавляет и берет элементы по индексу head - головной стек
// второй отвечает за конец очереди, добавляет и берет элементы по индексу tail - хвостовой стек
// причем головной и хвостовой стеки растут в разных направлениях - head увеличивается, а tail уменьшается
//
// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// каждая отдельная операция над деком происходит за О(1), поскольку мы храним индексы началов стеков, а значит просто двигаем их
// простыми арифметическими операциями
// поэтому общая временная сложность программы будет O(n), где n - количество операций в инпуте
//
// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// если максимальный размер стека = n, то дек будет хранить слайс размером n
// также мы храним указатели на вершины стеков и текущий размер, но они не зависят от входных данных
// поэтому дек будет потреблять O(n) памяти
type Deque struct {
	queue       []int
	maxSize     int // макисмально допустимый размер дека
	head        int // указатель на свободный индекс в головном стеке
	tail        int // указатель на последний добавленный элемент в хвостовой стек
	currentSize int // количество элементов в деке
}

// newDeque создает новый дек
func newDeque(maxSize int) *Deque {
	return &Deque{
		queue:   make([]int, maxSize),
		maxSize: maxSize,
		head:    0,
		tail:    0,
	}
}

// pushFront добавляет элемент в начало дека, кладет элемент наверх головного стека
// сначала определяем, можем ли мы добавить элемент в дек (не достигнут ли максимальный размер дека)
// если не можем, выходим из функции с ошибкой
//
// в head хранится указатель на свободный индекс в головном стеке, а значит можно
// просто записать в него элемент.
//
// после увеличиваем head, причем ограничиваем его по модулю максимального
// размера дека, это позволяет не выйти за пределы слайса.
//
// теперь в head опять хранится указатель на свободный индекс в головном стеке
//
// также увеличиваем текущий размер дека currentSize
func (d *Deque) pushFront(x int) error {
	if d.currentSize == d.maxSize {
		return errors.New("превышен максимальный размер дека")
	}
	d.queue[d.head] = x
	d.head = (d.head + 1) % d.maxSize
	d.currentSize++
	return nil
}

// pushBack добавляет элемент в конец дека, кладет элемент наверх хвостового стека
// сначала определяем, можем ли мы добавить элемент в дек (не достигнут ли максимальный размер дека)
// если не можем, выходим из функции с ошибкой
//
// в tail хранится указатель на последний добавленный элемент в хвостовой стек
//
// сначала в tail записываем свободный индекс в хвостовом стеке, причем ограничиваем его по модулю максимального
// размера дека, это позволяет не выйти за пределы слайса
//
// далее добавляем элемент на вершину хвостового стека
//
// теперь в tail опять хранится указатель на последний добавленный элемент в хвостовой стек
//
// также увеличиваем текущий размер дека currentSize
func (d *Deque) pushBack(x int) error {
	if d.currentSize == d.maxSize {
		return errors.New("превышен максимальный размер дека")
	}
	d.tail = (d.tail - 1 + d.maxSize) % d.maxSize
	d.queue[d.tail] = x
	d.currentSize++
	return nil
}

// popFront берет элемент с начала дека, берет элемент с вершины головного стека
//
// в head хранится указатель на свободный индекс в головном стеке
//
// сначала запишем в head указатель на последний добавленый эллемент - просто задекрементим индекс, причем
// закольцовывая индекс
//
// возвращаем элемент по вычисленному индексу
// теперь считаем, что в этом индексе нет элемента.
// при добавлении элемента в этот индекс значение просто перезапишется
//
// теперь в head опять хранится указатель на свободный индекс в головном стеке
//
// также уменьшаем текущий размер дека currentSize
func (d *Deque) popFront() (int, error) {
	if d.currentSize == 0 {
		return 0, errors.New("нет элементов в деке")
	}
	d.head = (d.head - 1 + d.maxSize) % d.maxSize
	d.currentSize--
	return d.queue[d.head], nil
}

// popBack берет элемент с конца дека, берет элемент с вершины хвостового стека
//
// в tail хранится указатель на последний добавленный элемент в хвостовой стек
//
// можем просто вывести элемент по этому индексу
//
// после заинкрементим tail, причем закольцовывая его
//
// теперь в tail опять хранится указатель на последний добавленный элемент в хвостовой стек
//
// также уменьшаем текущий размер дека currentSize
func (d *Deque) popBack() (int, error) {
	if d.currentSize == 0 {
		return 0, errors.New("нет элементов в деке")
	}
	result := d.queue[d.tail]
	d.tail = (d.tail + 1) % d.maxSize
	d.currentSize--
	return result, nil
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))
	queueSize, _ := strconv.Atoi(readLine(scanner))

	queue := newDeque(queueSize)

	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, queue)
	}
}

func runCommand(scanner *bufio.Scanner, queue *Deque) {
	command := readLine(scanner)

	if strings.Contains(command, "push_front") {
		arg, _ := strconv.Atoi(command[11:])
		err := queue.pushFront(arg)
		if err != nil {
			fmt.Println("error")
		}
		return
	}
	if strings.Contains(command, "push_back") {
		arg, _ := strconv.Atoi(command[10:])
		err := queue.pushBack(arg)
		if err != nil {
			fmt.Println("error")
		}
		return
	}

	if command == "pop_front" {
		result, err := queue.popFront()
		if err != nil {
			fmt.Println("error")
			return
		}
		fmt.Println(result)
		return
	}

	if command == "pop_back" {
		result, err := queue.popBack()
		if err != nil {
			fmt.Println("error")
			return
		}
		fmt.Println(result)
		return
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
