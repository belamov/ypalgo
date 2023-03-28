package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/24414/problems/B/
// последнее ОК решение - https://contest.yandex.ru/contest/24414/run-report/84606262/

//Тимофей, как хороший руководитель, хранит информацию о зарплатах своих
//сотрудников в базе данных и постоянно её обновляет.
//Он поручил вам написать реализацию хеш-таблицы, чтобы хранить
//в ней базу данных с зарплатами сотрудников.
//
//Хеш-таблица должна поддерживать следующие операции:
//
//Put key value —– добавление пары ключ-значение.
// Если заданный ключ уже есть в таблице, то соответствующее ему значение обновляется.
//Get key –— получение значения по ключу.
// Если ключа нет в таблице, то вывести «None».
// Иначе вывести найденное значение.
//Delete key –— удаление ключа из таблицы.
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
//элементы с одинаковыми номерами бакетов.
//номер бакета (хеш ключа) определяется математическим делением по модулю ключа на размер хеш-таблицы

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// итоговая сложность O(N), где N - количество команд
// каждая отдельная команда выполняется за O(1):
// поскольку у нас по условию количество ключей в хещ-таблице не превышает 10^5, то в каждом бакете будет
// храниться максимум 10^5 / 9371 = 11 элементов. поэтому поиск, вставка и удаление будет
// занимать О(1) на получение бакета по ключу + О(11) на пробег по всем элементам в бакете
//
// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// программа будет потреблять O(M) памяти, где
// M - количество различных добавляемых элементов
// но при значениях M < 9371, все равно будет выделяться память на 9371 бакет
type Element struct {
	key   int
	value int
}

type HashTable struct {
	buckets [][]Element
}

func newHashTable(m int) *HashTable {
	return &HashTable{
		buckets: make([][]Element, m),
	}
}

func (q *HashTable) Put(key int, value int) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пустая, то просто добавляем элемент в массив
	if q.buckets[bucketNumber] == nil {
		q.buckets[bucketNumber] = []Element{{
			key:   key,
			value: value,
		}}
		return
	}

	// если в корзине уже есть элементы, то пройдемся по всем, чтобы
	// заменить значение элемента с таким же ключом
	for i, element := range q.buckets[bucketNumber] {
		if element.key == key {
			q.buckets[bucketNumber][i].value = value
			return
		}
	}

	// тут мы убедились, что ключ элемента уникален - можем добавить его в массив
	q.buckets[bucketNumber] = append(q.buckets[bucketNumber], Element{
		key:   key,
		value: value,
	})
}

func (q *HashTable) Get(key int) (int, error) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пуста - элемента нет
	if q.buckets[bucketNumber] == nil {
		return 0, errors.New("None")
	}

	// пройдемся по каждому элементу в корзине и найдем среди них элемент с нужным ключом
	for _, element := range q.buckets[bucketNumber] {
		if element.key == key {
			return element.value, nil
		}
	}

	// среди элементов в корзине ключа не найдено
	return 0, errors.New("None")
}

func (q *HashTable) Delete(key int) (int, error) {
	// Вычисляем индекс корзины
	bucketNumber := q.getBucketNumber(key)

	// если корзина пуста - элемента нет
	if q.buckets[bucketNumber] == nil {
		return 0, errors.New("None")
	}

	// пройдемся по каждому элементу в корзине и найдем среди них элемент с нужным ключом
	for i, element := range q.buckets[bucketNumber] {
		if element.key == key {
			deletedValue := element.value

			//удалим элемент из массива - поменяем его местами с последним элементом,
			//а затем уменьшим размер массива на один
			q.buckets[bucketNumber][i], q.buckets[bucketNumber][len(q.buckets[bucketNumber])-1] = q.buckets[bucketNumber][len(q.buckets[bucketNumber])-1], q.buckets[bucketNumber][i]
			q.buckets[bucketNumber] = q.buckets[bucketNumber][:len(q.buckets[bucketNumber])-1]

			return deletedValue, nil
		}
	}

	// элемента с заданным ключем не найдено
	return 0, errors.New("None")
}

func (q *HashTable) getBucketNumber(key int) int {
	// хеширование ключа - простое деление по модулю, поскольку ключи у нас - целые числа
	return mathematicalModulus(key, len(q.buckets))
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
	return (d%m + m) % m
}

func runCommand(scanner *bufio.Scanner, hashTable *HashTable, writer *bufio.Writer) {
	command := readLine(scanner)

	if strings.Contains(command, "get") {
		key, _ := strconv.Atoi(command[4:])
		value, err := hashTable.Get(key)
		if err != nil {
			writer.WriteString(err.Error())
			writer.WriteString("\n")
			return
		}
		writer.WriteString(strconv.Itoa(value))
		writer.WriteString("\n")
		return
	}

	if strings.Contains(command, "put") {
		keyValue := strings.Split(command[4:], " ")
		key, _ := strconv.Atoi(keyValue[0])
		value, _ := strconv.Atoi(keyValue[1])
		hashTable.Put(key, value)
		return
	}

	if strings.Contains(command, "delete") {
		key, _ := strconv.Atoi(command[7:])
		value, err := hashTable.Delete(key)
		if err != nil {
			writer.WriteString(err.Error())
			writer.WriteString("\n")
			return
		}
		writer.WriteString(strconv.Itoa(value))
		writer.WriteString("\n")
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
