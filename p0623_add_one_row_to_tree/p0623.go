package p0623

import (
	"github.com/SlinSo/leetcode/structure"
)

type TreeNode = structure.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	return addNode(root, v, d, true)
}

func addNode(node *TreeNode, v int, d int, left bool) *TreeNode {
	if d == 1 {
		t := &TreeNode{Val: v}
		if left {
			t.Left = node
		} else {
			t.Right = node
		}
		return t
	} else if d == 2 {
		node.Left = addNode(node.Left, v, d-1, true)
		node.Right = addNode(node.Right, v, d-1, false)
	} else {
		if node.Left != nil {
			node.Left = addNode(node.Left, v, d-1, true)
		}
		if node.Right != nil {
			node.Right = addNode(node.Right, v, d-1, false)
		}
	}
	return node
}
