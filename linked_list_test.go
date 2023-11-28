package ds

import (
	"slices"
	"testing"
)

func TestHead(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Push(8)
	ll.Push(1)
	ll.Push(4)

	if r, ok := ll.Head(); !ok || r != 8 {
		t.Fatal(ok, r)
	}

	ll1 := NewLinkedList[int]()
	ll1.Push(8)

	if r, ok := ll1.Head(); !ok || r != 8 {
		t.Fatal(ok, r)
	}

	ll2 := NewLinkedList[int]()
	if _, ok := ll2.Head(); ok {
		t.Fatal()
	}
}

func TestTail(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4)

	if r, ok := ll.Tail(); !ok || r != 4 {
		t.Fatal(ok, r)
	}

	ll1 := NewLinkedList[int]()
	ll1.Push(8)

	if r, ok := ll1.Tail(); !ok || r != 8 {
		t.Fatal(ok, r)
	}

	ll2 := NewLinkedList[int]()
	if _, ok := ll2.Tail(); ok {
		t.Fatal()
	}
}

func TestPush(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if ll.head.val != 8 {
		t.Fail()
	}

	if ll.tail.val != 7 {
		t.Fail()
	}

	if ll.length != 5 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if r, _ := ll.Pop(); ll.tail.val != 2 || r != 7 {
		t.Fatal(ll.tail.val, r)
	}

	if r, _ := ll.Pop(); ll.tail.val != 4 || r != 2 {
		t.Fatal(ll.tail.val, r)
	}

	ll.Pop()
	if r, _ := ll.Pop(); ll.length != 1 || r != 1 {
		t.Fatal(ll.length, r)
	}

	if r, _ := ll.Pop(); ll.length != 0 || r != 8 {
		t.Fatal(ll.length, r)
	}

	if _, ok := ll.Pop(); ok {
		t.Fatal(ok)
	}
}

func TestUnshift(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Unshift(8)
	ll.Unshift(1)
	ll.Unshift(4)
	ll.Unshift(2)
	ll.Unshift(7)

	if ll.tail.val != 8 {
		t.Fail()
	}

	if ll.head.val != 7 {
		t.Fail()
	}

	if ll.length != 5 {
		t.Fail()
	}
}

func TestShift(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if r, ok := ll.Shift(); !ok || ll.head.val != 1 || r != 8 {
		t.Fatal(ok, ll.head.val, r)
	}

	if r, ok := ll.Shift(); !ok || ll.head.val != 4 || r != 1 {
		t.Fatal(ok, ll.head.val, r)
	}

	_, _ = ll.Shift()
	if r, ok := ll.Shift(); !ok || ll.length != 1 || r != 2 {
		t.Fatal(ok, ll.length, r)
	}

	if r, ok := ll.Shift(); !ok || ll.length != 0 || r != 7 {
		t.Fatal(ok, ll.length, r)
	}

	if _, ok := ll.Pop(); ok {
		t.Fatal(ok)
	}
}

func TestValueAt(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if r, e := ll.ValueAt(2); e != nil || r != 4 {
		t.Fatal(e, r)
	}

	if r, e := ll.ValueAt(4); e != nil || r != 7 {
		t.Fatal(e, r)
	}

	if r, e := ll.ValueAt(0); e != nil || r != 8 {
		t.Fatal(e, r)
	}

	if r, e := ll.ValueAt(1); e != nil || r != 1 {
		t.Fatal(e, r)
	}

	if r, e := ll.ValueAt(3); e != nil || r != 2 {
		t.Fatal(e, r)
	}

	if _, e := ll.ValueAt(5); e == nil {
		t.Fatal()
	}
}

func TestInsertAt(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if e := ll.InsertAt(1, 10); e != nil || ll.length != 6 {
		t.Fatal(e, ll.length)
	}

	if e := ll.InsertAt(0, 10); e != nil || ll.length != 7 {
		t.Fatal(e, ll.length)
	}

	if e := ll.InsertAt(6, 10); e != nil || ll.length != 8 {
		t.Fatal(e, ll.length)
	}

	if e := ll.InsertAt(8, 10); e == nil {
		t.Fatal()
	}
}

func TestFind(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if v, i, ok := ll.Find(func(v int) bool { return v == 2 }); v != 2 || i != 3 || !ok {
		t.Fatal(v, i, ok)
	}

	if v, i, ok := ll.Find(func(v int) bool { return v == 8 }); v != 8 || i != 0 || !ok {
		t.Fatal(v, i, ok)
	}

	if v, i, ok := ll.Find(func(v int) bool { return v == 7 }); v != 7 || i != 4 || !ok {
		t.Fatal(v, i, ok)
	}

	if _, _, ok := ll.Find(func(v int) bool { return v == 100 }); ok {
		t.Fatal()
	}
}

func TestToSlice(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if s := ll.ToSlice(); len(s) != 5 || !slices.Equal(s, []int{8, 1, 4, 2, 7}) {
		t.Fatal(len(s))
	}

	ll1 := NewLinkedList[int]()
	if s := ll1.ToSlice(); len(s) != 0 || !slices.Equal(s, []int{}) {
		t.Fatal(len(s))
	}
}

func TestTraverse(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	var sum int
	if ll.Traverse(func(v int) { sum += v }); sum != 8+1+4+2+7 {
		t.Fatal(sum)
	}

	ll1 := NewLinkedList[int]()
	sum = 0
	if ll1.Traverse(func(v int) { sum += v }); sum != 0 {
		t.Fatal(sum)
	}
}

func TestAppend(t *testing.T) {
	ll := NewLinkedList[int](8, 1, 4, 2, 7)

	if ll.Append([]int{1, 2, 3}); ll.length != 8 || ll.tail.val != 3 {
		t.Fatal(ll.length, ll.tail.val)
	}

	ll.Flush()

	if ll.Append([]int{1, 2, 3}); ll.length != 3 || ll.tail.val != 3 || ll.head.val != 1 {
		t.Fatal(ll.length, ll.tail.val, ll.head.val)
	}

}
