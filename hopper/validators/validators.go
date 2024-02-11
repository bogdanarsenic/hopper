package validators

import (
	"fmt"
	"strconv"
	"strings"

	models "../models"
	constants "../models/constants"
)

func ValidatePosition(tCases []*models.TCase, currentCase int, currentLine string, caseLine int) ([]*models.TCase, error) {
	givenRoute := strings.Split(currentLine, " ")
	obs := models.Obstacle{}

	var err error
	if len(givenRoute) != 4 {
		err = fmt.Errorf("Please write in the format 'x1 x2 y1 y2' \n")
		return nil, err
	}

	for i := range givenRoute {
		x, err := strconv.Atoi(givenRoute[i])
		if err != nil {
			err = fmt.Errorf("Please enter the number!")
			return nil, err
		}
		if i < 2 {
			if i%2 == 0 {
				if x >= tCases[currentCase].Grid.Column {
					err = fmt.Errorf("Number for x1 position is larger than the column number %d!\n", tCases[currentCase].Grid.Column)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.Start.X = x
				} else {
					obs.X1 = x
				}
			} else {
				if x >= tCases[currentCase].Grid.Row {
					err = fmt.Errorf("Number for y1 position is larger than the row number %d!\n", tCases[currentCase].Grid.Row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.Start.Y = x
				} else {
					obs.X2 = x
				}

			}
		} else {
			if i%2 == 0 {
				if x >= tCases[currentCase].Grid.Column {
					err = fmt.Errorf("Number for x2 position is larger than the column number %d!\n", tCases[currentCase].Grid.Column)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.End.X = x
				} else {
					obs.Y1 = x
				}
			} else {
				if x >= tCases[currentCase].Grid.Row {
					err = fmt.Errorf("Number for y2 position is larger than the row number %d!\n", tCases[currentCase].Grid.Row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].Route.End.Y = x
				} else {
					obs.Y2 = x
				}
			}
		}
	}
	if caseLine != 1 {
		tCases[currentCase].Obstacles = append(tCases[currentCase].Obstacles, obs)
	}
	return tCases, nil
}

func ValidateGrid(tCases []*models.TCase, currentCase int, currentLine string) ([]*models.TCase, error) {
	tCases[currentCase] = &models.TCase{}
	var err error
	givenGrid := strings.Split(currentLine, " ")
	if len(givenGrid) != 2 {
		err = fmt.Errorf("Please write in the format 'row column' for the grid numbers\n")
		return nil, err
	}
	for i := range givenGrid {
		x, err := strconv.Atoi(givenGrid[i])
		if err != nil || x > constants.MAX_GRID || x < 1 {
			err = fmt.Errorf("Grid columns and rows need to be between numbers 1 and %d!\n", constants.MAX_GRID)
			return nil, err
		}
		if i%2 == 0 {
			tCases[currentCase].Grid.Column = x
		} else {
			tCases[currentCase].Grid.Row = x
		}
	}
	return tCases, nil
}
