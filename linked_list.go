package ds

import (
	"fmt"
)

type LinkedList[T any] struct {
	*linkedListBase[T]
}

func NewLinkedList[T any](vals ...T) *LinkedList[T] {
	return &LinkedList[T]{newLinkedListBaseFromSlice[T](vals)}
}

func (ll LinkedList[T]) Size() uint {
	return ll.length
}

func (ll LinkedList[T]) Head() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	return ll.head.val, true
}

func (ll LinkedList[T]) Tail() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	return ll.tail.val, true
}

func (ll LinkedList[T]) Push(v T) {
	ll.push(v)
}

func (ll LinkedList[T]) Unshift(v T) {
	ll.unshift(v)
}

func (ll LinkedList[T]) Pop() (T, bool) {
	return ll.pop()
}

func (ll LinkedList[T]) Shift() (T, bool) {
	return ll.shift()
}

func (ll *LinkedList[T]) Append(s []T) {
	for i := range s {
		ll.push(s[i])
	}
}

func (ll LinkedList[T]) Flush() {
	ll.flush()
}

func (ll LinkedList[T]) Traverse(matcher func(v T)) {
	ll.traverse(matcher)
}

func (ll *LinkedList[T]) ValueAt(idx uint) (T, error) {
	if idx >= ll.length {
		var zero T
		return zero, fmt.Errorf("cannot iterate beyond list end")
	}

	if idx < 0 {
		var zero T
		return zero, fmt.Errorf("cannot Find below index 0")
	}

	fromTail := idx > ll.length/2

	curr := ll.head
	position := idx
	if fromTail {
		position = ll.length - idx - 1
		curr = ll.tail
	}

	for i := uint(0); i != position; i++ {
		if fromTail {
			curr = curr.prev
		} else {
			curr = curr.next
		}
	}

	return curr.val, nil
}

func (ll *LinkedList[T]) InsertAt(idx uint, v T) error {
	if idx >= ll.length {
		return fmt.Errorf("cannot insert beyond list end")
	}

	if idx < 0 {
		return fmt.Errorf("cannot insert below index 0")
	}

	if idx == 0 {
		new := &llNode[T]{
			prev: nil,
			next: ll.head,
			val:  v,
		}

		ll.head.prev = new
		ll.head = new
		ll.length++

		return nil
	}

	if idx == ll.length-1 {
		new := &llNode[T]{
			prev: ll.tail,
			next: nil,
			val:  v,
		}

		ll.tail.next = new
		ll.tail = new
		ll.length++

		return nil
	}

	fromTail := idx > ll.length/2

	curr := ll.head
	position := idx
	if fromTail {
		position = ll.length - idx - 1
		curr = ll.tail
	}

	for i := uint(0); i != position; i++ {
		if fromTail {
			curr = curr.prev
		} else {
			curr = curr.next
		}
	}

	new := &llNode[T]{
		prev: curr.prev,
		next: curr,
		val:  v,
	}

	curr.prev.next = new
	curr.prev = new

	ll.length++

	return nil
}

func (ll *LinkedList[T]) Find(matcher func(v T) bool) (T, uint, bool) {
	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		if matcher(curr.val) {
			return curr.val, i, true
		}

		curr = curr.next
	}

	var zero T
	return zero, 0, false
}

func (ll *linkedListBase[T]) traverse(action func(v T)) {
	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		action(curr.val)
		curr = curr.next
	}
}

func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, ll.length)

	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		result[i] = curr.val
		curr = curr.next
	}

	return result
}

// ToSlice + Flush
func (ll *LinkedList[T]) Close() []T {
	s := ll.ToSlice()
	ll.Flush()
	return s
}

// Traverse + Flush
func (ll *LinkedList[T]) Consume(action func(v T)) {
	ll.Traverse(action)
	ll.Flush()
}
