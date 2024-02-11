package main

import (
	"testing"
)

func Test_hopper(t *testing.T) {
	type args struct {
		g grid
		r route
		o []obstacle
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				g: grid{5, 5},
				r: route{point{
					x: 4,
					y: 0,
				}, point{
					x: 4,
					y: 4,
				}},
				o: []obstacle{
					{1, 4, 2, 3},
				},
			},
			want: 7,
		},
		{
			name: "case2",
			args: args{
				g: grid{3, 3},
				r: route{point{
					x: 0,
					y: 0,
				}, point{
					x: 2,
					y: 2,
				}},
				o: []obstacle{
					{1, 1, 0, 2},
					{0, 2, 1, 1},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTestCase(tt.args.g, tt.args.r, tt.args.o); got != tt.want {
				t.Errorf("hopper() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_hopperStart(t *testing.T) {
// 	type args struct {
// 		reader io.Reader
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []string
// 	}{
// 		{
// 			name: "test",
// 			args: args{reader: strings.NewReader(`2
// 5 5
// 4 0 4 4
// 1
// 1 4 2 3
// 3 3
// 0 0 2 2
// 2
// 1 1 0 2
// 0 2 1 1

// `)},
// 			want: []string{"Optimal solution takes 7 hops.", "No solution."},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := start(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("start() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
