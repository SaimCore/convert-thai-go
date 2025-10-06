# convert-thai-go ✨
This package helps to format dates according to Thai standards and convert numbers to Thai text, making it suitable for applications that need to display data in Thai specifically. (แพคเกจนี้ช่วยจัดรูปแบบวันที่ตามมาตรฐานไทยและแปลงตัวเลขให้เป็นข้อความภาษาไทย ทำให้เหมาะกับแอพพลิเคชันที่ต้องการแสดงข้อมูลเป็นภาษาไทยโดยเฉพาะ)

[![Go Report Card](https://goreportcard.com/badge/SaimCore/convert-thai-go)](https://goreportcard.com/report/SaimCore/convert-thai-go) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/SaimCore/convert-thai-go/blob/main/LICENSE)

## Install (ติดตั้ง) 🛠️
```
go get github.com/SaimCore/convert-thai-go
```
## How to use (วิธีใช้งาน) 💡
```
package main

import (
  "fmt"
  "time"
  "github.com/SaimCore/convert-thai-go"
)

func main(){
  // แปลงเลขอารบิกเป็นเลขไทย
  num := convertthai.NumberFormat(1000, 2, true)
  fmt.Println(num) // ผลลัพธ์ ๑,๐๐๐.๐๐

  // แปลงวันที่เป็นรูปแบบไทย
  date := convertthai.DateFormat(time.Now(),"dd-mm-yyyy")
  fmt.Println(date) // ผลลัพธ์ ๒๙-๐๘-๒๐๒๔

  // แปลงจำนวนเงินเป็นรูปแบบไทยบาท
  thaiBath := convertthai.ThaiBath(1000)
  fmt.Println(thaiBath) // ผลลัพธ์ หนึ่งพันบาทถ้วน
}
```
