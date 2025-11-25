package lists

type Queue[T any] struct {
	len  int
	head *singlyNode[T]
	tail *singlyNode[T]
}

func (q *Queue[T]) Enqueue(item T) {
	node := &singlyNode[T]{value: item}
	defer func() { q.len++ }()

	if q.len == 0 {
		q.tail = node
		q.head = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() {
	if q.len <= 0 {
		panic("queue is empty")
	}

	q.len--
	q.head = q.head.next
}

func (q *Queue[T]) Peek() T {
	if q.len <= 0 {
		panic("queue is empty")
	}

	return q.head.value
}

func (q *Queue[T]) Len() int {
	return q.len
}
