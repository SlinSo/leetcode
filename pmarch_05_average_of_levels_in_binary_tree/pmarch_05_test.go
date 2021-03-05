package pmarch_05

import (
	"reflect"
	"testing"

	"github.com/SlinSo/leetcode/structure"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "ex1",
			args: args{
				root: nil,
			},
			want: []float64{},
		},
		{
			name: "ex2",
			args: args{
				root: &structure.TreeNode{Val: 1},
			},
			want: []float64{1},
		},
		{
			name: "ex2",
			args: args{
				root: &structure.TreeNode{Val: 3, Left: &structure.TreeNode{Val: 9}, Right: &structure.TreeNode{Val: 20}},
			},
			want: []float64{3, 14.5},
		},
		{
			name: "ex3",
			args: args{
				root: &structure.TreeNode{Val: 3,
					Left: &structure.TreeNode{Val: 9},
					Right: &structure.TreeNode{Val: 20,
						Left:  &structure.TreeNode{Val: 15},
						Right: &structure.TreeNode{Val: 7},
					},
				},
			},
			want: []float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
