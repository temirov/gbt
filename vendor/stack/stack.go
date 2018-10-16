package stack

import (
	"fmt"
	"node"
)

type Stack []*node.Node

func (s Stack) Push(v *node.Node) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, *node.Node) {
	if len(s) == 0 {
		fmt.Errorf("Empty stack: %-v", s)
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

// Equal compares two stacks of nodes
// using node.Equal(other) for node comparison
func Equal(stack Stack, other Stack) bool {
	if len(stack) != len(other) {
		return false
	}
	for i, v := range stack {
		if !v.Equal(other[i]) {
			return false
		}
	}
	return true
}
