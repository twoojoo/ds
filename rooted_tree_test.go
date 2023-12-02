package ds

import "testing"

func TestRootedTreeAppendNode(t *testing.T) {
	rt := NewRootedTree[int, int](1, 10)

	if c := rt.CountNodes(); c != 1 {
		t.Fatal(c)
	}

	if err := rt.AppendNode(1, 2, 20); err != nil {
		t.Fatal(err)
	}

	if c := rt.CountNodes(); c != 2 {
		t.Fatal(c)
	}

	if err := rt.AppendNode(1, 3, 30); err != nil {
		t.Fatal(err)
	}

	if c := rt.CountNodes(); c != 3 {
		t.Fatal(c)
	}

	if c, err := rt.CountChildren(1); c != 2 || err != nil {
		t.Fatal(c, err)
	}

	if err := rt.AppendNode(3, 4, 40); err != nil {
		t.Fatal(err)
	}

	if err := rt.AppendNode(3, 5, 50); err != nil {
		t.Fatal(err)
	}

	if c := rt.CountNodes(); c != 5 {
		t.Fatal(c)
	}
}

func TestRootedTreeRemoveNodeChildren(t *testing.T) {
	rt := NewRootedTree[int, int](1, 10)
	rt.AppendNode(1, 2, 20)
	rt.AppendNode(1, 3, 30)
	rt.AppendNode(3, 4, 40)
	rt.AppendNode(3, 5, 50)

	if err := rt.RemoveNodeChildren(3); err != nil {
		t.Fatal(err)
	}

	if c := rt.CountNodes(); c != 3 {
		t.Fatal(c)
	}

	if c, err := rt.CountChildren(3); c != 0 || err != nil {
		t.Fatal(c, err)
	}

	if c, err := rt.CountChildren(1); c != 2 || err != nil {
		t.Fatal(c, err)
	}
}

func TestRootedTreeRemoveNode(t *testing.T) {
	rt := NewRootedTree[int, int](1, 10)
	rt.AppendNode(1, 2, 20)
	rt.AppendNode(1, 3, 30)
	rt.AppendNode(3, 4, 40)
	rt.AppendNode(3, 5, 50)

	if err := rt.RemoveNode(3); err != nil {
		t.Fatal(err)
	}

	if c := rt.CountNodes(); c != 2 {
		t.Fatal(c)
	}

	if c, err := rt.CountChildren(1); c != 1 || err != nil {
		t.Fatal(c, err)
	}
}
