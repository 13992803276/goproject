package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int类型的channel基本使用
	var chanInt = make(chan int, 10)
	chanInt <- 10
	chanInt <- 20
	chanInt <- 30
	chanInt <- 40
	//遍历管道取出各个元素,遍历前必须先关闭管道
	close(chanInt)
	for date := range chanInt {
		fmt.Println(strconv.Itoa(date))
	}
	//map类型的channel基本使用
	var chanMap = make(chan map[string]string, 10)
	map1 := make(map[string]string, 10)
	map1["name"] = "tom"
	map1["age"] = "10"
	chanMap <- map1
	val1 := <-chanMap
	//关闭管道后，从管道中获取数据，无法给管道加入存数据
	//close(chanMap)
	fmt.Println(val1)
}
