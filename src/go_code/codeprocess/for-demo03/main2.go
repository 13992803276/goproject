package main

import "fmt"

func main() {
	//打印金字塔
	/*var level int
	fmt.Println("请输入要打印的层数：")
	fmt.Scanln(&level)
	for i := 0; i <= level; i++ {
		//在打印*号前先打印空格
		for k := 0; k < level-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == level {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}*/

	//打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v ", j, i, i*j)
		}
		fmt.Println()
	}
}
