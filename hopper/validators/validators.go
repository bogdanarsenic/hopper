package validators

import (
	"fmt"
	"strconv"
	"strings"

	constants "constants"
	models "models"
)

func ValidatePosition(tCases []*models.TCase, currentCase int, currentLine string, caseLine int) ([]*models.tCase, error) {
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
				if x >= tCases[currentCase].g.column {
					err = fmt.Errorf("Number for x1 position is larger than the column number %d!\n", tCases[currentCase].g.column)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].r.start.x = x
				} else {
					obs.x1 = x
				}
			} else {
				if x >= tCases[currentCase].g.row {
					err = fmt.Errorf("Number for y1 position is larger than the row number %d!\n", tCases[currentCase].g.row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].r.start.y = x
				} else {
					obs.x2 = x
				}

			}
		} else {
			if i%2 == 0 {
				if x >= tCases[currentCase].g.column {
					err = fmt.Errorf("Number for x2 position is larger than the column number %d!\n", tCases[currentCase].g.column)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].r.end.x = x
				} else {
					obs.y1 = x
				}
			} else {
				if x >= tCases[currentCase].g.row {
					err = fmt.Errorf("Number for y2 position is larger than the row number %d!\n", tCases[currentCase].g.row)
					return nil, err
				}
				if caseLine == 1 {
					tCases[currentCase].r.end.y = x
				} else {
					obs.y2 = x
				}
			}
		}
	}
	if caseLine != 1 {
		tCases[currentCase].o = append(tCases[currentCase].o, obs)
	}
	return tCases, nil
}

func ValidateGrid(tCases []*models.tCase, currentCase int, currentLine string) ([]*models.tCase, error) {
	tCases[currentCase] = &models.tCase{}
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
			tCases[currentCase].g.column = x
		} else {
			tCases[currentCase].g.row = x
		}
	}
	return tCases, nil
}
