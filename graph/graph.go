package graph

import "github.com/hectorhammett/graphs/node"

// Graph defines a generic graph interface
type Graph interface {
	AddVertice(node.Node, node.Node)
	GetNodesList() map[string]node.Node
	GetAdjacencyList() map[string][]node.Node
	GetNodeAdjacencyList(node.Node) []node.Node
	GetNode(string) node.Node
}
