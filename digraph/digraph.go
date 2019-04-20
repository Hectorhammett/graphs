package digraph

import (
	"github.com/hectorhammett/graphs/graph"
	"github.com/hectorhammett/graphs/node"
)

type digraph struct {
	nodesList     map[string]node.Node
	adjacencyList map[string][]node.Node
}

// NewDigraph instantiates a unidirectional graph
func NewDigraph() graph.Graph {
	adjacencyList := map[string][]node.Node{}
	nodesList := map[string]node.Node{}
	return &digraph{
		nodesList, adjacencyList,
	}
}

// GetNodesList from the graph
func (g *digraph) GetNodesList() map[string]node.Node {
	return g.nodesList
}

// GetAdjacencyList of the graph
func (g *digraph) GetAdjacencyList() map[string][]node.Node {
	return g.adjacencyList
}

// GetNodeAdjacencyList for a single node
func (g *digraph) GetNodeAdjacencyList(n node.Node) []node.Node {
	return g.adjacencyList[n.GetName()]
}

// AddVertice to the graph
func (g *digraph) AddVertice(a node.Node, b node.Node) {
	g.addNodeToList(a)
	g.addNodeToList(b)

	g.addToAdjacencylist(a, b)
}

// GetNode from the graph
func (g *digraph) GetNode(nodeName string) node.Node {
	return g.nodesList[nodeName]
}

func (g *digraph) addNodeToList(n node.Node) {
	_, exists := g.nodesList[n.GetName()]
	if exists {
		return
	}

	g.nodesList[n.GetName()] = n
}

func (g *digraph) addToAdjacencylist(from node.Node, to node.Node) {
	g.adjacencyList[from.GetName()] = append(g.adjacencyList[from.GetName()], to)
}
