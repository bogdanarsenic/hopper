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

// start processing the inputs
func startProcess(reader io.Reader) []string {
	sc := bufio.NewReader(reader)

	// in case there is an invalid input, start everything again from this label
Retry:

	fmt.Println("Start Processing: ")
	lineNumber := 0
	numOfCases := 0
	caseAdded := 0
	caseLine := 0
	currentCase := 0
	obstacleLines := -1

	var tCases []*models.TCase

	for {
		// checking if the process is completed
		if numOfCases > 0 && caseAdded == numOfCases {
			break
		}
		// reading the input
		line, err := sc.ReadString('\n')
		if err != nil {
			break
		}
		lineNumber++

		currentLine := strings.TrimSpace(line)
		// process input number of test cases
		if lineNumber == 1 {
			numOfCases, err = strconv.Atoi(currentLine)
			if err != nil || numOfCases < 1 || numOfCases > constants.MAX_GRID {
				fmt.Printf("Invalid number of test cases. Please write one number between 1 and %d!\n", constants.MAX_GRID)
				goto Retry
			}
			tCases = make([]*models.TCase, numOfCases)
		} else {
			switch caseLine {
			// check if the input for the grid columns and rows is valid
			case 0:
				tCases, err = validators.ValidateGrid(tCases, currentCase, currentLine)
				// if there is an error, go back to label Retry
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				caseLine++
			// check if the start and end position is valid
			case 1:
				tCases, err = validators.ValidatePosition(tCases, currentCase, currentLine, caseLine)
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				caseLine++
			// process input number of obstacles
			case 2:
				obstacleLines, err = strconv.Atoi(currentLine)
				if err != nil || obstacleLines < 1 || obstacleLines > constants.MAX_GRID {
					fmt.Printf("Invalid number of obstacles. Please write one number between 1 and %d!\n", constants.MAX_GRID)
					goto Retry
				}
				caseLine++
			// checking if the inputs for obstacles are valid
			default:
				tCases, err = validators.ValidatePosition(tCases, currentCase, currentLine, caseLine)
				if err != nil {
					fmt.Println(err)
					goto Retry
				}
				// going to next obstacle positions and potential next test case
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
	// go through all the test cases individually and print the result
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
	//initializing grid
	grid := make([][]bool, g.Column)
	for i := 0; i < g.Column; i++ {
		grid[i] = make([]bool, g.Row)
	}

	//set obstacles on the grid
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
	// initialize hopper and its properties
	newHopper := models.HopperAt{Position: r.Start, Velocity: models.Point{X: 0, Y: 0}, Hop: 0}
	finish := r.End

	possibleHops := []*models.HopperAt{&newHopper}

	// initial velocity movements for the hopper
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for len(possibleHops) > 0 {
		// taking the first hopper
		currentHop := possibleHops[0]
		possibleHops = possibleHops[1:]

		if currentHop.Position == finish {
			return currentHop.Hop
		}

		// 8 possible hopper movements on the grid
		for i := 0; i < 8; i++ {
			// getting the next position and checking if it's valid
			nextPoint := models.Point{X: currentHop.Position.X + currentHop.Velocity.X + dx[i], Y: currentHop.Position.Y + currentHop.Velocity.Y + dy[i]}
			if nextPoint.X < 0 || nextPoint.Y < 0 || nextPoint.X >= len(grid) || nextPoint.Y >= len(grid[0]) || grid[nextPoint.X][nextPoint.Y] {
				continue
			}
			// calculating new velocity and updating the queue
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
	outputs := startProcess(os.Stdin)
	for _, output := range outputs {
		fmt.Println(output)
	}
}
