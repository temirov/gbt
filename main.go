package main

import (
	"fmt"
	bn "node"
)

func paths(node *bn.Node, forest []*bn.Node) []*bn.Node {
	if node.HasChildren() {
		for _, child := range node.Children() {
			forest = append(forest, node, child)
			paths(child, forest)
		}
	} else {
		forest = append(forest, node)
	}
	return forest

}

func dfs(root *bn.Node) []*bn.Node {
	stack := []*bn.Node{}
	stack = append(stack, root)
	fmt.Printf("len(stack): %-v\n", len(stack))
	var n *bn.Node
	accumulator := []*bn.Node{}
	for len(stack) > 0 {
		fmt.Printf("stack: %-v\n", stack)
		n, stack = stack[0], stack[1:]
		accumulator = append(accumulator, n)
		fmt.Printf("n: %-v\n", n)
		fmt.Printf("stack: %-v\n", stack)

		if n.Right != nil {
			fmt.Printf("added: %-v\n", n.Right)
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			fmt.Printf("added: %-v\n", n.Left)
			stack = append(stack, n.Left)
		}
		if n.Left == n.Right && n.Right == nil {
			accumulator = append(accumulator, &bn.Node{Val: 100})
		}
		fmt.Printf("stack: %-v\n", stack)
		fmt.Printf("len(stack): %-v\n", len(stack))

	}
	return accumulator
}

func main() {
	root := &bn.Node{Val: 1}
	root.Left = &bn.Node{Val: 10}
	root.Right = &bn.Node{Val: 6}
	root.Left.Left = &bn.Node{Val: 6}
	root.Left.Right = &bn.Node{Val: 3}
	root.Right.Left = &bn.Node{Val: 7}

	// fmt.Printf("node: %v", root)

	// trees := paths(root, []*Node{})
	// fmt.Printf("trees: %-v", trees)

	fullGraph := dfs(root)
	fmt.Printf("fullGraph: %-v\n", fullGraph)
}
