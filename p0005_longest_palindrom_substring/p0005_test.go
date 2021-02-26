package p0005

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "ex2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "ex3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "ex4",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "ex5",
			args: args{
				s: "abccba",
			},
			want: "abccba",
		},
		{
			name: "ex6",
			args: args{
				s: "caabcbae",
			},
			want: "abcba",
		},
		{
			name: "eabcb",
			args: args{
				s: "eabcb",
			},
			want: "bcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
