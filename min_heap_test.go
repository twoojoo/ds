package ds

import (
	"testing"
)

func TestMinHepPush(t *testing.T) {
	mh := NewMinHeap[int](func(a, b int) int {
		if a > b {
			return 1
		}

		if b > a {
			return -1
		}

		return 0
	})

	mh.Push(12)
	mh.Push(44)
	mh.Push(112)
	mh.Push(10)
	mh.Push(12)
	mh.Push(9)

	if arr := mh.getArray(); arr[0] != 112 {
		t.Fatal(arr)
	}
}

func TestMinHepPop(t *testing.T) {
	mh := NewMinHeap[int](func(a, b int) int {
		if a > b {
			return 1
		}

		if b > a {
			return -1
		}

		return 0
	})

	mh.Push(12)
	mh.Push(44)
	mh.Push(112)
	mh.Push(10)
	mh.Push(12)
	mh.Push(9)

	mh.Pop()
	mh.Pop()
	mh.Pop()
	mh.Pop()
	mh.Pop()
	mh.Pop()

	if arr := mh.getArray(); len(arr) != 0 {
		t.Fatal(arr)
	}
}
