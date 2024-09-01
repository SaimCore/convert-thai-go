# go-convert-thai ✨
It provides functions to format dates according to Thai conventions and convert numbers into Thai text, making it useful for applications that need to display localized information in the Thai language. (โดยมีฟังก์ชันที่ช่วยในการจัดรูปแบบวันที่ตามมาตรฐานไทย และแปลงตัวเลขเป็นข้อความภาษาไทย ทำให้เหมาะสำหรับแอปพลิเคชันที่ต้องการแสดงข้อมูลที่เป็นภาษาไทยโดยเฉพาะ)
## Install (ติดตั้ง) 🛠️
```
go get github.com/PurinPintakhiew/go-convert-thai
```
## How to use (วิธีใช้งาน) 💡
```
package main

import (
  "fmt"
  "time"
  "github.com/PurinPintakhiew/go-convert-thai"
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
