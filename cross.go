package cross

import (
	"errors"
	"strconv"
	"strings"
)

//CrosspointCountString is a string wrapper for CrosspointCount
//param:
//	s => input string to be parsed to the necessary CrosspointCount params
//		 format:
//			- number separated with a space " "
//			- X Axis and Y Axis is separated by a dash "-"
//			- example = "2 4 1 2 5 3-3 2 3 1"
func CrosspointCountString(s string) (int, error) {
	splitted := strings.Split(s, "-")
	if len(splitted) != 2 {
		return -1, errors.New("invalid-input")
	}

	xAxis, err := convertStringToIntSlice(splitted[0], " ")
	if err != nil {
		return -1, err
	}
	yAxis, err := convertStringToIntSlice(splitted[1], " ")
	if err != nil {
		return -1, err
	}
	return CrosspointCount(xAxis, yAxis), nil
}

func convertStringToIntSlice(s, separator string) ([]int, error) {
	numberStringSlice := strings.Split(s, separator)
	intSlice := make([]int, 0)
	for _, numberString := range numberStringSlice {
		intVal, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, intVal)
	}
	return intSlice, nil
}

//CrosspointCount counts the number of crosspoints of vertical and horizontal lines
//params:
//	xAxis is an int array, each element specifies the length of vertical lines in (index+1) position on x Axis
//	yAxis is an int array, each element specifies the length of horizontal lines in (index+1) position on y Axis
func CrosspointCount(xAxis, yAxis []int) int {
	//Counter var
	crosspointCount := 0
	//for each line in y Axis
	for yLinePos, yLineLength := range yAxis {
		//get the number of times we need to check for crosspoint for this line
		//if the number of lines in x Axis is smaller than the length of current Y axis line
		//then we don't need to check the overflowed lines
		limit := yLineLength
		if len(xAxis) < yLineLength {
			limit = len(xAxis)
		}
		for i := 0; i < limit; i++ {
			//if i'th position is not taller than the position of current y line
			//then ignore it
			//else that is a crosspoint
			if xAxis[i] < yLinePos+1 {
				continue
			}
			crosspointCount++
		}
	}
	return crosspointCount
}
