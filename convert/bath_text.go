package convert

import "strconv"

func BathText(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}
