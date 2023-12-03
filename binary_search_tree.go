package ds

type BinarySearchTree[V any] struct {
	comparer func(a V, b V) int
	root     *BinaryTreeNode[V]
}

func NewBinarySearchTree[V any](val V, comparer func(a V, b V) int) *BinarySearchTree[V] {
	return &BinarySearchTree[V]{
		root:     NewBinaryTreeNode(val, nil, nil),
		comparer: comparer,
	}
}

func NewBinarySearchTreeFromBinaryTree[V any](root *BinaryTreeNode[V], comparer func(a V, b V) int) *BinarySearchTree[V] {
	return &BinarySearchTree[V]{
		root:     root,
		comparer: comparer,
	}
}

func (bst *BinarySearchTree[V]) Find(val V) (V, bool) {
	return find(bst.root, val, bst.comparer)
}

func find[V any](root *BinaryTreeNode[V], val V, comparer func(a V, b V) int) (V, bool) {
	var zero V

	if root == nil {
		return zero, false
	}

	comp := comparer(val, root.Value())
	if comp == 0 {
		return val, true
	}

	if comp > 0 {
		if right, ok := root.Right(); ok {
			return find(right, val, comparer)
		} else {
			return zero, false
		}
	}

	if left, ok := root.Left(); ok {
		return find(left, val, comparer)
	} else {
		return zero, false
	}
}
