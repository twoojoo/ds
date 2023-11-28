package ds

type Node[K comparable, V any] struct {
	Data  V
	edges map[K]int
}

type Graph[K comparable, V any] struct {
	nodes map[K]*Node[K, V]
}

func NewGraph[K comparable, V any]() *Graph[K, V] {
	return &Graph[K, V]{
		nodes: map[K]*Node[K, V]{},
	}
}

func (g *Graph[K, V]) AddNode(ID K, data ...V) *Graph[K, V] {
	if _, ok := g.nodes[ID]; !ok {
		if len(data) > 0 {
			g.nodes[ID] = &Node[K, V]{
				Data:  data[0],
				edges: map[K]int{},
			}
		} else {
			g.nodes[ID] = &Node[K, V]{edges: map[K]int{}}
		}
	}

	return g
}

func (g *Graph[K, V]) RemoveNode(ID K) *Graph[K, V] {
	delete(g.nodes, ID)

	for k1 := range g.nodes {
		for k2 := range g.nodes[k1].edges {
			if k2 == ID {
				delete(g.nodes[k1].edges, ID)
			}
		}
	}

	return g
}

func (g *Graph[K, V]) AddEdge(nodeID1 K, nodeID2 K) *Graph[K, V] {
	if _, ok := g.nodes[nodeID1]; !ok {
		return g
	}

	if _, ok := g.nodes[nodeID2]; !ok {
		return g
	}

	g.nodes[nodeID1].edges[nodeID2] = 0

	return g
}

func (g *Graph[K, V]) RemoveEdge(nodeID1 K, nodeID2 K) *Graph[K, V] {
	if _, ok := g.nodes[nodeID1]; !ok {
		return g
	}

	if _, ok := g.nodes[nodeID2]; !ok {
		return g
	}

	delete(g.nodes[nodeID1].edges, nodeID2)

	return g
}

func (g *Graph[K, V]) DepthFirstSearch(startNodeID K, matcher func(n *Node[K, V]) bool) (K, []K, bool) {
	visited := make(map[K]bool, len(g.nodes))
	path := make([]K, 0, len(g.nodes))

	nodeID, ok := g.dfs(visited, &path, startNodeID, func(nodeID K, n *Node[K, V]) bool {
		return matcher(n)
	})

	return nodeID, path, ok
}

func (g *Graph[K, V]) dfs(visited map[K]bool, path *[]K, nodeID K, matcher func(nodeID K, n *Node[K, V]) bool) (K, bool) {
	if _, ok := visited[nodeID]; ok {
		return nodeID, false
	}

	visited[nodeID] = true

	curr := g.nodes[nodeID]

	if matcher(nodeID, curr) {
		return nodeID, true
	}

	*path = append((*path), nodeID)

	for edgeNodeID := range curr.edges {
		if ID, found := g.dfs(visited, path, edgeNodeID, matcher); found {
			return ID, true
		}
	}

	return nodeID, false
}

func (g *Graph[K, V]) FindConnectedComponents(startNodeID K) []map[K]struct{} {
	visited := make(map[K]bool, len(g.nodes))
	connComp := []map[K]struct{}{}
	currIdx := 0

	for len(visited) < len(g.nodes) {
		connComp = append(connComp, map[K]struct{}{})
		g.fcc(visited, connComp, currIdx, startNodeID)
		currIdx++

		//inefficient?
		for k := range g.nodes {
			if v, ok := visited[k]; !ok || !v {
				startNodeID = k
			}
		}
	}

	return connComp
}

func (g *Graph[K, V]) fcc(visited map[K]bool, connComp []map[K]struct{}, currIdx int, nodeID K) {
	g.dfs(visited, &[]K{}, nodeID, func(nodeID K, n *Node[K, V]) bool {
		connComp[currIdx][nodeID] = struct{}{}
		return false
	})
}

func (g *Graph[K, V]) BreadthFirstSearch(startNodeID K, matcher func(n *Node[K, V]) bool) (K, []K, bool) {
	q := NewQueue(startNodeID)
	visited := map[K]bool{}
	path := []K{}

	var currID K
	for !q.IsEmpty() {
		var ok bool
		currID, ok = q.Dequeue()
		if !ok {
			break
		}

		path = append(path, currID)
		visited[currID] = true

		curr := g.nodes[currID]

		if matcher(curr) {
			return currID, path, true
		}

		for k := range curr.edges {
			if !visited[k] {
				q.Enqueue(k)
			}
		}
	}

	var zero K
	return zero, path, false
}
