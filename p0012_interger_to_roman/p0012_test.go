package p0012

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "ex2",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "ex3",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "ex4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "ex5",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "ex6",
			args: args{
				num: 1,
			},
			want: "I",
		},
		{
			name: "ex7",
			args: args{
				num: 3999,
			},
			want: "MMMCMXCIX",
		},
		{
			name: "ex8",
			args: args{
				num: 1666,
			},
			want: "MDCLXVI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
