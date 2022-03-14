package main

import "fmt"

/**
逻辑运算符：连接多个关系运算符的操作，返回值也为bool类型，一版用于流程控制
		  逻辑&& 也叫短路与，即第一个为false时，不再判断后面的条件
          逻辑｜｜ 也叫短路或，即第一个为true时，不再执行后面的条件。
*/
func main() {
	var a int = 10
	var b int = 20
	var c int = 30
	flag := a > b || b > c && a > c
	fmt.Println("flag = ", flag)

	fmt.Println(!(a < b))

	if a > 5 && a < 20 {
		fmt.Println("OK1")
	}
}
