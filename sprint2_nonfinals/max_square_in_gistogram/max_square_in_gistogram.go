package max_square_in_gistogram

func getMaxSquare(gistogram []int) int {
	// leftLimits - массив, в котором на i позиции
	// находится число, которое обозначает длину максимально возможного прямоугольника, который
	// продолжается влево от столбца
	leftLimits := getLeftLimits(gistogram)
	// rightLimits - массив, в котором на i позиции
	// находится число, которое обозначает длину максимально возможного прямоугольника, который
	// продолжается вправо от столбца
	rightLimits := getRightLimits(gistogram)

	maxSquare := 0
	for i := 0; i < len(gistogram); i++ {
		// площадь прямоугольника с высотой gistogram[i] = сумма площадей прямоугольников, продолжающихся
		// влево и вправо из этой позиции
		s := gistogram[i] * (rightLimits[i] - leftLimits[i] - 1)
		if s > maxSquare {
			maxSquare = s
		}
	}
	return maxSquare
}

func getLeftLimits(gistogram []int) []int {
	n := len(gistogram)
	leftLimits := make([]int, n)
	var stack *ListNode

	//двигаемся вправо от каждого столбца с начала гистограмы
	for i := 0; i < n; i++ {
		//идем по стеку назад
		for stack != nil {
			// упираемся в столбец на стеке - записываем индекс, до которого дошел прямоугольник
			leftLimits[i] = stack.idx
			// если высота столбца на стеке меньше, чем высота текущего столбца, то останавливаемся - мы нашли
			// настоящий столбец, в который упирается прямоугольник
			if stack.h < gistogram[i] {
				break
			}
			// если высота столбца на стеке больше, то мы можем продолжить прямоугольник далее
			// убираем со стека столбец. следующий на стеке будет столбец, в который упирается последний столбец на стеке
			stack = stack.prev
		}
		// если стек пустой, то мы дошли до начала гистограмы - записываем как границу текущего столбца -1
		// это значит, что прямоугольник, продолженный из этого столбца можно продолжить до начала гистограмы
		if stack == nil {
			leftLimits[i] = -1
		}
		// добавляем на стек текущий столбец
		stack = &ListNode{
			h:    gistogram[i],
			idx:  i,
			prev: stack,
		}

	}
	return leftLimits
}

func getRightLimits(gistogram []int) []int {
	n := len(gistogram)
	rightLimits := make([]int, n)
	var stack *ListNode

	//двигаемся влево от каждого столбца с конца гистограмы
	for i := n - 1; i >= 0; i-- {
		//идем по стеку назад
		for stack != nil {
			// упираемся в столбец на стеке - записываем индекс, до которого дошел прямоугольник
			rightLimits[i] = stack.idx
			// если высота столбца на стеке меньше, чем высота текущего столбца, то останавливаемся - мы нашли
			// настоящий столбец, в который упирается прямоугольник
			if stack.h < gistogram[i] {
				break
			}
			// если высота столбца на стеке больше, то мы можем продолжить прямоугольник далее
			// убираем со стека столбец. следующий на стеке будет столбец, в который упирается последний столбец на стеке
			stack = stack.prev
		}
		// если стек пустой, то мы дошли до конца гистограмы - записываем как границу текущего столбца n
		// это значит, что прямоугольник, продолженный из этого столбца можно продолжить до конца гистограмы
		if stack == nil {
			rightLimits[i] = n
		}
		// добавляем на стек текущий столбец
		stack = &ListNode{
			h:    gistogram[i],
			idx:  i,
			prev: stack,
		}

	}
	return rightLimits
}

type ListNode struct {
	h    int
	idx  int
	prev *ListNode
}
