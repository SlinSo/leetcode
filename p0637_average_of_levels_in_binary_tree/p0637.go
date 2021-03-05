package p0637

import (
	"github.com/SlinSo/leetcode/structure"
)

type TreeNode = structure.TreeNode

// using BFS
func averageOfLevels(root *TreeNode) []float64 {
	var q []*TreeNode
	t := make([]*TreeNode, 0)
	sum := 0
	count := 0
	ans := []float64{}
	if root == nil {
		return ans
	}

	t = append(t, root)

	for len(t) > 0 {
		q = t
		t = []*TreeNode{}
		sum = 0
		count = 0
		for len(q) > 0 {
			sum += q[0].Val
			count++

			if q[0].Left != nil {
				t = append(t, q[0].Left)
			}
			if q[0].Right != nil {
				t = append(t, q[0].Right)
			}
			q = q[1:]
		}
		ans = append(ans, float64(sum)/float64(count))
	}

	return ans
}
