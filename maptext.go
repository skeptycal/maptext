package maptext

import (
	"crypto"
	"crypto/md5"
	"fmt"

	"github.com/skeptycal/errorlogger"
)

type direction = int

const (
	none direction = iota
	left
	right
	both
)

const (
	b1 = 0b0001
	b2 = 0b0010
	b4 = 0b0100
	b8 = 0b1000
)

var log = errorlogger.Log

func init() {
	if ok := crypto.MD5.Available(); !ok {
		log.Fatal("crypto.MD5 hashing is not available: %v", ok)
	}

	// h := md5.New()s

}

type (
	Any = interface{}

	id struct {
		n int
		s string
	}

	textMap struct {
		m map[int]string
	}

	// nodeMap represents a Binary Tree structure
	//
	/// TODO: rules:
	// - last == parent means size is 1 (root only)
	// - could last == nil work better?
	//
	// - name is stored in nodeMap and parent ...
	nodeMap struct {
		n     int          // n is the number of nodes
		name  string       // name is the name of the node set
		root  *node        // root is the root node of the tree
		last  *node        // last is the last node in the tree
		nodes map[id]*node // map of nodes by index id
		edges map[id]*edge // map of edges by index id
	}

	edge struct {
		n1  *node
		n2  *node
		dir direction
	}

	node struct {
		name     string
		data     []byte
		checksum [16]byte

		// Connecting nodes
		root   *node
		parent *node
		left   *node
		right  *node
	}
)

func (i *id) String() string {
	if i.s == "" {
		i.s = fmt.Sprintf("%09d", i.n)
	}
	return i.s
}

func (e *edge) CanLeft() bool { return e.dir&b2 != 0 }

func (e *edge) CanRight() bool { return e.dir&b1 != 0 }

func NewNodeMap(name string, parent *node) *nodeMap {
	return &nodeMap{
		name: name,
		root: parent,
		last: parent, // TODO: or nil initially??
	}
}

func NewNode(name string, data []byte, root, parent *node) *node {
	n := node{
		name:     name,
		data:     data,
		checksum: CheckSum(data),
		root:     root,
		parent:   parent,
		left:     nil,
		right:    nil,
	}
	return &n
}

func NewRoot(name string, data []byte) *node {
	n := node{
		name:     name,
		data:     data,
		checksum: CheckSum(data),
		root:     nil,
		parent:   nil,
		left:     nil,
		right:    nil,
	}
	n.root = &n
	return &n
}

func (n *node) String() string {
	return fmt.Sprintf("%s: %s", n.name, n.data)
}

func (n *node) Len() int {
	if n.parent == nil {
		return len(n.data)
	}
	return n.root.Len()
}
func (n *node) Swap(i, j int)          { n.data, n.name = n.parent.data, n.parent.name }
func (n *node) Less(i, j int) bool     { return len(n.data) < len(n.parent.data) }
func (n *node) Equal(other *node) bool { return string(n.data) == string(other.data) }
func (n *node) CheckSum() [16]byte     { return CheckSum(n.data) }
func (n *node) AddLeft(other *node) {
	n.left = other
	other.parent = n
}
func (n *node) AddRight(other *node) {
	n.right = other
	other.parent = n
}

func (n *node) Delete() {}

func CheckSum(data []byte) [16]byte { return md5.Sum(data) }
