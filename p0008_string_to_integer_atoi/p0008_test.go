package p0008

import (
	"math"
	"strconv"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "ex2",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "ex3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "ex4",
			args: args{
				s: "word and 987",
			},
			want: 0,
		},
		{
			name: "ex5",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "ex6",
			args: args{
				s: "-+12",
			},
			want: 0,
		},
		{
			name: "ex7",
			args: args{
				s: "--12",
			},
			want: 0,
		},
		{
			name: "ex8",
			args: args{
				s: "++12",
			},
			want: 0,
		},
		{
			name: "ex9",
			args: args{
				s: "+",
			},
			want: 0,
		},
		{
			name: "ex10",
			args: args{
				s: "-",
			},
			want: 0,
		},
		{
			name: "ex11",
			args: args{
				s: "0000-42",
			},
			want: 0,
		},
		{
			name: "ex12",
			args: args{
				s: "-42-",
			},
			want: -42,
		},
		{
			name: "ex13",
			args: args{
				s: "-42+",
			},
			want: -42,
		},
		{
			name: "ex14",
			args: args{
				s: strconv.Itoa(math.MaxInt32),
			},
			want: math.MaxInt32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
