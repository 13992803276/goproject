package main

import "fmt"

/*
	slice 切片。golang语言独有的一种数据类型，切片可以理解为长度不确定的数组，他是数组的一个引用。
*/
func main() {
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	// 1.切片slice直接引用intArr的元素，从下标1开始到4。左闭右开，即不包含最有边的元素
	slice := intArr[1:4]
	fmt.Println(slice)
	fmt.Println("slice 的容量是", cap(slice))
	//-------------------------------------------------------------------------------
	//2.通过make函数创建切片,其中第三个参数容量不是必填，可以省略。切片的容量大于等于长度
	var slice1 []int = make([]int, 2, 10)
	slice1[0] = 1
	slice1[1] = 2
	fmt.Println(slice1)
	//-------------------------------------------------------------------------------
	//3.直接定义并初始化切片
	var slice2 []int = []int{1, 2, 3, 4}
	fmt.Println(slice2)
	//-------------------------------------------------------------------------------
	//4.切片append功能
	slice3 := append(slice1, 1, 2, 3, 4)
	fmt.Println(slice3)
}
