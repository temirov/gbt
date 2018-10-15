package node

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("%d", node.Val)
}

func (node *Node) Children() (compact []*Node) {
	if node.Left != nil {
		compact = append(compact, node.Left)
	}

	if node.Right != nil {
		compact = append(compact, node.Right)
	}

	return compact
}

func (node *Node) HasChildren() bool {
	if len(node.Children()) == 0 {
		return false
	}

	return true
}

func (node *Node) NoChildren() bool {
	return !node.HasChildren()
}
