package main

import (
	"encoding/json"
	"fmt"
)

type Point1 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Circle1 struct {
	Point1
	Radius int `json:"radius"`
}

type Circle2 struct {
	Point1 `json:"point"`
	Radius int `json:"radius"`
}

func main() {
	data := `{"x":80,"y":60,"radius":20}`
	var c Circle1
	json.Unmarshal([]byte(data), &c)
	//fmt.Println(c)
	fmt.Println("*****************************")
	data1 := `{"point":{"x":80,"y":60},"radius":20}`
	var c1 Circle2
	json.Unmarshal([]byte(data1), &c1)
	//fmt.Println(c1)
}
