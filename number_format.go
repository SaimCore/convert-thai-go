package convertthai

import (
	"strconv"
	"strings"
)

func formatNumberWithComma(numberStr string) string {
	var result []string

	for i, c := range numberStr {
		if i > 0 && (len(numberStr)-i)%3 == 0 {
			result = append(result, ",")
		}
		result = append(result, string(c))
	}

	return strings.Join(result, "")
}

func NumberFormat(number float64, decimals int, thousandsSeparator bool) string {
	numberStr := strconv.FormatFloat(number, 'f', decimals, 64)

	if thousandsSeparator {
		integerArr := strings.Split(numberStr, ".")
		if len(integerArr) == 2 {
			formattedNumber := formatNumberWithComma(integerArr[0])
			numberStr = formattedNumber + "." + integerArr[1]
		}
	}

	var numberArr []string
	for _, digit := range numberStr {
		numberArr = append(numberArr, string(numberThConstants[string(digit)]))
	}

	result := strings.Join(numberArr, "")

	return result
}
