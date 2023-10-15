package main

func PostOrderTraversal(root *TreeNode) []int {
	var result []int
	var internal func(root *TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		internal(root.Left)
		internal(root.Right)
		result = append(result, root.Val)
	}
	internal(root)
	return result
}
