package ds

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

type linkedListBase[T any] struct {
	length uint
	head   *LlNode[T]
	tail   *LlNode[T]
}

func newLinkedListBase[T any]() *linkedListBase[T] {
	return &linkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func newLinkedListBaseFromSlice[T any](s []T) *linkedListBase[T] {
	ll := &linkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}

	for i := range s {
		ll.push(s[i])
	}

	return ll
}

func (ll *linkedListBase[T]) push(v T) {
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

func (ll *linkedListBase[T]) pop() (T, bool) {
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

	ll.tail = ll.tail.prev
	ll.tail.next = nil
	ll.length--

	return result, true
}

func (ll *linkedListBase[T]) unshift(v T) {
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

func (ll *linkedListBase[T]) shift() (T, bool) {
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
	ll.head.prev = nil
	ll.length--

	return result, true
}

func (ll *linkedListBase[T]) flush() {
	ll.tail = nil
	ll.head = nil
	ll.length = 0
}
