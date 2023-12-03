package ds

import (
	"math"
)

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

// O(h)
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

// O(h)
func (bst *BinarySearchTree[V]) Insert(val V) {
	insert(bst.root, val, bst.comparer)
}

func insert[V any](root *BinaryTreeNode[V], val V, comparer func(a V, b V) int) {
	comp := comparer(val, root.Value())

	if comp > 0 {
		if right, ok := root.Right(); ok {
			insert(right, val, comparer)
		} else {
			root.right = &BinaryTreeNode[V]{
				val: val,
			}
		}

		return
	}

	if left, ok := root.Left(); ok {
		insert(left, val, comparer)
	} else {
		root.left = &BinaryTreeNode[V]{
			val: val,
		}
	}
}

func (bst *BinarySearchTree[V]) Remove(val V) {
	var prevNode *BinaryTreeNode[V] = nil
	var prevFromRight bool

	currNode := bst.root
	for {
		comp := bst.comparer(val, currNode.Value())

		if comp == 0 {
			remove(currNode, prevNode, prevFromRight)
			return
		}

		if comp > 0 {
			if right, ok := currNode.Right(); ok {
				prevFromRight = true
				prevNode = currNode
				currNode = right
			}
		} else {
			if left, ok := currNode.Left(); ok {
				prevFromRight = false
				prevNode = currNode
				currNode = left
			}
		}
	}
}

func remove[V any](node *BinaryTreeNode[V], prevNode *BinaryTreeNode[V], prevFromRight bool) {
	if node.IsLeaf() {
		if prevFromRight {
			prevNode.right = nil
		} else {
			prevNode.left = nil
		}
	}

	stepsToSmallestOnRight := math.Inf(1)
	stepsToGreatestOnLeft := math.Inf(1)
	var smallestOnRight *BinaryTreeNode[V] = nil
	var greatestOnLeft *BinaryTreeNode[V] = nil
	var rightNode *BinaryTreeNode[V] = nil
	var leftNode *BinaryTreeNode[V] = nil

	var ok bool
	if rightNode, ok = node.Right(); ok {
		smallestOnRight, stepsToSmallestOnRight = FindSmallest(rightNode, 0)
	}

	if leftNode, ok = node.Left(); ok {
		greatestOnLeft, stepsToGreatestOnLeft = FindGreater(leftNode, 0)
	}

	// choose the best strategy to keep the tree well balanced
	if stepsToGreatestOnLeft < stepsToSmallestOnRight {
		if prevFromRight {
			prevNode.right = greatestOnLeft
			greatestOnLeft.right = rightNode
		} else {
			prevNode.left = greatestOnLeft
			greatestOnLeft.right = rightNode
		}
	} else {
		if prevFromRight {
			prevNode.right = smallestOnRight
			smallestOnRight.left = leftNode
		} else {
			prevNode.left = smallestOnRight
			smallestOnRight.left = leftNode
		}
	}
}

func FindGreater[V any](root *BinaryTreeNode[V], steps float64) (*BinaryTreeNode[V], float64) {
	if right, ok := root.Right(); ok {
		steps++
		return FindGreater(right, steps)
	}

	return root, steps
}

func FindSmallest[V any](root *BinaryTreeNode[V], steps float64) (*BinaryTreeNode[V], float64) {
	if left, ok := root.Left(); ok {
		steps++
		return FindSmallest(left, steps)
	}

	return root, steps
}
