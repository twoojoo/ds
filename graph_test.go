package ds

import (
	// "fmt"
	// "math/rand"
	"slices"
	"testing"
)

func TestGraphAddNode(t *testing.T) {
	g := NewGraph[int, int]()
	if l := len(g.nodes); l != 0 {
		t.Fatal(l)
	}

	g.AddNode(1, 10)
	if l := len(g.nodes); l != 1 {
		t.Fatal(l)
	}

	g.AddNode(2, 10)
	g.AddNode(3, 10)
	if l := len(g.nodes); l != 3 {
		t.Fatal(l)
	}

	g.AddNode(3, 10)
	if l := len(g.nodes); l != 3 {
		t.Fatal(l)
	}
}

func TestGraphAddEdge(t *testing.T) {
	g := NewGraph[int, int]()
	g.AddNode(1, 10)
	g.AddNode(2, 10)
	g.AddNode(3, 10)
	g.AddNode(4, 10)
	g.AddNode(5, 10)

	if e := g.CountEdges(); e != 0 {
		t.Fatal(e)
	}

	g.AddEdge(2, 3)
	if e := g.CountEdges(); e != 1 {
		t.Fatal(e)
	}

	g.AddEdge(5, 3)
	g.AddEdge(4, 1)
	if e := g.CountEdges(); e != 3 {
		t.Fatal(e)
	}

	g.AddEdge(4, 1)
	if e := g.CountEdges(); e != 3 {
		t.Fatal(e)
	}
}

func TestGraphRemoveEdge(t *testing.T) {
	g := NewGraph[int, int]()
	g.AddNode(1, 10)
	g.AddNode(2, 10)
	g.AddNode(3, 10)
	g.AddNode(4, 10)
	g.AddNode(5, 10)
	g.AddEdge(2, 3)
	g.AddEdge(5, 3)
	g.AddEdge(4, 1)
	g.AddEdge(3, 2)

	n1 := g.CountNodes()

	g.RemoveEdge(3, 2)
	if e := g.CountEdges(); e != 3 {
		t.Fatal(e)
	}

	g.RemoveEdge(2, 3)
	if e := g.CountEdges(); e != 2 {
		t.Fatal(e)
	}

	g.RemoveEdge(5, 4)
	if e := g.CountEdges(); e != 2 {
		t.Fatal(e)
	}

	g.RemoveEdge(5, 3)
	g.RemoveEdge(4, 1)
	if e := g.CountEdges(); e != 0 {
		t.Fatal(e)
	}

	n2 := g.CountNodes()

	if n1 != n2 {
		t.Fatal(n1, n2)
	}
}

func TestGraphRemoveNode(t *testing.T) {
	g := NewGraph[int, int]()
	g.AddNode(1, 10)
	g.AddNode(2, 10)
	g.AddNode(3, 10)
	g.AddNode(4, 10)
	g.AddNode(5, 10)
	g.AddEdge(2, 3)
	g.AddEdge(5, 3)
	g.AddEdge(4, 1)
	g.AddEdge(3, 2)

	g.RemoveNode(2)
	if n := g.CountNodes(); n != 4 {
		t.Fatal(n)
	}
	if e := g.CountEdges(); e != 2 {
		t.Fatal(e)
	}

	g.RemoveNode(2)
	if n := g.CountNodes(); n != 4 {
		t.Fatal(n)
	}
	if e := g.CountEdges(); e != 2 {
		t.Fatal(e)
	}

	g.RemoveNode(3)
	g.RemoveNode(5)
	if n := g.CountNodes(); n != 2 {
		t.Fatal(n)
	}
	if e := g.CountEdges(); e != 1 {
		t.Fatal(e)
	}

	g.RemoveNode(1)
	g.RemoveNode(4)
	if n := g.CountNodes(); n != 0 {
		t.Fatal(n)
	}
	if e := g.CountEdges(); e != 0 {
		t.Fatal(e)
	}
}

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

func TestGraphFindConnectedComponentes(t *testing.T) {
	g := NewGraph[int, int]()
	g.AddNode(1, 10)
	g.AddNode(2, 20)
	g.AddNode(3, 30)
	g.AddNode(4, 40)
	g.AddNode(5, 50)
	g.AddNode(6, 60)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	g.AddEdge(4, 5)
	g.AddEdge(5, 4)

	g.AddEdge(5, 6)

	g.FindConnectedComponents(1)
}

// func TestMaze(t *testing.T) {
// 	g := NewGraph[int, string]()

// 	cols := 30
// 	rows := 10
// 	for i := 0; i < cols*rows; i++ {
// 		cell := " "
// 		rnd := rand.Intn(5)
// 		if rnd == 2 {
// 			cell = "#"
// 		}

// 		if i == cols*rows-1 {
// 			cell = "E"
// 		}

// 		if i == 0 {
// 			cell = "S"
// 		}

// 		g.AddNode(i, cell)
// 	}

// 	for i := 0; i < cols*rows; i++ {
// 		if g.nodes[i].Data == "#" {
// 			continue
// 		}

// 		if (i+1)%cols != 1 {
// 			if n, ok := g.nodes[i-1]; ok && n.Data != "#" {
// 				g.AddEdge(i, i-1)
// 			}
// 		}

// 		if (i+1)%cols != 0 {
// 			if n, ok := g.nodes[i+1]; ok && n.Data != "#" {
// 				g.AddEdge(i, i+1)
// 			}
// 		}

// 		if n, ok := g.nodes[i-cols]; ok && n.Data != "#" {
// 			g.AddEdge(i, i-cols)
// 		}

// 		if n, ok := g.nodes[i+cols]; ok && n.Data != "#" {
// 			g.AddEdge(i, i+cols)
// 		}
// 	}

// 	for i := 0; i < cols*rows; i++ {
// 		if i%cols == 0 {
// 			fmt.Print("\n")
// 		}

// 		fmt.Print(g.nodes[i].Data)
// 	}

// 	fmt.Print("\n")

// 	p, ok := g.ShortestPathBFS(0, cols*rows-1)
// 	fmt.Println(p)

// 	if ok {
// 		for i := 0; i < cols*rows; i++ {
// 			if i%cols == 0 {
// 				fmt.Print("\n")
// 			}

// 			if g.nodes[i].Data == "S" {
// 				fmt.Print(g.nodes[i].Data)
// 			} else if g.nodes[i].Data == "E" {
// 				fmt.Print(g.nodes[i].Data)
// 			} else if slices.Contains(p, i) {
// 				fmt.Print("+")
// 			} else {
// 				fmt.Print(g.nodes[i].Data)
// 			}
// 		}

// 		fmt.Print("\n")
// 	}

// }
