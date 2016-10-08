package main

type graph struct {
	nodes []node
}

// function to see if the node is already in the graph
func (g graph) IsInGraph(n int) bool {
	for i := range g.nodes {
		if g.nodes[i].id == n {
			return true
		}
	}

	return false
}

func (g graph) InsertNode(n node) {
	if ( !g.IsInGraph(n.id) ) {
		g.nodes = append(g.nodes, n)
	}
}

// create a new graph with a specified number of nodes
func NewGraph(size int) graph {
	nodes := make([]node, size)
	return graph{nodes: nodes}
}