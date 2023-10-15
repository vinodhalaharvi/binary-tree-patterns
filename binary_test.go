// tree_test.go
package main

import (
	"testing"
)

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Val: value,
	}
}

func TestIsIdentical(t *testing.T) {
	tests := []struct {
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		// Both trees are empty.
		{nil, nil, true},

		// One tree is empty, and the other is not.
		{NewNode(1), nil, false},

		// Both trees have one node, but values are different.
		{NewNode(1), NewNode(2), false},

		// Both trees have one node with the same value.
		{NewNode(1), NewNode(1), true},

		// Complex trees that are identical.
		{
			&TreeNode{1, NewNode(2), NewNode(3)},
			&TreeNode{1, NewNode(2), NewNode(3)},
			true,
		},

		// Complex trees that are not identical.
		{
			&TreeNode{1, NewNode(2), NewNode(3)},
			&TreeNode{1, NewNode(3), NewNode(2)},
			false,
		},

		{
			&TreeNode{8,
				&TreeNode{4,
					&TreeNode{2, NewNode(1), NewNode(3)},
					&TreeNode{6, NewNode(5), NewNode(7)},
				},
				&TreeNode{12,
					&TreeNode{10, NewNode(9), NewNode(11)},
					&TreeNode{14, NewNode(13), NewNode(15)},
				},
			},
			&TreeNode{8,
				&TreeNode{4,
					&TreeNode{2, NewNode(1), NewNode(3)},
					&TreeNode{6, NewNode(5), NewNode(7)},
				},
				&TreeNode{12,
					&TreeNode{10, NewNode(9), NewNode(11)},
					&TreeNode{14, NewNode(13), NewNode(15)},
				},
			},
			true,
		},

		// Non-identical Trees: Changing one node's value (15 to 100) to make them different.
		{
			&TreeNode{8,
				&TreeNode{4,
					&TreeNode{2, NewNode(1), NewNode(3)},
					&TreeNode{6, NewNode(5), NewNode(7)},
				},
				&TreeNode{12,
					&TreeNode{10, NewNode(9), NewNode(11)},
					&TreeNode{14, NewNode(13), NewNode(100)}, // Changed this node's value
				},
			},
			&TreeNode{8,
				&TreeNode{4,
					&TreeNode{2, NewNode(1), NewNode(3)},
					&TreeNode{6, NewNode(5), NewNode(7)},
				},
				&TreeNode{12,
					&TreeNode{10, NewNode(9), NewNode(11)},
					&TreeNode{14, NewNode(13), NewNode(15)},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		if output := IsIdentical(test.root1, test.root2); output != test.expected {
			t.Errorf("For trees %v and %v, expected %v but got %v", test.root1, test.root2, test.expected, output)
		}
	}
}
