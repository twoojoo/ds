package ds

import (
	"slices"
	"testing"
)

func TestGraphBFS(t *testing.T) {
	g := NewGraph[int, int]()

	g.AddNode(1, 10)
	g.AddNode(2, 20)
	g.AddNode(3, 30)
	g.AddNode(4, 40)
	g.AddNode(5, 50)

	g.AddEdge(1, 3)
	g.AddEdge(2, 1)
	g.AddEdge(3, 5)
	g.AddEdge(3, 1)
	g.AddEdge(4, 3)
	g.AddEdge(2, 4)
	g.AddEdge(5, 2)
	g.AddEdge(5, 1)

	matcher := func(n *Node[int, int]) bool {
		return n.Data == 40
	}

	if i, p, ok := g.BreadthFirstSearch(1, matcher); !ok || i != 4 || !slices.Equal(p, []int{1, 3, 5, 2, 4}) {
		t.Fatal(ok, i, p)
	}
}

func TestGraphDFS(t *testing.T) {
	g := NewGraph[int, int]()

	g.AddNode(1, 10)
	g.AddNode(2, 20)
	g.AddNode(3, 30)
	g.AddNode(4, 40)
	g.AddNode(5, 50)

	g.AddEdge(1, 3)
	g.AddEdge(2, 1)
	g.AddEdge(3, 5)
	g.AddEdge(4, 3)
	g.AddEdge(2, 4)
	g.AddEdge(5, 2)

	matcher := func(n *Node[int, int]) bool {
		return n.Data == 40
	}

	if i, p, ok := g.DepthFirstSearch(1, matcher); !ok || i != 4 || slices.Equal(p, []int{1, 3, 5, 2, 4}) {
		t.Fatal(ok, i, p)
	}
}
