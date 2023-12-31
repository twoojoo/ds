package ds

import "testing"

func TestBinaryTreeFind(t *testing.T) {
	bst := NewBinarySearchTreeFromBinaryTree(
		NewBinaryTreeNode[int](10,
			NewBinaryTreeNode(12,
				NewBinaryTreeNode(25, nil, nil),
				nil,
			),
			NewBinaryTreeNode(7,
				NewBinaryTreeNode(9, nil, nil),
				NewBinaryTreeNode(2, nil, nil),
			),
		), func(a, b int) int {
			if a > b {
				return 1
			}

			if b > a {
				return -1
			}

			return 0
		},
	)

	val := 10
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	val = 12
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	val = 25
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	val = 7
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	val = 9
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	val = 2
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}
}

func TestBSTInsert(t *testing.T) {
	bst := NewBinarySearchTreeFromBinaryTree(
		NewBinaryTreeNode[int](10,
			NewBinaryTreeNode(12,
				NewBinaryTreeNode(25, nil, nil),
				nil,
			),
			NewBinaryTreeNode(7,
				NewBinaryTreeNode(9, nil, nil),
				NewBinaryTreeNode(2, nil, nil),
			),
		), func(a, b int) int {
			if a > b {
				return 1
			}

			if b > a {
				return -1
			}

			return 0
		},
	)

	bst.Insert(33)
	val := 33
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	bst.Insert(1)
	val = 1
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}

	bst.Insert(19)
	val = 19
	if v, ok := bst.Find(val); !ok || v != val {
		t.Fatal(ok, v)
	}
}

func TestBSTRemove(t *testing.T) {
	bst := NewBinarySearchTreeFromBinaryTree(
		NewBinaryTreeNode[int](10,
			NewBinaryTreeNode(12,
				NewBinaryTreeNode(25, nil, nil),
				nil,
			),
			NewBinaryTreeNode(7,
				NewBinaryTreeNode(9, nil, nil),
				NewBinaryTreeNode(2, nil, nil),
			),
		), func(a, b int) int {
			if a > b {
				return 1
			}

			if b > a {
				return -1
			}

			return 0
		},
	)

	bst.Insert(11)
	bst.Insert(33)
	bst.Insert(1)
	bst.Insert(19)

	bst.Remove(12)

	if _, ok := bst.Find(12); ok {
		t.Fatal(ok)
	}

	if _, ok := bst.Find(25); !ok {
		t.Fatal(ok)
	}

	bst.Remove(25)

	if _, ok := bst.Find(33); !ok {
		t.Fatal(ok)
	}

	if _, ok := bst.Find(11); !ok {
		t.Fatal(ok)
	}
}
