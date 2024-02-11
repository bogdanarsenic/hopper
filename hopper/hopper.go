package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	models "./models"
	constants "./models/constants"
	validators "./validators"
)

func start(reader io.Reader) []string {
	sc := bufio.NewReader(reader)

Retry:

	fmt.Println("Start Processing:")
	lineNumber := 0
	numOfCases := 0
	caseAdded := 0
	caseLine := 0
	currentCase := 0
	obstacleLines := -1

	var tCases []*models.TCase

	for {
		if numOfCases > 0 && caseAdded == numOfCases {
			break
		}
		line, err := sc.ReadString('\n')
		if err != nil {
			break
		}
		lineNumber++

		currentLine := strings.TrimSpace(line)
		if lineNumber == 1 {
			numOfCases, err = strconv.Atoi(currentLine)
			if err != nil || numOfCases < 1 || numOfCases > constants.MAX_GRID {
				fmt.Printf("Invalid number of test cases. Please write one number between 1 and %d!\n", constants.MAX_GRID)
				goto Retry
			}
			tCases = make([]*models.TCase, numOfCases)
		} else {
			switch caseLine {
			case 0:
				tCases, err = validators.ValidateGrid(tCases, currentCase, currentLine)
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				caseLine++
			case 1:
				tCases, err = validators.ValidatePosition(tCases, currentCase, currentLine, caseLine)
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				caseLine++
			case 2:
				obstacleLines, err = strconv.Atoi(currentLine)
				if err != nil || obstacleLines < 1 || obstacleLines > constants.MAX_GRID {
					fmt.Printf("Invalid number of obstacles. Please write one number between 1 and %d!\n", constants.MAX_GRID)
					goto Retry
				}
				caseLine++
			default:
				tCases, err = validators.ValidatePosition(tCases, currentCase, currentLine, caseLine)
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				obstacleLines--
				if obstacleLines == 0 {
					caseLine = 0
					caseAdded++
					currentCase++
				}
			}
		}
	}
	return gettingOutput(tCases)
}

func gettingOutput(tCases []*models.TCase) []string {
	output := make([]string, len(tCases))
	for i := range tCases {
		hops := processTestCase(tCases[i].Grid, tCases[i].Route, tCases[i].Obstacles)
		if hops == -1 {
			output[i] = "No solution."
		} else {
			output[i] = fmt.Sprintf("Optimal solution takes %d hops.", hops)
		}
	}
	return output
}

func processTestCase(g models.Grid, r models.Route, o []models.Obstacle) int {
	grid := make([][]bool, g.Column)

	for i := 0; i < g.Column; i++ {
		grid[i] = make([]bool, g.Row)
	}

	for _, obs := range o {
		for i := obs.X1; i <= obs.X2; i++ {
			for y := obs.Y1; y <= obs.Y2; y++ {
				grid[i][y] = true
			}
		}
	}

	minHops := findMinHops(grid, r)
	return minHops
}

func findMinHops(grid [][]bool, r models.Route) int {
	newHopper := models.HopperAt{Position: r.Start, Velocity: models.Point{X: 0, Y: 0}, Hop: 0}
	finish := r.End

	possibleHops := []*models.HopperAt{&newHopper}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for len(possibleHops) > 0 {
		currentHop := possibleHops[0]
		possibleHops = possibleHops[1:]

		if currentHop.Position == finish {
			return currentHop.Hop
		}

		for i := 0; i < 8; i++ {
			nextPoint := models.Point{X: currentHop.Position.X + currentHop.Velocity.X + dx[i], Y: currentHop.Position.Y + currentHop.Velocity.Y + dy[i]}
			if nextPoint.X < 0 || nextPoint.Y < 0 || nextPoint.X >= len(grid) || nextPoint.Y >= len(grid[0]) || grid[nextPoint.X][nextPoint.Y] {
				continue
			}
			newVelocity := models.Point{X: currentHop.Velocity.X + dx[i], Y: currentHop.Velocity.Y + dy[i]}
			if abs(newVelocity.X) <= constants.MAX_VELOCITY && abs(newVelocity.Y) <= constants.MAX_VELOCITY {
				possibleHops = append(possibleHops, &models.HopperAt{Position: nextPoint, Velocity: newVelocity, Hop: currentHop.Hop + 1})
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
	outputs := start(os.Stdin)
	for _, output := range outputs {
		fmt.Println(output)
	}
}
