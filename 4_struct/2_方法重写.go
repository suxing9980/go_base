package main

import "fmt"

type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func main() {
	s1 := Student{
		Human:  Human{"Daniel", "15020135681", 13},
		school: "十一附中",
	}
	e1 := Employee{
		Human:   Human{"stenven", "13612058790", 32},
		company: "美团",
	}
	s1.SayHi()
	e1.SayHi()
}

func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，联系方式%s\n", h.name, h.age, h.phone)
}

func (s *Student) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，我在%s上学，联系方式%s\n", s.name, s.age, s.school, s.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，我在%s工作，联系方式%s\n", e.name, e.age, e.company, e.phone)
}
