package p0002_add_two_numbers

import (
	"reflect"
	"testing"

	"github.com/SlinSo/leetcode/structure"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *structure.ListNode
		l2 *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "ex1",
			args: args{
				l1: structure.NewListNode([]int{2, 4, 3}),
				l2: structure.NewListNode([]int{5, 6, 4}),
			},
			want: structure.NewListNode([]int{7, 0, 8}),
		},
		{
			name: "ex2",
			args: args{
				l1: structure.NewListNode([]int{0}),
				l2: structure.NewListNode([]int{0}),
			},
			want: structure.NewListNode([]int{0}),
		},
		{
			name: "ex3",
			args: args{
				l1: structure.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: structure.NewListNode([]int{9, 9, 9, 9}),
			},
			want: structure.NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
