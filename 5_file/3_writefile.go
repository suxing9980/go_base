package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("E:/GoWorkSpace/exercise/5_file/write.txt", os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("打开文件异常", err.Error())
	} else {
		n, err := file.Write([]byte("adawdtrhthytheeferwcrswfg"))
		if err != nil {
			fmt.Println("1写入文件异常", err.Error())
		} else {
			fmt.Println("1写入ok！", n)
		}
		n, err = file.WriteString("中国字")
		if err != nil {
			fmt.Println("2写入文件异常", err.Error())
		} else {
			fmt.Println("2写入ok！", n)
		}
	}
}
