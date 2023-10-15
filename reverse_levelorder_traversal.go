package main

func ReverseLevelOrderTraversal(root *TreeNode) []int {
	var result []int
	var queue []*TreeNode
	var internal func(root *TreeNode)
	level := 0
	m := make(map[int][]*TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		queue = append(queue, root)
		for len(queue) > 0 {
			m[level] = append(m[level], queue[0])
			queue = queue[1:]
		}
	}
	internal(root)
	for i := len(m) - 1; i >= 0; i-- {
		for _, node := range m[i] {
			result = append(result, node.Val)
		}
	}
	return result
}
