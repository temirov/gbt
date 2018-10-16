// Given a binary tree (doesn't have to be a BST or any balanced), ie:

//       1
//      / \
//     10  6
//    /  \  \
//   6   3   7

// Find all the root->leaf paths and sum them, ie for the above tree:
// 1,10,6=>1106
// 1,10,3=>1103
// 1,6,7 =>167
// 1106+1103+167=2376

package main

import (
	"fmt"
	bn "node"
)

func dfsR(node *bn.Node, path *[]*bn.Node) {
	*path = append(*path, node)
	for _, n := range node.Children() {
		dfsR(n, path)
	}
}

func paths(node *bn.Node, path []*bn.Node, forest []*bn.Node) {
	path = append(path, node)
	if node.NoChildren() {
		forest = append(forest, path...)
		forest = append(forest, &bn.Node{Val: 0})
	}

	for _, child := range node.Children() {
		paths(child, path, forest)
		fmt.Printf("path before copy: %-v\n", path)
		// copy(path[:len(path)-1], path)
		// path[:len(path)-1] = nil
		fmt.Printf("path after copy: %-v\n", path)
	}
}

func dfs(root *bn.Node) []*bn.Node {
	stack := []*bn.Node{}
	accumulator := []*bn.Node{}
	stack = append(stack, root)

	for len(stack) > 0 {
		var n *bn.Node
		n, stack = stack[0], stack[1:]
		accumulator = append(accumulator, n)
		if n.HasChildren() {
			stack = append(stack, n.Children()...)
		}
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

	fullGraph := dfs(root)
	fmt.Printf("fullGraph: %-v\n", fullGraph)

	fullGraphR := &[]*bn.Node{}
	dfsR(root, fullGraphR)
	fmt.Printf("fullGraphR: %-v\n", fullGraphR)

	allPaths := []*bn.Node{}
	paths(root, []*bn.Node{}, allPaths)

	fmt.Printf("paths: %-v\n", allPaths)
}
