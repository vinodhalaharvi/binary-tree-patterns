package main

func DeleteBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		root = nil
		return nil
	}

	root.Left = DeleteBinaryTree(root.Left)
	root.Right = DeleteBinaryTree(root.Right)
	root = nil
	return root
}
