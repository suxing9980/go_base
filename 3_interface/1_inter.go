package main

import "fmt"

type Work interface {
	getNumber() int
}
type goWork struct {
	name string
	age  int
	//work 	Work
}

func (g *goWork) getNumber() int {
	return 100
}
func main() {
	a := goWork{
		name: "jim",
		age:  18,
	}
	n := a.getNumber()
	fmt.Println(n)
}
