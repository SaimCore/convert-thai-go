package convertthai

import (
	"fmt"
	"math"

	"strconv"
	"strings"
)

var bathNumbers = []string{
	"",
	"หนึ่ง",
	"สอง",
	"สาม",
	"สี่",
	"ห้า",
	"หก",
	"เจ็ด",
	"แปด",
	"เก้า",
}

var bathUnits = []string{
	"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน",
}

func numberToWord(numberArr []string) string {
	var ret string
	length := len(numberArr)

	for i, number := range numberArr {
		digit, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			pt := length - i - 1

			if digit != 0 {
				ret += bathNumbers[digit] + bathUnits[pt%6]
				if pt%6 == 0 && pt != 0 {
					ret += bathUnits[6]
				}
			}
		}
	}

	ret = strings.ReplaceAll(ret, "หนึ่งสิบ", "สิบ")
	ret = strings.ReplaceAll(ret, "สองสิบ", "ยี่สิบ")
	ret = strings.ReplaceAll(ret, "สิบหนึ่ง", "สิบเอ็ด")

	return ret

}

func ThaiBath(number float64) string {
	if number == 0 {
		return "ศูนย์บาทถ้วน"
	}

	n := math.Floor(number*100) / 100
	numberStr := strconv.FormatFloat(n, 'f', 2, 64)
	numberArr := strings.Split(numberStr, ".")

	integerArr := strings.Split(numberArr[0], "")
	decimalArr := strings.Split(numberArr[1], "")

	var bathText string
	length := len(integerArr)
	if length > 6 {
		millions := integerArr[0 : length-6]
		bathText += numberToWord(millions) + bathUnits[6]
	} else {
		bathText += numberToWord(integerArr)
	}

	if integerArr[0] != "0" {
		bathText += "บาท"
	}

	if len(decimalArr) > 0 && (decimalArr[0] != "0" && decimalArr[1] != "0") {
		bathText += numberToWord(decimalArr) + "สตางค์"
	} else {
		bathText += "ถ้วน"
	}

	if number < 0 {
		return "ลบ" + bathText
	}

	return bathText
}
