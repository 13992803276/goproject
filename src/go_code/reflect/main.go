package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (stu Student) GetName() string {
	return stu.Name
}

//反射测试方法
func reflectTest(b interface{}) {
	//基本类型操作
	typeName := reflect.TypeOf(b)
	value := reflect.ValueOf(b)
	fmt.Println("typeName = ", typeName, "value = ", value)
	kind := value.Kind()
	fmt.Println("kind = ", kind)
	iv := value.Interface()
	num2 := iv.(int)
	fmt.Printf("num2 的类型是 %T 值是%v \n", num2, num2)
}

func reflectTest02(s interface{}) {
	//基本类型操作
	typeName := reflect.TypeOf(s)
	value := reflect.ValueOf(s)
	fmt.Println("typeName = ", typeName, "value = ", value)
	kind := value.Kind()
	fmt.Println("kind = ", kind)
	iv := value.Interface()
	fmt.Printf("iv 的类型是 %T 值是%v \n", iv, iv)
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v stu.Age = %v", stu.Name, stu.Age)
	} else {
		fmt.Printf("转化失败")
	}
}

func main() {
	var num = 100
	reflectTest(num)
	var stu = Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
