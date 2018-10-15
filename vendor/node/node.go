package node

import (
	"fmt"
)

// Node is a type that describes a node in a binary tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// type Nodes []*Node

// String returns a value of a node as string
func (node *Node) String() string {
	return fmt.Sprintf("%d", node.Val)
}

// Children returns a slice containing the immediate children of a node
// The addition goes from left to right
func (node *Node) Children() (compact []*Node) {
	if node.Left != nil {
		compact = append(compact, node.Left)
	}

	if node.Right != nil {
		compact = append(compact, node.Right)
	}

	return compact
}

// HasChildren returns true when a node has children
// or false when a node has no children
func (node *Node) HasChildren() bool {
	if len(node.Children()) == 0 {
		return false
	}

	return true
}

// NoChildren returns true when a node has no children
// or false when a node has children
func (node *Node) NoChildren() bool {
	return !node.HasChildren()
}

// Equal compares a node to another node
// using node values for comparison
func (node *Node) Equal(other *Node) bool {
	return node.Val == other.Val
}

// Equal compares two slices of nodes
// using node.Equal(other) for node comparison
func Equal(array []*Node, other []*Node) bool {
	if len(array) != len(other) {
		return false
	}
	for i, v := range array {
		if !v.Equal(other[i]) {
			return false
		}
	}
	return true
}
