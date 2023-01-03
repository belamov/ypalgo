package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deque - реализация дека на основе кольцевого буфера
// по сути реализация представляет собой два стека:
// один отвечает за начало очереди, добавляет и берет элементы по индексу head - головной стек
// второй отвечает за конец очереди, добавляет и берет элементы по индексу tail - хвостовой стек
// причем головной и хвостовой стеки растут в разных направлениях - head увеличивается, а tail уменьшается
//
// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// все операции происходят за О(1), поскольку мы храним индексы началов стеков, а значит просто двигаем их
// простыми арифметическими операциями
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
func (d *Deque) pushFront(x int) {
	if d.currentSize == d.maxSize {
		fmt.Println("error")
		return
	}
	d.queue[d.head] = x
	d.head = (d.head + 1) % d.maxSize
	d.currentSize++
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
func (d *Deque) pushBack(x int) {
	if d.currentSize == d.maxSize {
		fmt.Println("error")
		return
	}
	d.tail = (d.tail - 1 + d.maxSize) % d.maxSize
	d.queue[d.tail] = x

	d.currentSize++
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
func (d *Deque) popFront() {
	if d.currentSize == 0 {
		fmt.Println("error")
		return
	}
	d.head = (d.head - 1 + d.maxSize) % d.maxSize
	fmt.Println(d.queue[d.head])
	d.currentSize--
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
func (d *Deque) popBack() {
	if d.currentSize == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(d.queue[d.tail])
	d.tail = (d.tail + 1) % d.maxSize
	d.currentSize--
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
		queue.pushFront(arg)
		return
	}
	if strings.Contains(command, "push_back") {
		arg, _ := strconv.Atoi(command[10:])
		queue.pushBack(arg)
		return
	}

	if command == "pop_front" {
		queue.popFront()
		return
	}

	if command == "pop_back" {
		queue.popBack()
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
