package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	res := [][]int{}

	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) != 0 {
		level := []int{}
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr != nil {
				level = append(level, curr.Val)
				queue = append(queue, curr.Left)
				queue = append(queue, curr.Right)
			}
		}

		if len(level) != 0 {
			res = append(res, level)
		}
	}

	return res
}
