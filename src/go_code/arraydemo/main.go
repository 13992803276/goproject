package main

import (
	"fmt"
)

/**
数组的使用
数组是可以存放多个同一数据类型的数据类型。在golang中，数组是值类型，即基本数据类型。

*/
func main() {
	//1。数组的基本用法，定义，初始化值，遍历。
	var hens [6]float64
	hens[0] = 1.0
	hens[1] = 3.0
	hens[2] = 5.6
	hens[3] = 3.4
	hens[4] = 2.5
	hens[5] = 50.0
	var total = 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}
	avg := fmt.Sprintf("%.2f", total/float64(len(hens)))
	fmt.Printf("total = %v avg = %v \n", total, avg)
	//--------------------------------------------------------------------------------
	//2.定义并初始化数组的四种方式
	var arrayInt1 [3]int = [3]int{1, 2, 3}
	var arrayInt2 = [3]int{1, 2, 3}
	var arrayInt3 = [...]int{1, 2, 3}
	var names [3]string = [3]string{0: "tom", 2: "jack", 1: "mark"}
	fmt.Printf("arrayInt1 = %v arrayInt2 = %v arrayInt3 = %v \n", arrayInt1, arrayInt2, arrayInt3)
	fmt.Printf("names = %v \n", names)
	//--------------------------------------------------------------------------------
	//3.for - range 遍历数组
	for index, value := range arrayInt1 {
		fmt.Printf("arrayInt1[%v] = %v ", index, value)
	}
	//--------------------------------------------------------------------------------
	//4.数组不初始化的时候，每一个元素都有默认值
	//5.golang语言中，数组默认是值传递。如果需要引用传递，则使用指针进行传递
}
