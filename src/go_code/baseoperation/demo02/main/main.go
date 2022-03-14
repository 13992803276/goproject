package main

import "fmt"

/**
关系运算符：关系算结果为bool类型
*/
func main() {
	var a int = 8
	var b int = 9
	fmt.Println("a > b 的结果为：", a > b)
	flag := b > a
	fmt.Println("flag 的值为：", flag)
}
