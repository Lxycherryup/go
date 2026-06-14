package main

import (
	"fmt"
	"time"
)

//时间和日期都是编程中经常会用到的
//time包提供了时间的现实和测量用的函数。日历的计算采用的是公历
// time.Time类型表示时间
func timeDemo() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()
	fmt.Println(timestamp1)
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp2)
	timestamp3 := now.UnixMicro()
	fmt.Println(timestamp3)
	//可以使用time.Unix()函数将时间出转换为时间格式
	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)

}
func main() {

	// timeDemo()
	// timestampDemo()
}
