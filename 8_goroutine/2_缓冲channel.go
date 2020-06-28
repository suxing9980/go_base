package main

import "fmt"

func main() {
	ch3 := make(chan string, 6)
	go GetData(ch3)
	for data := range ch3 {
		fmt.Println("\t 读取数据", data)
	}
}

func GetData(ch chan string) {
	defer close(ch)
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("data%d", i)
		fmt.Println("在通道放入数据", i)
	}
}
