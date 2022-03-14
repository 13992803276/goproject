package main

import "fmt"

func main() {
	//假如还有97天放假，问还有多少星期多少天。
	var days int = 97
	fmt.Println("离放假还有", days/7, "个星期", days%7, "天！")

	// 华氏温度和摄氏温度的转化

	var huashi float32 = 78.8
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%v转化为摄氏温度为：%v", huashi, sheshi)
}
