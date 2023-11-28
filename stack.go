package ds

type Stack[T any] struct {
	*LinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{NewLinkedList[T]()}
}

func NewStackFromSlice[T any](s []T) *Stack[T] {
	return &Stack[T]{NewLinkedListFromSlice[T](s)}
}

func (q *Stack[T]) Add(v T) {
	q.Push(v)
}

func (q *Stack[T]) Remove() (T, bool) {
	return q.Pop()
}

func (q *Stack[T]) Size() uint {
	return q.Length()
}

func (q *Stack[T]) IsEmpty() bool {
	return q.Length() == 0
}
