package lists

type Stack[T any] struct {
	len  int
	head *singlyNode[T]
}

func (s *Stack[T]) Push(item T) {
	node := &singlyNode[T]{value: item}
	defer func() { s.len++ }()

	if s.len == 0 {
		s.head = node
		return
	}

	node.next = s.head
	s.head = node
}

func (s *Stack[T]) Pop() {
	if s.len <= 0 {
		panic("stack is empty")
	}

	s.len--
	s.head = s.head.next
}

func (s *Stack[T]) Peek() T {
	if s.len <= 0 {
		panic("stack is empty")
	}

	return s.head.value
}

func (s *Stack[T]) Len() int {
	return s.len
}
