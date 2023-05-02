package main

func siftDown(heap []int, idx int) int {
	left := idx * 2
	right := idx*2 + 1

	if left >= len(heap) {
		return idx
	}

	idxLargest := left

	if right < len(heap) && heap[left] < heap[right] {
		idxLargest = right
	}

	if heap[idx] < heap[idxLargest] {
		heap[idx], heap[idxLargest] = heap[idxLargest], heap[idx]
		return siftDown(heap, idxLargest)
	}

	return idx
}

func test() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		panic("WA")
	}
}
