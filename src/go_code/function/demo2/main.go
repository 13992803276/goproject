package main

import (
	"fmt"
	"lib/src/go_code/function/model"
)

/*
	函数使用细节：
	1。golang中，函数的中的变量是局部的，在函数外部无法使用。
	2。基本类型和数组类型都是值传递，不是引用传递。
	3。***函数本身也是一种数据类型。可以赋值给一个变量，通过该变量可以实现对函数的调用*** （类似把方法进行的rename）
	4。函数可以当作形参，并且可以调用。
	5。支持对返回值进行命名，此时，return不需要写具体返回值，只要名字对应就行。如test2
*/
/*func main() {
	//var age2 int = age 此处的age作用域是test方法。
	age := 30
	n1 := test(age)
	fmt.Println("main 中的age = ", age) //30
	fmt.Println("n1 = ", n1)
}*/

func main() {
	a := test
	fmt.Printf("a 的类型是 %T \n", a) // 此时说明a是一个函数类型的变量。
	res := a(20)
	fmt.Println("res = ", res)
	res1 := getSum(5, test)
	fmt.Println(res1)
	sum, sub := test2(34, 20)
	fmt.Println("sum = ", sum, "sub =", sub)
}
func getSum(n int, n2 func(int) int) int { //将函数作为参数进行传递和调用
	return n + n2(n)
}
func test(age int) int {
	return model.Operation(3, 4, '+') + age
}

//返回值命名
func test2(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
