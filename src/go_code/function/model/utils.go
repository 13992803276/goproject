package model

import "fmt"

var Name string = "tom"

// Operation 整数计算器
func Operation(a int, b int, op byte) int {
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
