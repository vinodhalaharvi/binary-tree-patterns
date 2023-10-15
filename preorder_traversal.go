package main

func PreOrderTraversal(root *TreeNode) []int {
	var result []int
	var internal func(root *TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		internal(root.Left)
		internal(root.Right)
	}
	internal(root)
	return result
}
