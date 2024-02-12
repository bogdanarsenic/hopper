package validators

import (
	"fmt"
	"reflect"
	"testing"

	models "../models"
)

func Test_ValidatePositionFailed(t *testing.T) {
	tCases := make([]*models.TCase, 1)

	tCases[0] = &models.TCase{Grid: models.Grid{Row: 30, Column: 30}}

	type args struct {
		tCases      []*models.TCase
		currentCase int
		currentLine string
		caseLine    int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "test1",
			args: args{tCases: tCases, currentCase: 0, currentLine: "4 1 22", caseLine: 1},
			want: fmt.Errorf("Please write in the format 'x1 x2 y1 y2' \n"),
		},
		{
			name: "test2",
			args: args{tCases: tCases, currentCase: 0, currentLine: "s 3 s 2", caseLine: 1},
			want: fmt.Errorf("Please enter the number and it cannot be less then 0!"),
		},
		{
			name: "test3",
			args: args{tCases: tCases, currentCase: 0, currentLine: "31 5 3 5", caseLine: 1},
			want: fmt.Errorf("x1 position has to be smaller than the number of columns - 30!\n"),
		},
		{
			name: "test4",
			args: args{tCases: tCases, currentCase: 0, currentLine: "2 32 5 3", caseLine: 1},
			want: fmt.Errorf("y1 position has to be smaller than the number of rows - 30!\n"),
		},
		{
			name: "test5",
			args: args{tCases: tCases, currentCase: 0, currentLine: "5 6 33 30", caseLine: 1},
			want: fmt.Errorf("x2 position has to be smaller than the number of columns - 30!\n"),
		},
		{
			name: "test6",
			args: args{tCases: tCases, currentCase: 0, currentLine: "5 6 25 33", caseLine: 1},
			want: fmt.Errorf("y2 position has to be smaller than the number of rows - 30!\n"),
		},
		{
			name: "test7",
			args: args{tCases: tCases, currentCase: 0, currentLine: "-1 5 0 4", caseLine: 1},
			want: fmt.Errorf("Please enter the number and it cannot be less then 0!"),
		},
		{
			name: "test7",
			args: args{tCases: tCases, currentCase: 0, currentLine: "-1 5 0 4", caseLine: 0},
			want: fmt.Errorf("Please enter the number and it cannot be less then 0!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ValidatePosition(tt.args.tCases, tt.args.currentCase, tt.args.currentLine, tt.args.caseLine); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("start() = %s, want %s", err, tt.want)
			}
		})
	}
}
func Test_ValidatePositionSuccess(t *testing.T) {
	tCases := make([]*models.TCase, 1)

	tCases[0] = &models.TCase{Grid: models.Grid{Row: 30, Column: 30}}
	type args struct {
		tCases      []*models.TCase
		currentCase int
		currentLine string
		caseLine    int
	}
	tests := []struct {
		name string
		args args
		want models.TCase
	}{
		{
			name: "test1",
			args: args{tCases: tCases, currentCase: 0, currentLine: "29 25 5 20", caseLine: 1},
			want: models.TCase{Grid: models.Grid{Column: 30, Row: 30}, Route: models.Route{Start: models.Point{X: 29, Y: 25}, End: models.Point{X: 5, Y: 20}}},
		},
		{
			name: "test2",
			args: args{tCases: tCases, currentCase: 0, currentLine: "5 10 15 10", caseLine: 0},
			want: models.TCase{Grid: models.Grid{Column: 30, Row: 30}, Obstacles: []models.Obstacle{{X1: 5, X2: 10, Y1: 15, Y2: 10}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ValidatePosition(tt.args.tCases, tt.args.currentCase, tt.args.currentLine, tt.args.caseLine)

			if tt.args.caseLine == 1 {
				if !reflect.DeepEqual(got[tt.args.currentCase].Route.Start.X, tt.want.Route.Start.X) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Route.Start.Y, tt.want.Route.Start.Y) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Route.End.X, tt.want.Route.End.X) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Route.End.Y, tt.want.Route.End.Y) {
					t.Errorf("start() = %v, want %v", got[tt.args.currentCase].Route, tt.want.Route)
				}
			} else {
				if !reflect.DeepEqual(got[tt.args.currentCase].Obstacles[0].X1, tt.want.Obstacles[0].X1) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Obstacles[0].X2, tt.want.Obstacles[0].X2) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Obstacles[0].Y1, tt.want.Obstacles[0].Y1) ||
					!reflect.DeepEqual(got[tt.args.currentCase].Obstacles[0].Y2, tt.want.Obstacles[0].Y2) {
					t.Errorf("start() = %v, want %v", got[tt.args.currentCase].Obstacles[0], tt.want.Obstacles[0])
				}
			}
		})
	}
}

func Test_ValidateGridFailed(t *testing.T) {
	tCases := make([]*models.TCase, 1)
	type args struct {
		tCases      []*models.TCase
		currentCase int
		currentLine string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "test1",
			args: args{tCases: tCases, currentCase: 0, currentLine: "zzs"},
			want: fmt.Errorf("Please write in the format 'row column' for the grid numbers\n"),
		},
		{
			name: "test2",
			args: args{tCases: tCases, currentCase: 0, currentLine: "0 5"},
			want: fmt.Errorf("Grid columns and rows need to be between numbers 1 and 30!\n"),
		},
		{
			name: "test3",
			args: args{tCases: tCases, currentCase: 0, currentLine: "31 5"},
			want: fmt.Errorf("Grid columns and rows need to be between numbers 1 and 30!\n"),
		},
		{
			name: "test4",
			args: args{tCases: tCases, currentCase: 0, currentLine: " "},
			want: fmt.Errorf("Grid columns and rows need to be between numbers 1 and 30!\n"),
		},
		{
			name: "test5",
			args: args{tCases: tCases, currentCase: 0, currentLine: "5 6 3"},
			want: fmt.Errorf("Please write in the format 'row column' for the grid numbers\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ValidateGrid(tt.args.tCases, tt.args.currentCase, tt.args.currentLine); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("start() = %s, want %s", err, tt.want)
			}
		})
	}
}
func Test_ValidateGridSuccess(t *testing.T) {
	type args struct {
		tCases      []*models.TCase
		currentCase int
		currentLine string
	}
	tests := []struct {
		name string
		args args
		want models.TCase
	}{
		{
			name: "test1",
			args: args{tCases: make([]*models.TCase, 1), currentCase: 0, currentLine: "30 30"},
			want: models.TCase{Grid: models.Grid{Column: 30, Row: 30}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ValidateGrid(tt.args.tCases, tt.args.currentCase, tt.args.currentLine); !reflect.DeepEqual(got[tt.args.currentCase].Grid.Column, tt.want.Grid.Column) {
				t.Errorf("start() = %d, want %d", got[tt.args.currentCase].Grid.Column, tt.want.Grid.Column)
			}
		})
	}
}
