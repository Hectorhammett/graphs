package main

import (
	"fmt"

	"github.com/hectorhammett/graphs/digraph"
	"github.com/hectorhammett/graphs/graph"
	"github.com/hectorhammett/graphs/node"
	"github.com/hectorhammett/graphs/ugraph"
)

func main() {
	a := node.NewNode("a", nil)
	b := node.NewNode("b", nil)
	c := node.NewNode("c", nil)
	d := node.NewNode("d", nil)

	g := ugraph.NewUGraph()
	g.AddVertice(a, c)
	g.AddVertice(b, c)
	g.AddVertice(d, c)

	ug := digraph.NewDigraph()
	ug.AddVertice(a, c)
	ug.AddVertice(b, c)
	ug.AddVertice(d, c)

	fmt.Println("UGraph:")
	traverseGraph(g)

	fmt.Println("Digraph:")
	traverseGraph(ug)
}

func traverseGraph(g graph.Graph) {
	for _, n := range g.GetNodesList() {
		for _, c := range g.GetNodeAdjacencyList(n) {
			fmt.Printf("%s is connected to %s\n", n.GetName(), c.GetName())
		}
	}
}
