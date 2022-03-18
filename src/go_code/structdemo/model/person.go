package model

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (person Person) Say() {
	eat()
	fmt.Printf("%v 可以说 %v 地方方言 \n", person.Name, person.Address)
}

func eat() {
	var person = Person{"zhangsan", 30, "23545", "北京"}
	marshal, _ := json.Marshal(person)
	fmt.Println(marshal)
	fmt.Println("人需要先吃饭。。。")
}
