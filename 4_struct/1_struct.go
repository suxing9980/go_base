package main

import "fmt"

type Address struct {
	province, city string
}

type Person struct {
	name    string
	age     int
	address *Address
}

func main() {
	p := Person{}
	p.name = "Steven"
	p.age = 35
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = &addr
	// 修改Person数据，，是否影响Address数据
	p.address.city = "昌平区"
	fmt.Println(p.address.city, p.address.province)
	fmt.Println(addr.city, addr.province)
	fmt.Println("********************")

}
