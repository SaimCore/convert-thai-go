package convertthai

import (
	"strconv"
	"strings"
)

var numberThConstants = map[string]string{
	"0": "๐",
	"1": "๑",
	"2": "๒",
	"3": "๓",
	"4": "๔",
	"5": "๕",
	"6": "๖",
	"7": "๗",
	"8": "๘",
	"9": "๙",
	".": ".",
	",": ",",
	"-": "-",
}

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
