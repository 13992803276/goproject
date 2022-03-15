package main

import (
	"fmt"
	"lib/src/go_code/function/model"
)

/*
	golang 函数入门
	1.golang函数可以有多个返回值。
	2.golang函数的返回值写在参数后面，多个返回值时，用（int， float）形式
	3.语法：func 函数名 (形参列表) （返回值列表）{
			代码块
			return 返回值列表
	}
	4.入参中，参数名在前，参数类型在后，这点和java有区别
	5.包中的函数、变量首字母大写表示共有，小写表示私有
	6.包引用后，支持别名，一旦起别名后，原来的包名失效，不可用
	7.一个包中不能定义相同名称的函数。即使是在不同的go文件中
*/
func main() {
	fmt.Println(model.Name)
	res := operation(2, 3, '*')
	result := model.Operation(1, 3, '-')
	fmt.Println(res, result)
}

func operation(a int, b int, op byte) int {
	res := 0
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a - b
	default:
		fmt.Println("输入有误。。。")
	}
	return res
}
