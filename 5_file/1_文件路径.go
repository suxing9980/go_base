package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 绝对路径
	filepath1 := "D:/go/ch09/picture/sold.jpg"
	filepath2 := "./picture/timg.jpg"
	fmt.Println(filepath.IsAbs(filepath1))
	fmt.Println(filepath.IsAbs(filepath2))

	fmt.Println(filepath.Rel("D:/go/ch09", filepath1))
	fmt.Println(filepath.Abs(filepath1))
	fmt.Println(filepath.Abs(filepath2))

	fmt.Println(filepath.Join(filepath1, ".."))
	fmt.Println(filepath.Join(filepath1, "."))
	fmt.Println(filepath.Join("D:/blockChain", filepath2))
}
