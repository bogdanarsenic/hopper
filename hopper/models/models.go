package models

type Point struct {
	X, Y int
}

type Grid struct {
	Column, Row int
}

type HopperAt struct {
	Position Point
	Velocity Point
	Hop      int
}

type Route struct {
	Start, End Point
}

type Obstacle struct {
	X1, X2, Y1, Y2 int
}

type TCase struct {
	Grid      Grid
	Route     Route
	Obstacles []Obstacle
}
