package ds

type Queue[T any] struct {
	*linkedListBase[T]
}

func NewQueue[T any](vals ...T) *Queue[T] {
	return &Queue[T]{newLinkedListBaseFromSlice[T](vals)}
}

func (q *Queue[T]) Enqueue(v T) {
	q.push(v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.shift()
}

func (q *Queue[T]) Size() uint {
	return q.length
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Flush() {
	q.flush()
}
