package main

func LevelOrderTraversal(root *TreeNode) []int {
	var result []int
	var queue []*TreeNode
	if root == nil {
		return result
	}
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil || node.Right != nil {
			level++
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}
