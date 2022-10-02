package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	stack := []*Node{root}
	outputArr := []int{}

	if root == nil {
		return outputArr
	}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr != nil {
			outputArr = append(outputArr, curr.Val)
		}

		if len(curr.Children) == 0 {
			continue
		}

		for i := len(curr.Children) - 1; i >= 0; i-- {
			stack = append(stack, curr.Children[i])
		}
	}

	return outputArr
}
