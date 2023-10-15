package main

func SpiralOrderTraversal(root *TreeNode) []int {
	var result []int
	var currentLevel []*TreeNode
	var nextLevel []*TreeNode
	level := 0
	if root == nil {
		return result
	}
	currentLevel = append(currentLevel, root)
	for len(currentLevel) > 0 {
		node := currentLevel[len(currentLevel)-1]
		currentLevel = currentLevel[:len(currentLevel)-1]
		result = append(result, node.Val)
		if level%2 == 1 {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		} else {
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
		}
		if len(currentLevel) == 0 {
			currentLevel, nextLevel = nextLevel, currentLevel
			level++
		}
	}
	return result
}
