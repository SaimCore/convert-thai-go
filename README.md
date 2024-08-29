# go-convert-thai
This package is used for converting date and number data into Thai format. (แพ็คเกจนี้ใช้สำหรับแปลงข้อมูลวันที่และตัวเลขเป็นรูปแบบภาษาไทย)
## Install (ติดตั้ง)
```
go get github.com/PurinPintakhiew/go-convert-thai
```
## How to use
```
package main

import (
  "fmt"
  "github.com/PurinPintakhiew/go-convert-thai"
)

func main(){
  num := convertthai.NumberFormat(1000, 2, true)
  fmt.Println(num) // ผลลัพธ์ ๑,๐๐๐.๐๐

  date := convertthai.DateFormat(time.Now(),"dd-mm-yyyy")
  fmt.Println(date) // ผลลัพธ์ ๒๙-๐๘-๒๐๒๔
}
```