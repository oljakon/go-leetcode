package main

func preorder2(root *Node) []int {
	path := []int{}
	dfs(root, &path)
	return path
}

func dfs(node *Node, path *[]int) {
	if node == nil {
		return
	}

	*path = append(*path, node.Val)

	for i := range node.Children {
		dfs(node.Children[i], path)
	}
}
