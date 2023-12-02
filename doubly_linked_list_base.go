package ds

type dllNode[T any] struct {
	prev *dllNode[T]
	next *dllNode[T]
	val  T
}

type doublyLinkedListBase[T any] struct {
	length uint
	head   *dllNode[T]
	tail   *dllNode[T]
}

func newDoublyLinkedListBase[T any]() *doublyLinkedListBase[T] {
	return &doublyLinkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func newDoublyLinkedListBaseFromSlice[T any](s []T) *doublyLinkedListBase[T] {
	ll := &doublyLinkedListBase[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}

	for i := range s {
		ll.push(s[i])
	}

	return ll
}

func (ll *doublyLinkedListBase[T]) push(v T) {
	if ll.length == 0 {
		new := &dllNode[T]{
			prev: nil,
			next: nil,
			val:  v,
		}

		ll.tail = new
		ll.head = new

		ll.length++

		return
	}

	new := &dllNode[T]{
		prev: ll.tail,
		val:  v,
	}

	ll.tail.next = new

	ll.tail = new

	ll.length++
}

func (ll *doublyLinkedListBase[T]) pop() (T, bool) {
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

func (ll *doublyLinkedListBase[T]) unshift(v T) {
	new := &dllNode[T]{
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

func (ll *doublyLinkedListBase[T]) shift() (T, bool) {
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

func (ll *doublyLinkedListBase[T]) flush() {
	ll.tail = nil
	ll.head = nil
	ll.length = 0
}

func (ll *doublyLinkedListBase[T]) traverse(action func(v T)) {
	curr := ll.head
	for i := uint(0); i < ll.length; i++ {
		action(curr.val)
		curr = curr.next
	}
}
