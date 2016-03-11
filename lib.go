package converterlib

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToFloat64(myString string) float64 {

	retFloat, err := strconv.ParseFloat(strings.Trim(myString, "\n"), 64)

	if err != nil {
		fmt.Println("ERROR: StringToFloat64( ", myString, " ): ", err)
	}

	return retFloat
}

func Float64ToString2DecimalPlaces(myFloat float64) string {

	// Format the float64 to 2 decimal places
	return strconv.FormatFloat(myFloat, 'f', 2, 64)
}
