package model

import "fmt"

type Student struct {
	Name     string
	Age      int
	Score    int
	Identity string
}

func (s *Student) NewStudent(name string, age int, score int, identity string) *Student {
	return &Student{
		Name:     name,
		Age:      age,
		Score:    score,
		Identity: identity,
	}
}

func (s *Student) Testing() {
	fmt.Printf("%v %v 正在考试....\n", s.Identity, s.Name)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func (s *Student) GetInfo() {
	fmt.Printf("学生名字 = %v 年龄 = %v 成绩 = %v 身份 = %v \n", s.Name, s.Age, s.Score, s.Identity)
}
