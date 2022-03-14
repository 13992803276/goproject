package main

import "fmt"

/**
赋值运算符：
*/
func main() {
	var a = 10
	var b = a + 1
	fmt.Printf("b的值为：%d \n", b)
	t := a
	a = b
	b = t
	fmt.Printf("a 的值为 %d 且 b的值为 %d ", a, b)
}
