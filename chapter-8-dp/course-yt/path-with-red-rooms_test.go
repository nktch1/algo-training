package course_yt

import "testing"

func Test_maxPriceOnTheWayWithRedRooms(t *testing.T) {
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
				k: 3,
			},
			want: 20,
		},
		{
			args: args{
				i: 1,
				j: 1,
				k: 2,
			},
			want: 9,
		},
		{
			args: args{
				i: 3,
				j: 3,
				k: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPriceOnTheWayWithRedRooms(tt.args.i, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("maxPriceOnTheWayWithRedRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
