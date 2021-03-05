package pmarch_05

import (
	"github.com/SlinSo/leetcode/structure"
)

type TreeNode = structure.TreeNode

type Pair struct {
	Elems int
	Sum   int
}

func averageOfLevels(root *TreeNode) []float64 {
	sums := make(map[int]Pair)

	if root == nil {
		return []float64{}
	}

	calcSum(sums, 0, root)

	ans := make([]float64, len(sums))
	for k, v := range sums {
		ans[k] = float64(v.Sum) / float64(v.Elems)
	}

	return ans
}

func calcSum(sums map[int]Pair, level int, node *TreeNode) {
	if node == nil {
		return
	}

	if p, ok := sums[level]; ok {
		p.Elems++
		p.Sum += node.Val
		sums[level] = p
	} else {
		p := Pair{}
		p.Elems = 1
		p.Sum = node.Val
		sums[level] = p
	}

	calcSum(sums, level+1, node.Left)
	calcSum(sums, level+1, node.Right)
}
