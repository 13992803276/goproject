package model

import (
	"encoding/json"
	"fmt"
)

// person 首字母小写后只能在本包使用，要在其他包使用，需要使用工厂模式进行改造
type person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func NewPerson(name string, age int, email string, address string) *person {
	return &person{
		Name:    name,
		Age:     age,
		Email:   email,
		Address: address,
	}
}
func (person person) Say() {
	eat()
	fmt.Printf("%v 可以说 %v 地方方言 \n", person.Name, person.Address)
}

func eat() {
	var person = person{"zhangsan", 30, "23545", "北京"}
	marshal, _ := json.Marshal(person)
	fmt.Println("marshal = ", string(marshal))
	fmt.Println("人需要先吃饭。。。")
}
