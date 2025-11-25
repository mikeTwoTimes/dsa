package lists

type Queue[T any] interface {
	Enqueue(item T)
	Deque()
	Peek() T
	Len() int
}

type queue[T any] struct {
	len  int
	head *singlyNode[T]
	tail *singlyNode[T]
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(item T) {
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

func (q *queue[T]) Deque() {
	if q.len <= 0 {
		panic("queue is empty")
	}

	q.len--
	q.head = q.head.next
}

func (q *queue[T]) Peek() T {
	if q.len <= 0 {
		panic("queue is empty")
	}

	return q.head.value
}

func (q *queue[T]) Len() int {
	return q.len
}
