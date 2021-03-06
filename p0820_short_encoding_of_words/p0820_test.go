package p0820

import "testing"

func Test_minimumLengthEncoding(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{
				words: []string{"time", "me", "bell"},
			},
			want: 10,
		},
		{
			name: "ex2",
			args: args{
				words: []string{"t"},
			},
			want: 2,
		},
		{
			name: "ex3",
			args: args{
				words: []string{"me", "time"},
			},
			want: 5,
		},
		{
			name: "ex4",
			args: args{
				words: []string{"feipyxx", "e"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
