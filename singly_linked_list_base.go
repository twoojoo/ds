package ds

import "fmt"

type sllNode[T any] struct {
	next *sllNode[T]
	val  T
}

type singlyLinkedListBase[T any] struct {
	length uint
	head   *sllNode[T]
	tail   *sllNode[T]
}

func newSinglyLinkedListBase[T any]() *singlyLinkedListBase[T] {
	return &singlyLinkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func newSinglyLinkedListBaseFromSlice[T any](s []T) *singlyLinkedListBase[T] {
	ll := &singlyLinkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}

	for i := range s {
		ll.push(s[i])
	}

	return ll
}

func (ll *singlyLinkedListBase[T]) push(v T) {
	if ll.length == 0 {
		new := &sllNode[T]{
			next: nil,
			val:  v,
		}

		ll.tail = new
		ll.head = new

		ll.length++

		return
	}

	new := &sllNode[T]{
		val: v,
	}

	ll.tail.next = new

	ll.tail = new

	ll.length++
}

func (ll *singlyLinkedListBase[T]) nodeAtIndex(idx uint) (*sllNode[T], error) {
	curr := ll.head

	if idx >= ll.length {
		return nil, fmt.Errorf("index %v out of length", idx)
	}

	var i uint
	for ; i < idx; i++ {
		curr = curr.next
	}

	return curr, nil
}

func (ll *singlyLinkedListBase[T]) pop() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	if ll.length == 1 {
		val := ll.tail.val
		ll.flush()
		return val, true
	}

	result := ll.tail.val

	ll.length--

	newTail, _ := ll.nodeAtIndex(ll.length - 1)
	ll.tail = newTail
	ll.tail.next = nil

	return result, true
}

func (ll *singlyLinkedListBase[T]) unshift(v T) {
	new := &sllNode[T]{
		val: v,
	}

	if ll.length == 0 {
		ll.head = new
		ll.tail = new
	} else {
		ll.head = new
	}

	ll.length++
}

func (ll *singlyLinkedListBase[T]) shift() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	if ll.length == 1 {
		val := ll.tail.val
		ll.flush()
		return val, true
	}

	result := ll.head.val

	ll.head = ll.head.next
	ll.length--

	return result, true
}

func (ll *singlyLinkedListBase[T]) flush() {
	ll.tail = nil
	ll.head = nil
	ll.length = 0
}

func (ll *singlyLinkedListBase[T]) traverse(action func(v T)) {
	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		action(curr.val)
		curr = curr.next
	}
}
