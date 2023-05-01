package main

func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return idx
	}

	parentIndex := idx / 2

	if heap[parentIndex] < heap[idx] {
		heap[parentIndex], heap[idx] = heap[idx], heap[parentIndex]
		return siftUp(heap, parentIndex)
	}

	return idx
}

func test() {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	if siftUp(sample, 5) != 1 {
		panic("WA")
	}
}
