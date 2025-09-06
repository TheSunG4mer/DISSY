package helper

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2 plus 2 equals 4", args{a: 2, b: 2}, 4},
		{
			name: "0 plus 0 equals 0",
			args: args{a: 0, b: 0},
			want: 0,
		},
		{
			name: "negative plus positive",
			args: args{a: -3, b: 5},
			want: 2,
		},
		{
			name: "negative plus negative",
			args: args{a: -2, b: -2},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2 plus 2 equals 4", args{a: 2, b: 2}, 4},
		{
			name: "0 plus 0 equals 0",
			args: args{a: 0, b: 0},
			want: 0,
		},
		{
			name: "negative plus positive",
			args: args{a: -3, b: 5},
			want: 2,
		},
		{
			name: "negative plus negative",
			args: args{a: -2, b: -2},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
