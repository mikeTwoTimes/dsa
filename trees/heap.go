package trees

type Heap[T ordered] struct {
	data []T
}

func swap[T ordered](first, second *T) {
	*first, *second = *second, *first
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return (idx * 2) + 1
}

func rightChild(idx int) int {
	return (idx * 2) + 2
}

func (h *Heap[T]) Insert(item T) {
	h.data = append(h.data, item)
	h.heapifyUp(h.Len() - 1)
}

func (h *Heap[T]) Pop() {
	if h.Len() == 0 {
		panic("heap is empty")
	}

	swap(&h.data[h.Len()-1], &h.data[0])
	h.data = h.data[:h.Len()-1]
	h.heapifyDown(0)
}

func (h *Heap[T]) Peek() T {
	return h.data[0]
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := parent(idx)

	if h.data[parentIdx] > h.data[idx] {
		swap(&h.data[parentIdx], &h.data[idx])
		h.heapifyUp(parentIdx)
	}
}

func (h *Heap[T]) heapifyDown(idx int) {
	leftIdx := leftChild(idx)

	if leftIdx >= h.Len() {
		return
	}

	rightIdx := rightChild(idx)
	minIdx := idx

	if h.data[leftIdx] < h.data[minIdx] {
		minIdx = leftIdx
	}

	if rightIdx < h.Len() && h.data[rightIdx] < h.data[minIdx] {
		minIdx = rightIdx
	}

	if minIdx != idx {
		swap(&h.data[idx], &h.data[minIdx])
		h.heapifyDown(minIdx)
	}
}
