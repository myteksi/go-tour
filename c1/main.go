package main

import (
	"strconv"
)

// format takes 2 arguments and output a formatted string
// sample: format(1,2.0) => "formatted-int(1)-float(2.00)" the float number shall always have 2 digits after decimal point.
func format(intValue int64, floatValue float64) string {
	float64Str := strconv.FormatFloat(floatValue, 'f', 2, 64)
	return "formatted-int(" + strconv.FormatInt(intValue, 10) + ")-float(" + float64Str + ")"
}
