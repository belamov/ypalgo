package main

type Node struct {
	value int
	left  *Node
	right *Node
}

// https://contest.yandex.ru/contest/24810/problems/B/
// https://leetcode.com/problems/delete-node-in-a-bst/
// https://contest.yandex.ru/contest/24810/run-report/86835719/ - последнее ОК решение
//Дано бинарное дерево поиска, в котором хранятся ключи.
//Ключи — уникальные целые числа.
//Найдите вершину с заданным ключом и удалите её из дерева так, чтобы дерево осталось
//корректным бинарным деревом поиска.
//Если ключа в дереве нет, то изменять дерево не надо.
//На вход вашей функции подаётся корень дерева и ключ, который надо удалить.
//Функция должна вернуть корень изменённого дерева.
//Сложность удаления узла должна составлять O(h), где h — высота дерева.
//Создавать новые вершины (вдруг очень захочется) нельзя.

// Для удаления сначала необходимо найти удаляемый элемент, спускаясь вниз по дереву.
// У найденного элемента может быть несколько вариантов наличия потомков - без потомков, один потомок, оба потомка
// Для каждого варианта действуем по-своему:
//  - если удаляемый элемент - лист, то просто удаляем его
//  - если удаляемый элемент имеет одного потомка, то меняем родителя этого потомка с удаляемого элемента на родителя удаляемого элемента
//  - если удаляемый элемент имеет и правого, и левого потомка, то заменяем удаляемый элемент
//    минимальным элементом в правом потомке и потом удаляем копию
//    этого минимального элемента в правом потомке рекурсивно (так мы покроем
//    случаи, когда минимальный элемент имеет одного потомка)

// -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
// алгоритм выполняется за O(H), где H - высота дерева.
// при поиске удаляемого элемента мы опускаемся ниже на каждом шаге рекурсии, максимум таких переходов - H
// после нахождения удаляемого элемента (если это не лист) мы можем спуститься до конца дерева, чтобы удалить копию
// в итоге спусков будет не больше, чем H, поэтому оценка будет O(H)

// -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
// мы не создаем никакие структуры, но
// поскольку мы обходим дерево рекурсивно, то на поддержку
// стека рекурсии понадобится O(H) памяти, где H - высота дерева
func remove(node *Node, key int) *Node {
	// если дерево пустое - возвращаем корень
	if node == nil {
		return node
	}

	// если удаляемый элемент больше, чем текущий узел - то продолжим поиск в правом поддереве
	// причем, нам нужно заменить корень правого поддерева у текущего корня, поскольку он может поменяться,
	// если искомый элемент будет корнем правого поддерева
	if key > node.value {
		node.right = remove(node.right, key)
		return node
	}

	// если удаляемый элемент меньше, чем текущий узел - то продолжим поиск в левом поддереве
	// причем, нам нужно заменить корень левого поддерева у текущего корня, поскольку он может поменяться,
	// если искомый элемент будет корнем левого поддерева
	if key < node.value {
		node.left = remove(node.left, key)
		return node
	}

	// тут node - элемент, который необходимо удалить
	// у нас четыре варианта развития событий:

	// первый вариант - элемент является листом
	// в этом случае мы можем просто удалить элемент и дерево останется валидным, поскольку
	// дерево не распадется на части
	if node.left == nil && node.right == nil {
		return nil
	}

	// второй вариант - у искомого элемента есть только правое поддерево
	// тогда мы можем просто назначить потомком текущего элемента правое поддерево
	if node.left == nil {
		node = node.right
		return node
	}

	// третий вариант - у искомого элемента есть только левое поддерево
	// тогда мы можем просто назначить потомком текущего элемента левое поддерево
	if node.right == nil {
		node = node.left
		return node
	}

	// четвертый вариант - у искомого элемента есть и правое, и левое поддерево
	// в этом случае нам нужно заменить искомый элемент подходящим элементом среди потомков
	// подходящим будем считать самый минимальный элемент в правом поддереве
	replacementNode := findMinNode(node.right)
	node.value = replacementNode.value

	// после замены у нас получается копия элемента в правом поддереве
	// причем эта копия не будет иметь правого поддерева
	// получается, что мы можем рекурсивно удалить этот элемент, и алгоритм остановится
	// либо на первом, либо на третьем варианте выше
	// также не забудем переназначить корень правого поддерева, поскольку он может измениться
	node.right = remove(node.right, replacementNode.value)

	return node
}

func findMinNode(node *Node) *Node {
	// спускаемся по левым поддеревьям, пока не дойдем до крайнего элемента
	for {
		if node.left == nil {
			return node
		}
		node = node.left
	}
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		panic("WA")
	}
	if newHead.right != &node5 {
		panic("WA")
	}
	if newHead.right.value != 8 {
		panic("WA")
	}
}
