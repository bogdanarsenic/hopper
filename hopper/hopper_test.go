package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	models "./models"
)

func Test_hopper(t *testing.T) {
	type args struct {
		Grid      models.Grid
		Route     models.Route
		Obstacles []models.Obstacle
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				Grid: models.Grid{5, 5},
				Route: models.Route{models.Point{
					X: 4,
					Y: 0,
				}, models.Point{
					X: 4,
					Y: 4,
				}},
				Obstacles: []models.Obstacle{
					{1, 4, 2, 3},
				},
			},
			want: 7,
		},
		{
			name: "case2",
			args: args{
				Grid: models.Grid{3, 3},
				Route: models.Route{models.Point{
					X: 0,
					Y: 0,
				}, models.Point{
					X: 2,
					Y: 2,
				}},
				Obstacles: []models.Obstacle{
					{1, 1, 0, 2},
					{0, 2, 1, 1},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTestCase(tt.args.Grid, tt.args.Route, tt.args.Obstacles); got != tt.want {
				t.Errorf("hopper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hopperStart(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{reader: strings.NewReader(`2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1

`)},
			want: []string{"Optimal solution takes 7 hops.", "No solution."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := start(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("start() = %v, want %v", got, tt.want)
			}
		})
	}
}
