package structure

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	s := "V: " + strconv.Itoa(t.Val)
	if t.Left == nil {
		s += ", L: nil"
	} else {
		s += ", L: [" + t.Left.String() + "] "
	}
	if t.Right == nil {
		s += ", R: nil"
	} else {
		s += ", R: [" + t.Right.String() + "] "
	}

	return s
}
