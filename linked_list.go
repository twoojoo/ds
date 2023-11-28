package ds

import "fmt"

type LlNode[T any] struct {
	prev *LlNode[T]
	next *LlNode[T]
	val  T
}

func (n *LlNode[T]) Value() T {
	return n.val
}

func (n *LlNode[T]) Next() *LlNode[T] {
	return n.next
}

func (n *LlNode[T]) Prev() *LlNode[T] {
	return n.prev
}

type LinkedList[T any] struct {
	length uint
	head   *LlNode[T]
	tail   *LlNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func NewLinkedListFromSlice[T any](s []T) *LinkedList[T] {
	ll := &LinkedList[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}

	ll.Append(s)

	return ll
}

func (ll LinkedList[T]) Length() uint {
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

func (ll *LinkedList[T]) Push(v T) {
	if ll.length == 0 {
		new := &LlNode[T]{
			prev: nil,
			next: nil,
			val:  v,
		}

		ll.tail = new
		ll.head = new

		ll.length++

		return
	}

	new := &LlNode[T]{
		prev: ll.tail,
		val:  v,
	}

	ll.tail.next = new

	ll.tail = new

	ll.length++
}

func (ll *LinkedList[T]) Append(s []T) {
	for i := range s {
		ll.Push(s[i])
	}
}

func (ll *LinkedList[T]) Pop() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	if ll.length == 1 {
		val := ll.tail.val
		ll.Flush()
		return val, true
	}

	result := ll.tail.val

	ll.tail = ll.tail.prev
	ll.tail.next = nil
	ll.length--

	return result, true
}

func (ll *LinkedList[T]) Unshift(v T) {
	new := &LlNode[T]{
		prev: ll.head,
		val:  v,
	}

	if ll.length == 0 {
		ll.head = new
		ll.tail = new
	} else {
		ll.head.prev = new
		ll.head = new
	}

	ll.length++
}

func (ll *LinkedList[T]) Shift() (T, bool) {
	if ll.length == 0 {
		var zero T
		return zero, false
	}

	if ll.length == 1 {
		val := ll.tail.val
		ll.Flush()
		return val, true
	}

	result := ll.head.val

	ll.head = ll.head.next
	ll.head.prev = nil
	ll.length--

	return result, true
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
		new := &LlNode[T]{
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
		new := &LlNode[T]{
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

	new := &LlNode[T]{
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

func (ll *LinkedList[T]) Traverse(action func(v T)) {
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

func (ll *LinkedList[T]) Flush() {
	ll.tail = nil
	ll.head = nil
	ll.length = 0
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
