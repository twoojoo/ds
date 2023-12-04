package ds

//heap = priority queue
//binary tree where
//maxheap: children and grandchildren
//are always smaller than the current node

//implemented using a array
//see https://frontendmasters.com/courses/algorithms/heap/

type MinHeap[V any] struct {
	comparer func(a V, b V) int
	data     []V
}

func NewMinHeap[V any](comparer func(a V, b V) int) *MinHeap[V] {
	return &MinHeap[V]{
		comparer: comparer,
		data:     []V{},
	}
}

func getLeftChildIndex(parentIdx int) int {
	return (2 * parentIdx) + 1
}

func getRightChildIndex(parentIdx int) int {
	return (2 * parentIdx) + 2
}

func getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (mh *MinHeap[V]) heapifyUp(currIdx int) {
	if currIdx == 0 {
		return
	}

	parentIdx := getParentIndex(currIdx)
	parentVal := mh.data[parentIdx]
	currVal := mh.data[currIdx]

	comp := mh.comparer(currVal, parentVal)

	//if parent is greater
	if comp > 0 {
		//swap
		mh.data[currIdx] = parentVal
		mh.data[parentIdx] = currVal

		//and recurse up to root
		mh.heapifyUp(parentIdx)
	}
}

func (mh *MinHeap[V]) getValueAtIndex(idx int) (V, bool) {
	if idx >= len(mh.data) || idx < 0 {
		var zero V
		return zero, false
	}

	return mh.data[idx], true
}

func (mh *MinHeap[V]) heapifyDown(currIdx int) {
	if currIdx >= len(mh.data) {
		return
	}

	leftChildIdx := getLeftChildIndex(currIdx)
	rightChildIdx := getRightChildIndex(currIdx)

	//can't heapify down more
	if leftChildIdx >= len(mh.data) {
		return
	}

	//get right val only if right node exists
	rightVal, hasRightChild := mh.getValueAtIndex(rightChildIdx)
	currVal := mh.data[currIdx]
	leftVal := mh.data[leftChildIdx]

	//if right val is the smallest child, and curr are grater than him
	if hasRightChild && (mh.comparer(leftVal, rightVal) > 0 && mh.comparer(currVal, rightVal) > 0) {
		//swap with smallest child
		mh.data[currIdx] = rightVal
		mh.data[rightChildIdx] = currVal

		//and heapifydown
		mh.heapifyDown(rightChildIdx)
	} else if mh.comparer(rightVal, leftVal) > 0 && mh.comparer(currVal, leftVal) > 0 {
		//swap with smallest child
		mh.data[currIdx] = leftVal
		mh.data[leftChildIdx] = currVal

		//and heapifydown
		mh.heapifyDown(leftChildIdx)
	}
}

func (mh *MinHeap[V]) Push(val V) {
	mh.data = append(mh.data, val)
	mh.heapifyUp(len(mh.data) - 1)
}

func (mh *MinHeap[V]) Pop() (V, bool) {
	var zero V

	if len(mh.data) == 0 {
		return zero, false
	}

	val := mh.data[0]

	if len(mh.data) == 1 {
		mh.data = []V{}
		return val, true
	}

	//put last value in first position
	mh.data[0] = mh.data[len(mh.data)-1]

	//and remove last element
	mh.data = mh.data[:len(mh.data)-1]

	// and bubble it down
	mh.heapifyDown(0)

	return val, true
}

func (mh *MinHeap[V]) getArray() []V {
	return mh.data
}
