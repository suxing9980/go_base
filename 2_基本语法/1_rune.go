package main

import "fmt"

func main() {
	var str = "搭建hadoop开发环境"
	str2 := []rune(str)
	for i, ch := range str2 {
		fmt.Printf("str[%d]=%c ", i, ch)
	}
}
