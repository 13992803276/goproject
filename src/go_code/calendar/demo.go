package main

import (
	"fmt"
	"time"
)

/*
	golang 中的日期时间使用
*/
func main() {
	fmt.Println("golang 中的日期时间使用")
	now := time.Now()
	fmt.Println(now)
	fmt.Println("年 = ", now.Year())
	fmt.Println("月 = ", now.Month())
	fmt.Println("日 = ", now.Day())
	fmt.Println("时 = ", now.Hour())
	fmt.Println("分 = ", now.Minute())
	fmt.Println("秒 = ", now.Second())
	//格式化日期必须使用2006-01-02 15:04:05对应的数字进行格式组合
	formatDate := now.Format("2006-01-02 15:04:05")
	fmt.Println(formatDate)
}
