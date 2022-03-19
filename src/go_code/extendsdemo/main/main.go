package main

import (
	"fmt"
	"lib/src/go_code/extendsdemo/model"
)

/*
	golang面向对象--继承
	golang 可以多继承。
*/

type Pupil struct {
	favorite string
	model.Student
	model.Person
}

func (p *Pupil) getFavorite() {
	fmt.Printf("小学生 %v 的爱好是 %v \n", p.Name, p.favorite)
}

func main() {
	pupil := &Pupil{
		favorite: "写字",
		Student: model.Student{
			Name:     "小明",
			Age:      10,
			Identity: "小学生",
		},
		Person: model.Person{
			Id: 1,
		},
	}
	pupil.Testing()
	pupil.SetScore(80)
	pupil.GetInfo()
	pupil.getFavorite()
	pupil.SetAddress("陕西省西安市")
	address := pupil.GetAddress()
	fmt.Printf("%v %v 的家住在%v", pupil.Identity, pupil.Name, address)
}
