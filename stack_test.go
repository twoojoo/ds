package ds

import (
	"testing"
)

func TestStackAdd(t *testing.T) {
	s := NewStack[int]()
	s.Add(5)
	s.Add(12)
	s.Add(7)
	s.Add(9)

	if s.length != 4 || s.tail.val != 9 || s.head.val != 5 {
		t.Fatal(s.length, s.head.val, s.tail.val)
	}
}

func TestStackRemove(t *testing.T) {
	s := NewStack(5, 12, 7, 9, 1)

	if r, ok := s.Remove(); !ok || r != 1 || s.tail.val != 9 || s.length != 4 {
		t.Fatal(ok, r, s.tail.val, s.length)
	}

	if r, ok := s.Remove(); !ok || r != 9 || s.tail.val != 7 || s.length != 3 {
		t.Fatal(ok, r, s.tail.val, s.length)
	}

	s.Remove()
	if r, ok := s.Remove(); !ok || r != 12 || s.tail.val != 5 || s.length != 1 {
		t.Fatal(ok, r, s.tail.val, s.length)
	}

	s.Remove()
	if _, ok := s.Remove(); ok {
		t.Fatal(ok)
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack(5, 12, 7, 9, 1)

	if s := s.Size(); s != 5 {
		t.Fatal(s)
	}

	s.Remove()
	if s := s.Size(); s != 4 {
		t.Fatal(s)
	}

	s.Remove()
	s.Remove()
	s.Remove()
	if s := s.Size(); s != 1 {
		t.Fatal(s)
	}

	s.Remove()
	if s := s.Size(); s != 0 {
		t.Fatal(s)
	}
}

func TestStackFlush(t *testing.T) {
	s := NewStack(5, 12, 7)

	s.Flush()
	if s := s.Size(); s != 0 {
		t.Fatal(s)
	}

	s.Flush()
	if s := s.Size(); s != 0 {
		t.Fatal(s)
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := NewStack(5, 12, 7)

	if s.IsEmpty() {
		t.Fatal()
	}

	s.Flush()
	if !s.IsEmpty() {
		t.Fatal()
	}
}
