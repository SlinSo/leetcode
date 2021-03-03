package p0013

import "testing"

func Test_romanToInt(t *testing.T) {
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
				s: "III",
			},
			want: 3,
		},
		{
			name: "ex2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "ex3",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "ex4",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "ex5",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "ex6",
			args: args{
				s: "I",
			},
			want: 1,
		},
		{
			name: "ex7",
			args: args{
				s: "MMMCMXCIX",
			},
			want: 3999,
		},
		{
			name: "ex8",
			args: args{
				s: "MDCLXIV",
			},
			want: 1664,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
