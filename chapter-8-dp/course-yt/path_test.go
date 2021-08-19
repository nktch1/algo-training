package course_yt

import (
	"testing"
)

func Test_maxPriceOnTheWay(t *testing.T) {
	type args struct {
		i int
		j int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				i: 0,
				j: 0,
				k: 0,
			},
			want: 20,
		},
		{
			args: args{
				i: 1,
				j: 1,
				k: 0,
			},
			want: 9,
		},
		{
			args: args{
				i: 3,
				j: 3,
				k: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPriceOnTheWay(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("maxPriceOnTheWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
