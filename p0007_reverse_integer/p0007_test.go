package p0007

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "ex2",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "ex3",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "ex4",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "ex5",
			args: args{
				x: math.MaxInt32 >> 1,
			},
			want: 0,
		},
		{
			name: "ex6",
			args: args{
				x: -2147483412,
			},
			want: -2143847412,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
