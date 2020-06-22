package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Circle struct {
	Point  `json:"point"`
	Radius int `json:"radius"`
}

func main() {
	c := Circle{
		Point:  Point{50, 60},
		Radius: 25,
	}
	if data, err := json.Marshal(c); err == nil {
		fmt.Printf("%s", data)
	}
}
