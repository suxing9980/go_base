package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "E:/GoWorkSpace/exercise/5_file/blocckchain.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开错误", err.Error())
		return
	} else {
		bs := make([]byte, 1024)
		n := -1
		for {
			n, err = file.Read(bs)
			if n == 0 || err == io.EOF {
				fmt.Println("读取完毕！")
				break
			}
			fmt.Println(string(bs[:n]))
		}

	}
}
