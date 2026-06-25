package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLevels := maxDepth(root.Left)
	rightLevels := maxDepth(root.Right)

	return max(leftLevels, rightLevels) + 1
}
