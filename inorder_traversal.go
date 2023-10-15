package main

func InorderTraversal(root *TreeNode) []int {
	var result []int
	var internal func(root *TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		internal(root.Left)
		result = append(result, root.Val)
		internal(root.Right)
	}
	internal(root)
	return result
}
