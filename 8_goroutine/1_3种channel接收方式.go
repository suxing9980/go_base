package main

import "fmt"

func SendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 100; i++ {
		ch1 <- fmt.Sprintf("发送数据%d", i)

	}
	fmt.Println("发送完毕……")
}

func main() {
	ch1 := make(chan string)
	go SendData(ch1)

	//// 1
	//for {
	//	data := <- ch1
	//	// 如果通道关闭，通道中传输的数据为各数据类型的默认值
	//	if data == "" {
	//		break
	//	}
	//	fmt.Println("从通道中读取数据方式1：", data)
	//}

	//// 2
	//for {
	//	data, ok := <- ch1
	//	// 如果通道关闭，通道中传输的数据为各数据类型的默认值
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("从通道中读取数据方式2：", data)
	//}
	////
	// 3
	for value := range ch1 {
		fmt.Println("从通道中读取数据方式2：", value)
	}
}
