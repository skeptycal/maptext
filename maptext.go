package maptext

import "fmt"

type (
	Any = interface{}

	textMap struct {
		m map[int]string
	}

	node struct {
		name string
		data []byte

		// Connecting nodes
		parent *node
		left   *node
		right  *node
	}
)

func NewRoot(name string, data []byte) *node {
	return &node{name, data, nil, nil, nil}
}

func (n *node) String() string {
	return fmt.Sprintf("%s: %s", n.name, n.data)
}

func (n *node) Len() int           { return len(n.data) }
func (n *node) Swap(i, j int)      { n.data, n.name = n.parent.data, n.parent.name }
func (n *node) Less(i, j int) bool { return len(n.data) < len(n.parent.data) }

func (n *node) Delete()
