package chapter_8_dp

import "testing"

func Test_maxCommonSubstring(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: 0,
		},
		{
			args: args{
				s1: "abcdafc",
				s2: "abcdafc",
			},
			want: 0,
		},
		{
			args: args{
				s1: "abcdrfk",
				s2: "abcdafc",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCommonSubstring(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("maxCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: 3,
		},
		{
			args: args{
				s1: "abcdafc",
				s2: "abcdafc",
			},
			want: 4,
		},
		{
			args: args{
				s1: "abcdrfk",
				s2: "abcdafc",
			},
			want: 4,
		},	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCommonSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("maxCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}