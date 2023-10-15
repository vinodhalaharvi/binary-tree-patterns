package main

import (
	"reflect"
	"testing"
)

func TestSpiralOrderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
	}
	expected := []int{2, 1, 7, 5, 4, 6}
	DrawInputTree(root)
	actual := SpiralOrderTraversal(root)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v\n", expected, actual)
	}
}
