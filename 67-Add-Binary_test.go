package main

import "testing"

func Test_sumBinaryString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s1: "11",
				s2: "10",
			},
			want: "101",
		},
		{
			args: args{
				s1: "0",
				s2: "10",
			},
			want: "10",
		},
		{
			args: args{
				s1: "0",
				s2: "0",
			},
			want: "0",
		},
		{
			args: args{
				s1: "10",
				s2: "10",
			},
			want: "100",
		},
		{
			args: args{
				s1: "0",
				s2: "110100",
			},
			want: "110100",
		},
		{
			args: args{
				s1: "1010",
				s2: "1011",
			},
			want: "10101",
		},
		{
			args: args{
				s1: "1111",
				s2: "1111",
			},
			want: "11110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumBinaryStrings(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("sumBinaryStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
