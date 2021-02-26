package p0002

import (
	"github.com/SlinSo/leetcode/structure"
)

func addTwoNumbers(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	ans := &structure.ListNode{}
	pAns := ans
	val := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		pAns.Val = val % 10

		if val >= 10 {
			val = 1
		} else {
			val = 0
		}
		if l1 != nil || l2 != nil {
			l := &structure.ListNode{}
			pAns.Next = l
			pAns = pAns.Next
		}
	}

	if val == 1 {
		l := &structure.ListNode{}
		l.Val = 1

		pAns.Next = l
	}
	return ans
}
