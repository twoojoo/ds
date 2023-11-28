package ds

type Stack[T any] struct {
	*linkedListBase[T]
}

func NewStack[T any](vals ...T) *Stack[T] {
	return &Stack[T]{newLinkedListBaseFromSlice[T](vals)}
}

func (q *Stack[T]) Add(v T) {
	q.push(v)
}

func (q *Stack[T]) Remove() (T, bool) {
	return q.pop()
}

func (q *Stack[T]) Size() uint {
	return q.length
}

func (q *Stack[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Stack[T]) Flush() {
	q.flush()
}
