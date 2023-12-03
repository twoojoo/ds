package ds

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := NewBinaryTreeNode[string]("a",
		NewBinaryTreeNode("b",
			NewBinaryTreeNode("d", nil, nil),
			nil,
		),
		NewBinaryTreeNode("c",
			NewBinaryTreeNode("e", nil, nil),
			nil,
		),
	)

	if r, ok := root.Right(); !ok || r.Value() != "b" {
		t.Fatal(ok, r)

		if r, ok = r.Right(); !ok || r.Value() != "d" {
			t.Fatal(ok, r)
		}
	}

	if l, ok := root.Left(); !ok || l.Value() != "c" {
		t.Fatal(ok, l)

		if r, ok := l.Right(); !ok || r.Value() != "e" {
			t.Fatal(ok, r)
		}
	}
}

func TestBinaryTreeDFS(t *testing.T) {
	root := NewBinaryTreeNode[string]("a",
		NewBinaryTreeNode("b",
			NewBinaryTreeNode("d", nil, nil),
			nil,
		),
		NewBinaryTreeNode("c",
			NewBinaryTreeNode("e", nil, nil),
			NewBinaryTreeNode("f", nil, nil),
		),
	)

	if v, ok := BinaryTreeBreadthFirstSearch(root, func(v *BinaryTreeNode[string]) bool { return v.Value() == "f" }); !ok || v != "f" {
		t.Fatal(ok, v)
	}

	if v, ok := BinaryTreeBreadthFirstSearch(root, func(v *BinaryTreeNode[string]) bool { return v.Value() == "g" }); ok {
		t.Fatal(ok, v)
	}
}

func TestBinaryTreeBFS(t *testing.T) {
	root := NewBinaryTreeNode[string]("a",
		NewBinaryTreeNode("b",
			NewBinaryTreeNode("d", nil, nil),
			nil,
		),
		NewBinaryTreeNode("c",
			NewBinaryTreeNode("e", nil, nil),
			NewBinaryTreeNode("f", nil, nil),
		),
	)

	if v, ok := BinaryTreeBreadthFirstSearch(root, func(v *BinaryTreeNode[string]) bool { return v.Value() == "f" }); !ok || v != "f" {
		t.Fatal(ok, v)
	}

	if v, ok := BinaryTreeBreadthFirstSearch(root, func(v *BinaryTreeNode[string]) bool { return v.Value() == "g" }); ok {
		t.Fatal(ok, v)
	}
}

func TestBinaryTreeSum(t *testing.T) {
	root := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(1, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	if v := BinaryTreeSum(root); v != 36 {
		t.Fatal(v)
	}
}

func TestBinaryTreeMaxPathToLeaf(t *testing.T) {
	root := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(1, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	if v := BinaryTreeMaxPathToLeaf(root); v != 21 {
		t.Fatal(v)
	}
}

func TestBinaryTreeCompare(t *testing.T) {
	root1 := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(1, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	root2 := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(1, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	root3 := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(1, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	root4 := NewBinaryTreeNode[int](5,
		NewBinaryTreeNode(12,
			NewBinaryTreeNode(2, nil, nil),
			nil,
		),
		NewBinaryTreeNode(7,
			NewBinaryTreeNode(2, nil, nil),
			NewBinaryTreeNode(9, nil, nil),
		),
	)

	if v := CompareBinaryTrees(root1, root2, root3); !v {
		t.Fatal(v)
	}

	if v := CompareBinaryTrees(root1, root2, root3, root4); v {
		t.Fatal(v)
	}
}
