package pmarch_08

import "testing"

func Test_removePalindromeSub(t *testing.T) {
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
				s: "ababa",
			},
			want: 1,
		},
		{
			name: "ex2",
			args: args{
				s: "abb",
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				s: "baabb",
			},
			want: 2,
		},
		{
			name: "ex4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "ex5",
			args: args{
				s: "bbaabaaa",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePalindromeSub(tt.args.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
