package ds

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(5)
	q.Enqueue(12)
	q.Enqueue(7)
	q.Enqueue(9)

	if q.length != 4 || q.head.val != 5 || q.tail.val != 9 {
		t.Fatal(q.length, q.head.val, q.tail.val)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue(5, 12, 7, 9, 1)

	if r, ok := q.Dequeue(); !ok || r != 5 || q.head.val != 12 || q.length != 4 {
		t.Fatal(ok, r, q.head.val, q.length)
	}

	if r, ok := q.Dequeue(); !ok || r != 12 || q.head.val != 7 || q.length != 3 {
		t.Fatal(ok, r, q.head.val, q.length)
	}

	q.Dequeue()
	if r, ok := q.Dequeue(); !ok || r != 9 || q.head.val != 1 || q.length != 1 {
		t.Fatal(ok, r, q.head.val, q.length)
	}

	q.Dequeue()
	if _, ok := q.Dequeue(); ok {
		t.Fatal(ok)
	}
}

func TestQueueSize(t *testing.T) {
	q := NewQueue(5, 12, 7, 9, 1)

	if s := q.Size(); s != 5 {
		t.Fatal(s)
	}

	q.Dequeue()
	if s := q.Size(); s != 4 {
		t.Fatal(s)
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	if s := q.Size(); s != 1 {
		t.Fatal(s)
	}

	q.Dequeue()
	if s := q.Size(); s != 0 {
		t.Fatal(s)
	}
}

func TestQueueFlush(t *testing.T) {
	q := NewQueue(5, 12, 7, 9, 1)

	q.Flush()
	if s := q.Size(); s != 0 {
		t.Fatal(s)
	}

	q.Flush()
	if s := q.Size(); s != 0 {
		t.Fatal(s)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue(5, 12, 7)

	if q.IsEmpty() {
		t.Fatal()
	}

	q.Flush()
	if !q.IsEmpty() {
		t.Fatal()
	}
}
