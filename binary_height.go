package main

func HeightOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(HeightOfBinaryTree(root.Left), HeightOfBinaryTree(root.Right))
}
