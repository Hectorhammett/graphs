package main

import (
	"fmt"

	"github.com/hectorhammett/graphs/graph"
	"github.com/hectorhammett/graphs/node"
)

func main() {
	a := node.NewNode("a", nil)
	b := node.NewNode("b", nil)
	c := node.NewNode("c", nil)
	d := node.NewNode("d", nil)

	g := graph.NewGraph()
	g.AddVertice(a, c)
	g.AddVertice(b, c)
	g.AddVertice(d, c)

	for _, n := range g.GetNodesList() {
		for _, c := range g.GetNodeAdjacencyList(n) {
			fmt.Printf("%s is connected to %s\n", n.GetName(), c.GetName())
		}
	}
}
