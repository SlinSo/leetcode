package structure

import "strconv"

func NewListNode(nums []int) *ListNode {
	l := &ListNode{}
	t := l

	for _, n := range nums {
		t.Next = &ListNode{Val: n}
		t = t.Next
	}

	return l.Next
}

func ListCombine(l, r *ListNode) *ListNode {
	t := l

	for t.Next != nil {
		t = t.Next
	}
	t.Next = r

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
