package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var isValid func(node *TreeNode, left, right int) bool

	isValid = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}

		if !(node.Val > left && node.Val < right) {
			return false
		}

		return isValid(node.Left, left, node.Val) && isValid(node.Right, node.Val, right)
	}

	return isValid(root, math.MinInt, math.MaxInt)
}
