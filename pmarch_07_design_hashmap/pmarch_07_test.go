package pmarch_07

import (
	"testing"
)

func TestMyHashMap_Get(t *testing.T) {
	type fields struct {
		m []Pair
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "not",
			fields: fields{
				m: []Pair{{1, 1}},
			},
			args: args{
				key: 0,
			},
			want: -1,
		},
		{
			name: "ok",
			fields: fields{
				m: []Pair{{1, 2}},
			},
			args: args{
				key: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyHashMap{
				m: tt.fields.m,
			}
			if got := m.Get(tt.args.key); got != tt.want {
				t.Errorf("MyHashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyHashMap_Remove(t *testing.T) {
	type fields struct {
		m []Pair
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "do nothing",
			fields: fields{
				m: []Pair{{1, 2}},
			},
			args: args{
				key: 0,
			},
			want: 1,
		},
		{
			name: "delete and empty",
			fields: fields{
				m: []Pair{{1, 2}},
			},
			args: args{
				key: 1,
			},
			want: 0,
		},
		{
			name: "delete",
			fields: fields{
				m: []Pair{{1, 2}, {2, 2}},
			},
			args: args{
				key: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyHashMap{
				m: tt.fields.m,
			}
			m.Remove(tt.args.key)
			if got := len(m.m); got != tt.want {
				t.Errorf("MyHashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
