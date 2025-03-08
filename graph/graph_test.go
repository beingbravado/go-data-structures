package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
	assert.Empty(t, g.nodes)
}

func TestAddNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1)
	assert.Contains(t, g.nodes, 1)
	assert.Equal(t, 1, g.nodes[1].num)

	// Adding duplicate node
	g.AddNode(1)
	assert.Len(t, g.nodes, 1) // Ensure no duplicate nodes
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)

	assert.Contains(t, g.nodes, 1)
	assert.Contains(t, g.nodes, 2)
	assert.Equal(t, 2, g.nodes[1].next[0].num)

	// Adding duplicate edge
	g.AddEdge(1, 2)
	assert.Len(t, g.nodes[1].next, 2) // Ensure multiple edges can exist

	// Adding self-loop
	g.AddEdge(3, 3)
	assert.Contains(t, g.nodes, 3)
	assert.Equal(t, 3, g.nodes[3].next[0].num)
}

func TestCreateGraph(t *testing.T) {
	prereqs := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	g := CreateGraph(4, prereqs)

	assert.Contains(t, g.nodes, 0)
	assert.Contains(t, g.nodes, 1)
	assert.Contains(t, g.nodes, 2)
	assert.Contains(t, g.nodes, 3)

	assert.Len(t, g.nodes[0].next, 2) // 0 -> [1, 2]
	assert.Len(t, g.nodes[1].next, 1) // 1 -> [3]
	assert.Len(t, g.nodes[2].next, 1) // 2 -> [3]
	assert.Len(t, g.nodes[3].next, 0) // 3 -> []

	// Test empty prerequisites
	emptyGraph := CreateGraph(5, [][]int{})
	assert.Len(t, emptyGraph.nodes, 0) // No edges means no nodes should be created

	// Test single node with no edges
	singleNodeGraph := CreateGraph(1, [][]int{})
	assert.Len(t, singleNodeGraph.nodes, 0) // Single node without edges shouldn't exist
}
