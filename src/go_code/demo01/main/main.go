package main

import (
	"fmt"
)

/*
	算数运算符讲解
*/
func main() {
	fmt.Println("hello world")
	//golang 语言的运算符
	// /运算： 运算会自动去掉小数点后面的数据，如果需要保留，参与运算的数字必须是小数
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	var n2 float32 = 10.0 / 3
	fmt.Println(n2)

	// % 运算：a % b = a - a / b * b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) //-1

	// ++ 运算 和 -- 的使用
	var i int = 10
	i++ // i = i + 1
	fmt.Println("i = ", i)
	var j int = 20
	j-- // j = j - 1
	fmt.Println("j = ", j)

	/**
	注意细节
	1.在golang中，++ 和 -- 只能独立使用，不存在--i 和 ++i 的使用方法。
	2.在golang中不能使用 j = i++的写法。
	*/
}
