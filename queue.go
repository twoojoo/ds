package ds

type Queue[T any] struct {
	*LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{NewLinkedList[T]()}
}

func (q *Queue[T]) Enqueue(v T) {
	q.Push(v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.Shift()
}

func (q *Queue[T]) Size() uint {
	return q.Length()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}
