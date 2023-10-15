package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsIdentical Check if two given binary trees are identical or not | Iterative & Recursive
func IsIdentical(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return IsIdentical(root1.Left, root2.Left) && IsIdentical(root1.Right, root2.Right)
}
