package ds

import "fmt"

// TREE = undirected graph with no cycles

type rootedTreeNode[K comparable, V any] struct {
	value    V
	children map[K]struct{}
}

type RootedTree[K comparable, V any] struct {
	nodes map[K]*rootedTreeNode[K, V]
}

func NewRootedTree[K comparable, V any](rootKey K, rootVal V) *RootedTree[K, V] {
	return &RootedTree[K, V]{
		nodes: map[K]*rootedTreeNode[K, V]{
			rootKey: {
				value:    rootVal,
				children: map[K]struct{}{},
			},
		},
	}
}

func (t *RootedTree[K, V]) CountNodes() uint {
	return uint(len(t.nodes))
}

func (t *RootedTree[K, V]) CountChildren(ID K) (uint, error) {
	if node, ok := t.nodes[ID]; ok {
		return uint(len(node.children)), nil
	}

	return 0, fmt.Errorf("node %v doesn't exist", ID)
}

func (t *RootedTree[K, V]) AppendNode(parentID K, ID K, val V) error {
	if p, ok := t.nodes[parentID]; ok {
		t.nodes[ID] = &rootedTreeNode[K, V]{
			value:    val,
			children: map[K]struct{}{},
		}

		p.children[ID] = struct{}{}
		return nil
	}

	return fmt.Errorf("parent node doesn't exist")
}

func (t *RootedTree[K, V]) RemoveNode(ID K) error {
	if _, ok := t.nodes[ID]; ok {
		t.RemoveNodeChildren(ID)
		delete(t.nodes, ID)

		for k := range t.nodes {
			delete(t.nodes[k].children, ID)
		}

		return nil
	}

	return fmt.Errorf("node %v doesn't exist", ID)
}

func (t *RootedTree[K, V]) RemoveNodeChildren(ID K) error {
	if node, ok := t.nodes[ID]; ok {
		for childID := range node.children {
			t.RemoveNodeChildren(childID)
			delete(t.nodes, childID)
		}

		node.children = map[K]struct{}{}
		return nil
	}

	return fmt.Errorf("node %v doesn't exist", ID)
}
