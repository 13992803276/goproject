package main

import (
	"encoding/json"
	"fmt"
	"lib/src/go_code/structdemo/model"
)

/*
	结构体类型：golang用结构体来实现面向对象编程（OOP），类似与java中的对象
*/
func main() {
	type a struct {
		name string
		age  int
	}
	a2 := a{"tom", 3}
	var a3 a = a{"jack", 10}
	//1。字段显式赋值
	a2.age = 30
	a2.name = "tom~"
	fmt.Println(a2, a3)
	person := model.Person{Name: "张三", Age: 20, Email: "37765432@qq.com", Address: "陕西西安"}
	person.Say()
	//2。使用指针定义一个结构体变量
	var person1 = &model.Person{}
	person1.Age = 30
	person1.Name = "jack"
	fmt.Println("person1 = ", person1)
	//3。结构体的字段可以通过tag来处理字段名称首字母大写的问题
	var person3 = model.Person{Name: "张三", Age: 30, Email: "23545", Address: "北京"}
	marshal, _ := json.Marshal(person3)
	fmt.Println(string(marshal))
}
