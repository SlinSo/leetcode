package p0009

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				x: -101,
			},
			want: false,
		},
		{
			name: "12321",
			args: args{
				x: 12321,
			},
			want: true,
		},
		{
			name: "123321",
			args: args{
				x: 123321,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "ex3",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "ex4",
			args: args{
				x: -101,
			},
			want: false,
		},
		{
			name: "12321",
			args: args{
				x: 12321,
			},
			want: true,
		},
		{
			name: "123321",
			args: args{
				x: 123321,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
