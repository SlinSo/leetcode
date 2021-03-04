package pmarch_04

import (
	"reflect"
	"testing"

	"github.com/SlinSo/leetcode/structure"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		common *structure.ListNode
		headA  *structure.ListNode
		headB  *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "ex1",
			args: args{
				common: structure.NewListNode([]int{8, 4, 5}),
				headA:  structure.NewListNode([]int{4, 1}),
				headB:  structure.NewListNode([]int{5, 6, 1}),
			},
			want: structure.NewListNode([]int{8, 4, 5}),
		},
		{
			name: "ex2",
			args: args{
				common: structure.NewListNode([]int{2, 4}),
				headA:  structure.NewListNode([]int{1, 9, 1}),
				headB:  structure.NewListNode([]int{3}),
			},
			want: structure.NewListNode([]int{2, 4}),
		},
		{
			name: "ex3",
			args: args{
				common: nil,
				headA:  structure.NewListNode([]int{2, 6, 4}),
				headB:  structure.NewListNode([]int{1, 5}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.common != nil {
				tt.args.headA = structure.ListCombine(tt.args.headA, tt.args.common)
				tt.args.headB = structure.ListCombine(tt.args.headB, tt.args.common)
			}
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
