package ds

import (
	"fmt"
)

type SinglyLinkedList[T any] struct {
	*singlyLinkedListBase[T]
}

func NewSinglyLinkedList[T any](vals ...T) *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{newSinglyLinkedListBaseFromSlice[T](vals)}
}

func (ll SinglyLinkedList[T]) Size() uint {
	return ll.length
}

func (ll SinglyLinkedList[T]) Head() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	return ll.head.val, true
}

func (ll SinglyLinkedList[T]) Tail() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	return ll.tail.val, true
}

func (ll SinglyLinkedList[T]) Push(v T) {
	ll.push(v)
}

func (ll SinglyLinkedList[T]) Unshift(v T) {
	ll.unshift(v)
}

func (ll SinglyLinkedList[T]) Pop() (T, bool) {
	return ll.pop()
}

func (ll SinglyLinkedList[T]) Shift() (T, bool) {
	return ll.shift()
}

func (ll *SinglyLinkedList[T]) Append(s []T) {
	for i := range s {
		ll.push(s[i])
	}
}

func (ll SinglyLinkedList[T]) Flush() {
	ll.flush()
}

func (ll SinglyLinkedList[T]) Traverse(matcher func(v T)) {
	ll.traverse(matcher)
}

func (ll *SinglyLinkedList[T]) ValueAt(idx uint) (T, error) {
	if idx >= ll.length {
		var zero T
		return zero, fmt.Errorf("cannot iterate beyond list end")
	}

	curr := ll.head
	for i := uint(0); i != idx; i++ {
		curr = curr.next
	}

	return curr.val, nil
}

func (ll *SinglyLinkedList[T]) InsertAt(idx uint, v T) error {
	if idx >= ll.length {
		return fmt.Errorf("cannot insert beyond list end")
	}

	if idx < 0 {
		return fmt.Errorf("cannot insert below index 0")
	}

	if idx == 0 {
		new := &sllNode[T]{
			next: ll.head,
			val:  v,
		}

		ll.head = new
		ll.length++

		return nil
	}

	if idx == ll.length-1 {
		new := &sllNode[T]{
			next: nil,
			val:  v,
		}

		ll.tail.next = new
		ll.tail = new
		ll.length++

		return nil
	}

	curr := ll.head
	for i := uint(0); i == idx-1; i++ {
		curr = curr.next
	}

	temp := curr.next
	curr.next = &sllNode[T]{
		next: temp,
		val:  v,
	}

	ll.length++

	return nil
}

func (ll *SinglyLinkedList[T]) Find(matcher func(v T) bool) (T, uint, bool) {
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

// func (ll *LinkedList[T])

func (ll *SinglyLinkedList[T]) ToSlice() []T {
	result := make([]T, ll.length)

	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		result[i] = curr.val
		curr = curr.next
	}

	return result
}

func (ll *SinglyLinkedList[T]) Reverse() {
	ll.reverse()
}

// ToSlice + Flush
func (ll *SinglyLinkedList[T]) Close() []T {
	s := ll.ToSlice()
	ll.Flush()
	return s
}

// Traverse + Flush
func (ll *SinglyLinkedList[T]) Consume(action func(v T)) {
	ll.Traverse(action)
	ll.Flush()
}
