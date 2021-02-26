package p0006

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "ex2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "ex3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "ex4",
			args: args{
				s:       "A",
				numRows: 3,
			},
			want: "A",
		},
		{
			name: "ex5",
			args: args{
				s:       "AB",
				numRows: 2,
			},
			want: "AB",
		},
		{
			name: "ex6",
			args: args{
				s:       "ABC",
				numRows: 2,
			},
			want: "ACB",
		},
		{
			name: "ex7",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 8,
			},
			want: "PAGYNPIARLIIHS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
