package main

import "fmt"

/*
	map数据结构
*/
func main() {
	//1.申明并初始化map
	var myMap = make(map[int]string, 9)
	myMap[1] = "张三"
	myMap[2] = "李四"
	fmt.Println(myMap)
	//2.map的crud
	myMap[3] = "王五" //新增修改
	//删除
	delete(myMap, 1)
	fmt.Println(myMap)
	//查询
	val, findRes := myMap[2]
	if findRes {
		print("val = ", val)
	} else {
		print("不存在该Key对应的value值")
	}
	//map 遍历 for - range
	for k, v := range myMap {
		fmt.Printf("myMap[%v] = %v \n", k, v)
	}
}
