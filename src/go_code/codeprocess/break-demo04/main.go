package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	在golang语言中break可以指定跳出具体label的循环,默认跳出当前循环
*/
func main() {
	var num int = 0
	rand.Seed(time.Now().Unix())
	for {
		n := rand.Intn(100) + 1
		fmt.Println(n)
		num++
		if n == 99 {
			break
		}
	}
	//lable1:
	for i := 0; i < 10; i++ {
		fmt.Println("第一层循环第", i, "次")
	lable2:
		for j := 0; j < 6; j++ {
			fmt.Println("第二层循环第", j, "次")
			for k := 0; k < 5; k++ {
				fmt.Println("第三层循环第", k, "次")
				if k == 3 {
					//break lable1
					break lable2
				}
			}
		}
	}
}
