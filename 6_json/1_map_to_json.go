package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string][]string{
		"level":   {"debug"},
		"message": {"File not found!", "Stack overflow"},
	}
	// 将map解析成json格式
	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%T \n%s", data, data)
	}
}
