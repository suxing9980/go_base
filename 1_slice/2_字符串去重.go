package main

import "fmt"

func removeDuplication(strList []string) []string {
	after := strList[:1]
	for _, value := range strList[1:] {
		i := 0
		for ; i < len(after); i++ {
			if value == after[i] {
				break
			}
		}
		if i == len(after) {
			after = append(after, value)
		}
	}
	return after
}

func main() {
	list := []string{"red", "black", "red", "yellow", "yellow", "pink", "blue", "pink", "blue"}
	after := removeDuplication(list)
	fmt.Println("原来的：", list)
	fmt.Println("去重后：", after)
}
