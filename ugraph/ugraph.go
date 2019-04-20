package ugraph

import (
	"github.com/hectorhammett/graphs/graph"
	"github.com/hectorhammett/graphs/node"
)

type ugraph struct {
	nodesList     map[string]node.Node
	adjacencyList map[string][]node.Node
}

// NewGraph instantiates a unidirectional graph
func NewUGraph() graph.Graph {
	adjacencyList := map[string][]node.Node{}
	nodesList := map[string]node.Node{}
	return &ugraph{
		nodesList, adjacencyList,
	}
}

// GetNodesList from the graph
func (g *ugraph) GetNodesList() map[string]node.Node {
	return g.nodesList
}

// GetAdjacencyList of the graph
func (g *ugraph) GetAdjacencyList() map[string][]node.Node {
	return g.adjacencyList
}

// GetNodeAdjacencyList for a single node
func (g *ugraph) GetNodeAdjacencyList(n node.Node) []node.Node {
	return g.adjacencyList[n.GetName()]
}

// AddVertice to the graph
func (g *ugraph) AddVertice(a node.Node, b node.Node) {
	g.addNodeToList(a)
	g.addNodeToList(b)

	g.addToAdjacencylist(a, b)
}

// GetNode from the graph
func (g *ugraph) GetNode(nodeName string) node.Node {
	return g.nodesList[nodeName]
}

func (g *ugraph) addNodeToList(n node.Node) {
	_, exists := g.nodesList[n.GetName()]
	if exists {
		return
	}

	g.nodesList[n.GetName()] = n
}

func (g *ugraph) addToAdjacencylist(a node.Node, b node.Node) {
	g.adjacencyList[a.GetName()] = append(g.adjacencyList[a.GetName()], b)
	g.adjacencyList[b.GetName()] = append(g.adjacencyList[b.GetName()], a)
}
