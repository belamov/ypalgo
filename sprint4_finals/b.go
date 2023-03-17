package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/24414/problems/B/
// последнее ОК решение - https://contest.yandex.ru/contest/24414/run-report/84163918/

//Тимофей, как хороший руководитель, хранит информацию о зарплатах своих
//сотрудников в базе данных и постоянно её обновляет.
//Он поручил вам написать реализацию хеш-таблицы, чтобы хранить
//в ней базу данных с зарплатами сотрудников.
//
//Хеш-таблица должна поддерживать следующие операции:
//
//put key value —– добавление пары ключ-значение.
// Если заданный ключ уже есть в таблице, то соответствующее ему значение обновляется.
//get key –— получение значения по ключу.
// Если ключа нет в таблице, то вывести «None».
// Иначе вывести найденное значение.
//delete key –— удаление ключа из таблицы.
// Если такого ключа нет, то вывести «None», иначе вывести
// хранимое по данному ключу значение и удалить ключ.
//В таблице хранятся уникальные ключи.
//
//Требования к реализации:
//
//Нельзя использовать имеющиеся в языках программирования реализации хеш-таблиц (std::unordered_map в С++, dict в Python, HashMap в Java, и т. д.)
//Разрешать коллизии следует с помощью метода цепочек или с помощью открытой адресации.
//Все операции должны выполняться за O(1) в среднем.
//Поддерживать рехеширование и масштабирование хеш-таблицы не требуется.
//Ключи и значения, id сотрудников и их зарплата, —– целые числа. Поддерживать произвольные хешируемые типы не требуется.

//пояснения к реализации: поскольку масштабирование не требуется, создадим сразу много бакетов, чтобы хватило для тестов
//в качестве размера было выбрано 9371 как большое простое число.
//для решения коллизий используется метод цепочек - в бакетах хранится массив из элементов, куда складываются
//элементы с одинаковыми номерами бакетов. в такой реализации мы тратим чуть больше времени на получение и удаление элементов,
//но поскольку у нас по условию количество ключей в хещ-таблице не превышает 10^5, то в каждом бакете будет
//храниться максимум 10^5 / 9371 = 11 элементов. поэтому поиск и удаление будет занимать О(1)+О(10) на пробег по всем элементам
//в бакете, что достаточно для прохождения тестов
//номер бакета (хеш ключа) определяется математическим делением по модулю ключа на размер хеш-таблицы
type Element struct {
	key   int
	value int
}

type HashTable struct {
	buckets [][]*Element
}

func newHashTable(m int) *HashTable {
	return &HashTable{
		buckets: make([][]*Element, m),
	}
}

func (q *HashTable) put(key int, value int) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пустая, то просто добавляем элемент в массив
	if q.buckets[bucketNumber] == nil {
		q.buckets[bucketNumber] = []*Element{{
			key:   key,
			value: value,
		}}
		return
	}

	// если в корзине уже есть элементы, то пройдемся по всем, чтобы
	// заменить значение элемента с таким же ключом
	for i, element := range q.buckets[bucketNumber] {
		if element != nil && element.key == key {
			q.buckets[bucketNumber][i].value = value
			return
		}
	}

	// тут мы убедились, что ключ элемента уникален - можем добавить его в массив
	q.buckets[bucketNumber] = append(q.buckets[bucketNumber], &Element{
		key:   key,
		value: value,
	})
}

func (q *HashTable) getBucketNumber(key int) int {
	// хеширование ключа - простое деление по модулю, поскольку ключи у нас - целые числа
	return mathematicalModulus(key, len(q.buckets))
}

func (q *HashTable) get(key int, writer *bufio.Writer) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пуста - элемента нет
	if q.buckets[bucketNumber] == nil {
		writer.WriteString("None\n")
		return
	}

	// пройдемся по каждому элементу в корзине и найдем среди них элемент с нужным ключом
	for _, element := range q.buckets[bucketNumber] {
		if element.key == key {
			writer.WriteString(strconv.Itoa(element.value))
			writer.WriteString("\n")
			return
		}
	}

	// среди элементов в корзине ключа не найдено
	writer.WriteString("None\n")
}

func (q *HashTable) delete(key int, writer *bufio.Writer) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пуста - элемента нет
	if q.buckets[bucketNumber] == nil {
		writer.WriteString("None\n")
		return
	}

	// пройдемся по каждому элементу в корзине и найдем среди них элемент с нужным ключом
	for i, element := range q.buckets[bucketNumber] {
		if element.key == key {
			writer.WriteString(strconv.Itoa(element.value))
			writer.WriteString("\n")

			//удалим элемент из массива - сместим все элементы справа от него на один влево,
			//а затем уменьшим размер массива на один
			for j := i; j < len(q.buckets[bucketNumber])-1; j++ {
				q.buckets[bucketNumber][j] = q.buckets[bucketNumber][j+1]
			}
			q.buckets[bucketNumber] = q.buckets[bucketNumber][:len(q.buckets[bucketNumber])-1]
			return
		}
	}

	// элемента с заданным ключем не найдено
	writer.WriteString("None\n")
}

func main() {
	scanner := makeScanner()
	commandsCount, _ := strconv.Atoi(readLine(scanner))

	m := 9371
	hashTable := newHashTable(m)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < commandsCount; i++ {
		runCommand(scanner, hashTable, writer)
	}
	writer.Flush()

}
func mathematicalModulus(d, m int) int {
	var res = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func runCommand(scanner *bufio.Scanner, hashTable *HashTable, writer *bufio.Writer) {
	command := readLine(scanner)

	if strings.Contains(command, "get") {
		key, _ := strconv.Atoi(command[4:])
		hashTable.get(key, writer)
		return
	}

	if strings.Contains(command, "put") {
		keyValue := strings.Split(command[4:], " ")
		key, _ := strconv.Atoi(keyValue[0])
		value, _ := strconv.Atoi(keyValue[1])
		hashTable.put(key, value)
		return
	}

	if strings.Contains(command, "delete") {
		key, _ := strconv.Atoi(command[7:])
		hashTable.delete(key, writer)
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
