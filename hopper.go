package main

import "fmt"

type point struct {
	x, y int
}

type grid struct {
	column, row int
}

type hopperAt struct {
	position point
	velocity point
	hop      int
}

type route struct {
	start, end point
}

type obstacle struct {
	x1, x2, y1, y2 int
}

type tCase struct {
	g    grid
	r    route
	o    []obstacle
	next *tCase
}

const MAX_VELOCITY = 3

func processInput(test tCase) int {

	grid := make([][]bool, test.g.column)

	for i := 0; i < test.g.column; i++ {
		grid[i] = make([]bool, test.g.row)
	}

	for _, obs := range test.o {
		for i := obs.x1; i <= obs.x2; i++ {
			for y := obs.y1; y <= obs.y2; y++ {
				grid[i][y] = true
			}
		}
	}

	minHops := findMinHops(grid, test)
	return minHops
}

func findMinHops(grid [][]bool, test tCase) int {
	newHopper := hopperAt{}

	newHopper.position = test.r.start
	newHopper.velocity.x = 0
	newHopper.velocity.y = 0
	newHopper.hop = 0
	finish := test.r.end

	possibleHops := []*hopperAt{&newHopper}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for len(possibleHops) > 0 {
		currentHop := possibleHops[0]
		possibleHops = possibleHops[1:]

		if currentHop.position == finish {
			return currentHop.hop
		}

		for i := 0; i < 8; i++ {
			nextPoint := point{currentHop.position.x + currentHop.velocity.x + dx[i], currentHop.position.y + currentHop.velocity.y + dy[i]}
			if nextPoint.x < 0 || nextPoint.y < 0 || nextPoint.x >= len(grid) || nextPoint.y >= len(grid[0]) || grid[nextPoint.x][nextPoint.y] {
				continue
			}
			newVelocity := point{currentHop.velocity.x + dx[i], currentHop.velocity.y + dy[i]}
			if abs(newVelocity.x) <= MAX_VELOCITY && abs(newVelocity.y) <= MAX_VELOCITY {
				possibleHops = append(possibleHops, &hopperAt{nextPoint, newVelocity, currentHop.hop + 1})
			}
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	obstacles := []obstacle{{1, 4, 2, 3}}

	testInput := tCase{g: grid{5, 5}, r: route{start: point{4, 0}, end: point{4, 4}}, o: obstacles}

	hop := processInput(testInput)
	fmt.Println(hop)
}
