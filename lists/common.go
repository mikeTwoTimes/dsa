package lists

type singlyNode[T any] struct {
	value T
	next  *singlyNode[T]
}
