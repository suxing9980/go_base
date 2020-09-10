package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func main() {
	p := Person{
		Id:   12,
		Name: "LiMing",
	}
	fmt.Printf("%T \n", p)
	fmt.Printf("%v \n", p)
	fmt.Printf("%+v \n", p)
	fmt.Printf("%#v \n", p)
}
