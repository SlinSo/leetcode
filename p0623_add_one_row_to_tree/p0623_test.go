package p0623

import (
	"reflect"
	"testing"

	"github.com/SlinSo/leetcode/structure"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root *TreeNode
		v    int
		d    int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex1",
			args: args{
				root: &structure.TreeNode{Val: 4, Left: &structure.TreeNode{Val: 2}, Right: &structure.TreeNode{Val: 6}},
				v:    1,
				d:    2,
			},
			want: &structure.TreeNode{Val: 4,
				Left: &structure.TreeNode{Val: 1,
					Left: &structure.TreeNode{Val: 2},
				},
				Right: &structure.TreeNode{Val: 1,
					Right: &structure.TreeNode{Val: 6}}},
		},
		{
			name: "ex2",
			args: args{
				root: &structure.TreeNode{Val: 4,
					Left:  &structure.TreeNode{Val: 2},
					Right: &structure.TreeNode{Val: 6}},
				v: 1,
				d: 3,
			},
			want: &structure.TreeNode{Val: 4,
				Left: &structure.TreeNode{Val: 2,
					Left:  &structure.TreeNode{Val: 1},
					Right: &structure.TreeNode{Val: 1},
				},
				Right: &structure.TreeNode{Val: 6,
					Left:  &structure.TreeNode{Val: 1},
					Right: &structure.TreeNode{Val: 1},
				},
			},
		},
		{
			name: "ex3",
			args: args{
				root: &structure.TreeNode{Val: 4,
					Left: &structure.TreeNode{Val: 2,
						Left:  &structure.TreeNode{Val: 3},
						Right: &structure.TreeNode{Val: 1},
					},
					Right: &structure.TreeNode{Val: 6,
						Left: &structure.TreeNode{Val: 5},
					},
				},
				v: 1,
				d: 3,
			},
			want: &structure.TreeNode{Val: 4,
				Left: &structure.TreeNode{Val: 2,
					Left: &structure.TreeNode{Val: 1,
						Left: &structure.TreeNode{Val: 3},
					},
					Right: &structure.TreeNode{Val: 1,
						Right: &structure.TreeNode{Val: 1},
					},
				},
				Right: &structure.TreeNode{Val: 6,
					Left: &structure.TreeNode{Val: 1,
						Left: &structure.TreeNode{Val: 5},
					},
					Right: &structure.TreeNode{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.v, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
