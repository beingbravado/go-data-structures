package graph

// Graph structure represents a directed graph using an adjacency list.
// The graph is represented as a map where the key is the node number and the value is a pointer to the Node.
type Graph struct {
	nodes map[int]*Node
}

// Node structure represents each node in the graph.
// Each node has a number (num) and a slice of pointers to its adjacent nodes (next).
type Node struct {
	num  int
	next []*Node
}

// AddNode function adds a new node to the graph with the given number.
// If the node already exists, it does nothing.
func (g *Graph) AddNode(num int) {
	if _, ok := g.nodes[num]; !ok {
		g.nodes[num] = &Node{num: num}
	}
}

// AddEdge function adds a directed edge from the node with number 'from' to the node with number 'to'.
// It ensures both nodes exist in the graph before adding the edge.
func (g *Graph) AddEdge(from, to int) {
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from].next = append(g.nodes[from].next, g.nodes[to])
}

// NewGraph function initializes and returns a new empty graph.
func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]*Node)}
}

// CreateGraph function creates a graph from the given number of nodes (n) and a list of prerequisite pairs (prereqs).
// Each pair in prereqs represents a directed edge where the first element is the prerequisite for the second element.
func CreateGraph(n int, prereqs [][]int) *Graph {
	graph := NewGraph()

	for _, edge := range prereqs {
		graph.AddEdge(edge[1], edge[0])
	}

	return graph
}
