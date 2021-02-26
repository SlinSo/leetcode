package structure

import "strconv"

func NewListNode(nums []int) *ListNode {
	l := &ListNode{}
	lTemp := l

	for i := 0; i < len(nums); i++ {
		lTemp.Val = nums[i]
		if i+1 < len(nums) {
			lTemp.Next = &ListNode{}
			lTemp = lTemp.Next
		}
	}

	return l
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "[]"
	}
	s := "["

	for l.Next != nil {
		s += strconv.Itoa(l.Val)
		s += ","
		l = l.Next
	}
	s += strconv.Itoa(l.Val)

	s += "]"
	return s
}
