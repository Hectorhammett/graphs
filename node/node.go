package node

// Node defines the Node interface
type Node interface {
	GetName() string
	GetContent() interface{}
}

type node struct {
	name    string
	content interface{}
}

// NewNode instantiates a Node interface
func NewNode(name string, content interface{}) Node {
	return &node{
		name,
		content,
	}
}

// GetContent of the node
func (n *node) GetContent() interface{} {
	return n.content
}

// GetName of the node
func (n *node) GetName() string {
	return n.name
}
