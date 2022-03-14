package main

import (
	"fmt"
)

/*
	循环分支控制
	1.golang 新增了for-range的循环方式，但是没有while 和 do while的语法。
	2.多层循环，嵌套循环。

*/
func main() {
	// 方式1：
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}
	// 方式2：
	j := 0
	for j < 10 {
		fmt.Println("hello go")
		j++
	}
	//字符串遍历方式1
	var str string = "hello world 你好世界"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	//字符串遍历方式2 for-range方式
	str = "xulegll北京"
	for index, val := range str {
		fmt.Printf("%d - %c \n", index, val)
	}
}
