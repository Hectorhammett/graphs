package graph

import (
	"github.com/hectorhammett/graphs/node"
)

// Graph defines the unidirectional graph interface
type Graph interface {
	AddVertice(node.Node, node.Node)
	GetNodesList() map[string]node.Node
	GetAdjacencyList() map[string][]node.Node
	GetNodeAdjacencyList(node.Node) []node.Node
	GetNode(string) node.Node
}

type graph struct {
	nodesList     map[string]node.Node
	adjacencyList map[string][]node.Node
}

// NewGraph instantiates a unidirectional graph
func NewGraph() Graph {
	adjacencyList := map[string][]node.Node{}
	nodesList := map[string]node.Node{}
	return &graph{
		nodesList, adjacencyList,
	}
}

// GetNodesList from the graph
func (g *graph) GetNodesList() map[string]node.Node {
	return g.nodesList
}

// GetAdjacencyList of the graph
func (g *graph) GetAdjacencyList() map[string][]node.Node {
	return g.adjacencyList
}

// GetNodeAdjacencyList for a single node
func (g *graph) GetNodeAdjacencyList(n node.Node) []node.Node {
	return g.adjacencyList[n.GetName()]
}

// AddVertice to the graph
func (g *graph) AddVertice(a node.Node, b node.Node) {
	g.addNodeToList(a)
	g.addNodeToList(b)

	g.addToAdjacencylist(a, b)
}

// GetNode from the graph
func (g *graph) GetNode(nodeName string) node.Node {
	return g.nodesList[nodeName]
}

func (g *graph) addNodeToList(n node.Node) {
	_, exists := g.nodesList[n.GetName()]
	if exists {
		return
	}

	g.nodesList[n.GetName()] = n
}

func (g *graph) addToAdjacencylist(a node.Node, b node.Node) {
	g.adjacencyList[a.GetName()] = append(g.adjacencyList[a.GetName()], b)
	g.adjacencyList[b.GetName()] = append(g.adjacencyList[b.GetName()], a)
}
