package converterlib

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToFloat64(myString string) float64 {

	retFloat, err := strconv.ParseFloat(strings.Trim(myString, "\n"), 64)

	if err != nil {
		fmt.Println("ERROR: StringToFloat( ", myString, " ): ", err)
	}

	return retFloat
}
