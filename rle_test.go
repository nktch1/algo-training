package main

import "testing"

func Test_rle(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				raw: "aaabbbcccd",
			},

			want: "a3b3c3d1",
		},
		{
			args: args{
				raw: "abcd",
			},

			want: "a1b1c1d1",
		},
		{
			args: args{
				raw: "a",
			},

			want: "a1",
		},
		{
			args: args{
				raw: "abcc",
			},

			want: "a1b1c2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rle(tt.args.raw); got != tt.want {
				t.Errorf("rle() = %v, want %v", got, tt.want)
			}
		})
	}
}
