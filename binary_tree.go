package ds

import (
	"math"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

type BinaryTreeNode[V any] struct {
	val   V
	left  *BinaryTreeNode[V]
	right *BinaryTreeNode[V]
}

func NewBinaryTreeNode[V any](val V, right *BinaryTreeNode[V], left *BinaryTreeNode[V]) *BinaryTreeNode[V] {
	return &BinaryTreeNode[V]{
		val:   val,
		left:  left,
		right: right,
	}
}

func (btn *BinaryTreeNode[V]) Value() V {
	return btn.val
}

func (btn *BinaryTreeNode[V]) Right() (*BinaryTreeNode[V], bool) {
	if btn.right != nil {
		return btn.right, true
	}

	return nil, false
}

func (btn *BinaryTreeNode[V]) IsLeaf() bool {
	return btn.right == nil && btn.left == nil
}

func (btn *BinaryTreeNode[V]) Left() (*BinaryTreeNode[V], bool) {
	if btn.left != nil {
		return btn.left, true
	}

	return nil, false
}

func BinaryTreeDepthFirstSearch[V any](root *BinaryTreeNode[V], matcher func(node *BinaryTreeNode[V]) bool) (V, bool) {
	if root == nil {
		var zero V
		return zero, false
	}

	stack := NewStack(root)

	for !stack.IsEmpty() {
		curr, _ := stack.Remove()

		if matcher(curr) {
			return curr.val, true
		}

		if r, ok := curr.Right(); ok {
			stack.Add(r)
		}

		if l, ok := curr.Left(); ok {
			stack.Add(l)
		}
	}

	var zero V
	return zero, false
}

func BinaryTreeBreadthFirstSearch[V any](root *BinaryTreeNode[V], matcher func(node *BinaryTreeNode[V]) bool) (V, bool) {
	if root == nil {
		var zero V
		return zero, false
	}

	queue := NewQueue(root)

	for !queue.IsEmpty() {
		curr, _ := queue.Dequeue()

		if matcher(curr) {
			return curr.val, true
		}

		if r, ok := curr.Right(); ok {
			queue.Enqueue(r)
		}

		if l, ok := curr.Left(); ok {
			queue.Enqueue(l)
		}
	}

	var zero V
	return zero, false
}

func BinaryTreeSum[V number](root *BinaryTreeNode[V]) V {
	var sum V

	BinaryTreeBreadthFirstSearch(root, func(node *BinaryTreeNode[V]) bool {
		sum += node.Value()
		return false
	})

	return sum
}

func BinaryTreeMaxPathToLeaf[V number](root *BinaryTreeNode[V]) V {
	if root == nil {
		return V(math.Inf(-1))
	}

	if root.IsLeaf() {
		return root.val
	}

	s1 := V(math.Inf(-1))
	if r, ok := root.Right(); ok {
		s1 = BinaryTreeMaxPathToLeaf(r)
	}

	s2 := V(math.Inf(-1))
	if l, ok := root.Left(); ok {
		s2 = BinaryTreeMaxPathToLeaf(l)
	}

	maxSum := V(math.Max(float64(s1), float64(s2)))

	return maxSum + root.val
}

func CompareBinaryTrees[V comparable](roots ...*BinaryTreeNode[V]) bool {
	validRoots := []*BinaryTreeNode[V]{}
	rightChildren := []*BinaryTreeNode[V]{}
	leftChildren := []*BinaryTreeNode[V]{}
	values := map[V]struct{}{}

	for i := range roots {
		if roots[i] != nil {
			validRoots = append(validRoots, roots[i])

			if r, ok := roots[i].Right(); ok {
				rightChildren = append(rightChildren, r)
			}

			if l, ok := roots[i].Left(); ok {
				leftChildren = append(leftChildren, l)
			}

			values[roots[i].val] = struct{}{}
		}
	}

	if len(validRoots) == 0 {
		return true
	}

	if len(validRoots) != len(roots) {
		return false
	}

	if len(values) != 1 {
		return false
	}

	rightEq := CompareBinaryTrees(rightChildren...)
	leftEq := CompareBinaryTrees(leftChildren...)

	return rightEq && leftEq
}
