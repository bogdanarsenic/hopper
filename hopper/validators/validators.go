package validators

import (
	"fmt"
	"strconv"
	"strings"

	models "../models"
	constants "../models/constants"
)

// Checking if the inputs for positions (routes and obstacles) correct
func ValidatePosition(tCases []*models.TCase, currentCase int, currentLine string, caseLine int) ([]*models.TCase, error) {
	givenPosition := strings.Split(currentLine, " ")
	obs := models.Obstacle{}

	var err error
	// has to be 4 positions
	if len(givenPosition) != 4 {
		err = fmt.Errorf("Please write in the format 'x1 x2 y1 y2' \n")
		return nil, err
	}

	for i := range givenPosition {
		x, err := strconv.Atoi(givenPosition[i])
		if err != nil || x < 0 {
			err = fmt.Errorf("Please enter the number and it cannot be less then 0!")
			return nil, err
		}
		// the first two numbers
		if i < 2 {
			// the first position
			if i%2 == 0 {
				// checking if the number is valid - less or equal to number of columns
				if x >= tCases[currentCase].Grid.Column {
					err = fmt.Errorf("x1 position has to be smaller than the number of columns - %d!\n", tCases[currentCase].Grid.Column)
					return nil, err
				}
				// this means that we are checking the routes - start and end position
				if caseLine == 1 {
					tCases[currentCase].Route.Start.X = x
				} else {
					// this is for the obstacles
					// if start position is the same as the obstacle, then we won't consider that obstacle
					if tCases[currentCase].Route.Start.X != x {
						obs.X1 = x
					}
				}
			} else {
				// same logic for the second element
				if x >= tCases[currentCase].Grid.Row {
					err = fmt.Errorf("y1 position has to be smaller than the number of rows - %d!\n", tCases[currentCase].Grid.Row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.Start.Y = x
				} else {
					if tCases[currentCase].Route.Start.Y != x {
						obs.X2 = x
					}
				}
			}
		} else {
			// same logic for the third and fourth element
			if i%2 == 0 {
				if x >= tCases[currentCase].Grid.Column {
					err = fmt.Errorf("x2 position has to be smaller than the number of columns - %d!\n", tCases[currentCase].Grid.Column)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.End.X = x
				} else {
					if tCases[currentCase].Route.End.X != x {
						obs.Y1 = x
					}
				}
			} else {
				if x >= tCases[currentCase].Grid.Row {
					err = fmt.Errorf("y2 position has to be smaller than the number of rows - %d!\n", tCases[currentCase].Grid.Row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.End.Y = x
				} else {
					if tCases[currentCase].Route.End.Y != x {
						obs.Y2 = x
					}
				}
			}
		}
	}
	// if it's obstacle checking, we are adding new ones
	if caseLine != 1 {
		tCases[currentCase].Obstacles = append(tCases[currentCase].Obstacles, obs)
	}
	return tCases, nil
}

// checking if the grid is valid
func ValidateGrid(tCases []*models.TCase, currentCase int, currentLine string) ([]*models.TCase, error) {
	tCases[currentCase] = &models.TCase{}
	var err error
	givenGrid := strings.Split(currentLine, " ")

	// has to have column and row number
	if len(givenGrid) != 2 {
		err = fmt.Errorf("Please write in the format 'row column' for the grid numbers\n")
		return nil, err
	}
	for i := range givenGrid {
		x, err := strconv.Atoi(givenGrid[i])
		// checking if the numbers are valid
		if err != nil || x > constants.MAX_GRID || x < 1 {
			err = fmt.Errorf("Grid columns and rows need to be between numbers 1 and %d!\n", constants.MAX_GRID)
			return nil, err
		}
		// the first value is for the columns, the second for the rows
		if i%2 == 0 {
			tCases[currentCase].Grid.Column = x
		} else {
			tCases[currentCase].Grid.Row = x
		}
	}
	return tCases, nil
}
